package main

import (
	websockethub "ApiInspector/websocket-hub"
	"net/http"

	"github.com/gin-gonic/gin"
)

const OPEN_API_ADDRESS = ":8080"
const WEBSOCKET_ADDRESS = ":8081"

var wsConnectionHub = websockethub.NewWsConnectionHub()

func handleAnyRoute(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}

func main() {
	// Serve the Open API for loggin the request.
	route_open_api := gin.Default()
	route_open_api.NoRoute(handleAnyRoute)
	go route_open_api.Run(OPEN_API_ADDRESS)

	// Serve the API for websocket
	route_ws := gin.Default()
	route_ws.GET("/ws", wsConnectionHub.AddClient)
	route_ws.Run(WEBSOCKET_ADDRESS)
}
