package main

import (
	"net/http"
	requesthub "open-api-inspector-backend/request-hub"
	websockethub "open-api-inspector-backend/websocket-hub"

	"github.com/gin-gonic/gin"
)

const OPEN_API_ADDRESS = ":8080"
const WEBSOCKET_ADDRESS = ":8081"

var wsConnectionHub = websockethub.NewWsConnectionHub()
var requestHub = requesthub.NewRequestHub(wsConnectionHub)

func handleAnyRoute(c *gin.Context) {
	// TODO: Add Support for other type of request
	requestHub.AddRequest(c)
	wsConnectionHub.Broadcast()
	// TODO: Wait for response from frontend.
	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}

func main() {

	// Serve the Open API for loggin the request.
	routeOpenApi := gin.Default()
	routeOpenApi.NoRoute(handleAnyRoute)
	go routeOpenApi.Run(OPEN_API_ADDRESS)

	// Serve the API for websocket
	route_ws := gin.Default()
	route_ws.GET("/ws", wsConnectionHub.AddClient)
	route_ws.Run(WEBSOCKET_ADDRESS)
}
