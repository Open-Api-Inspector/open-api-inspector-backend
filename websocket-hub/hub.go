package websockethub

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type WSConnection interface {
	AddClient(websocketConnection *websocket.Conn)
	RemoveClient()
	Broadcast()
}

type wsConnection struct {
	clients []webSocketClient
}

func NewWsConnectionHub() wsConnection {
	return wsConnection{
		clients: nil,
	}
}

func (wsc wsConnection) AddClient(c *gin.Context) {
	webSocketClient := NewWebSocketClient(c)
	wsc.clients = append(wsc.clients, webSocketClient)
}

func (wsc wsConnection) RemoveClient() {
	// TODO: Handle This
}

func (wsc wsConnection) Broadcast() {
	// TODO: Broadcast
}
