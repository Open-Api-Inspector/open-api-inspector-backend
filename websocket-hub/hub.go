package websockethub

import (
	"github.com/gin-gonic/gin"
)

type WSConnectionHub interface {
	AddClient(c *gin.Context)
	RemoveClient()
	Broadcast()
}

type wsConnectionHub struct {
	clients []webSocketClient
}

func NewWsConnectionHub() *wsConnectionHub {
	return &wsConnectionHub{
		clients: nil,
	}
}

func (wsc *wsConnectionHub) AddClient(c *gin.Context) {
	webSocketClient := NewWebSocketClient(c)
	wsc.clients = append(wsc.clients, *webSocketClient)
}

func (wsc *wsConnectionHub) RemoveClient() {
	// TODO: Handle This
}

func (wsc *wsConnectionHub) Broadcast() {
	message := []byte("Hello!")
	for _, client := range wsc.clients {
		client.SendMessage(message)
	}
}
