package websockethub

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type WSConnection interface {
	AddClient(c *gin.Context)
	RemoveClient()
	Broadcast()
}

type wsConnection struct {
	clients []webSocketClient
}

func NewWsConnectionHub() wsConnection {
	return wsConnection{
		clients: make([]webSocketClient, 0),
	}
}

func (wsc wsConnection) AddClient(c *gin.Context) {
	webSocketClient := NewWebSocketClient(c)
	fmt.Println(webSocketClient)
	data := append(wsc.clients, webSocketClient)
	fmt.Println(len(data))
	wsc.clients = data
}

func (wsc wsConnection) RemoveClient() {
	// TODO: Handle This
}

func (wsc wsConnection) Broadcast() {
	message := []byte("Hello!")
	fmt.Println("Broadcasting")
	fmt.Println(len(wsc.clients))
	for _, client := range wsc.clients {
		client.SendMessage(message)
	}
}
