package websockethub

import (
	"fmt"
	requestmanager "open-api-inspector-backend/request-manager"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type WSConnectionHub interface {
	AddClient(c *gin.Context)
	RemoveClient(clientId string)
	Broadcast()
	updateClientState(clientId string, status WebSocketStatus)
}

type wsConnectionHub struct {
	clients map[string]webSocketClient
}

func NewWsConnectionHub() *wsConnectionHub {
	return &wsConnectionHub{
		clients: make(map[string]webSocketClient, 0),
	}
}

func (wsc *wsConnectionHub) AddClient(c *gin.Context, apiRequests map[string]requestmanager.ApiRequest) {
	clientId := uuid.New().String()
	webSocketClient := NewWebSocketClient(c, clientId, wsc)
	// Send old request to the new client
	message := []byte("Test")
	for _, request := range apiRequests {
		fmt.Println(request)
		webSocketClient.SendMessage(message)
	}
	wsc.clients[clientId] = *webSocketClient
}

func (wsc *wsConnectionHub) RemoveClient(clientId string) {
	// TODO: Handle this
}

func (wsc *wsConnectionHub) updateClientState(clientId string, status WebSocketStatus) {
	if status == CLOSED {
		delete(wsc.clients, clientId)
	}
}

func (wsc *wsConnectionHub) Broadcast() {
	message := []byte("Hello!")
	for _, element := range wsc.clients {
		element.SendMessage(message)
	}
}
