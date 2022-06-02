package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleAnyRoute(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}

func main() {
	route := gin.Default()
	route.NoRoute(handleAnyRoute)
	route.Run(":8080")
}
