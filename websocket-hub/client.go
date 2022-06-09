package websockethub

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type WebSocketStatus int

const (
	INIT      WebSocketStatus = 0
	CONNECTED                 = 1
	CLOSED                    = 2
)

var wsUpgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type WebSocketClient interface {
	HandleSocket()
	SendMessage(s []byte)
	setStatus()
	CloseConnection()
}

type webSocketClient struct {
	ginContext          *gin.Context
	webSocketConnection *websocket.Conn
	status              WebSocketStatus
	ClientId            string
	wscHub              *wsConnectionHub
}

func NewWebSocketClient(c *gin.Context, id string, wscHub *wsConnectionHub) *webSocketClient {

	ws, _ := wsUpgrader.Upgrade(c.Writer, c.Request, nil)
	wsc := &webSocketClient{
		ginContext:          c,
		webSocketConnection: ws,
		status:              INIT,
		ClientId:            id,
		wscHub:              wscHub,
	}
	fmt.Println("New connection with ID ", id)
	// Handle Each client in each thread.
	go wsc.HandleSocket()
	wsc.setStatus(CONNECTED)

	return wsc
}

func (wsc *webSocketClient) HandleSocket() {
	defer wsc.CloseConnection()
	for {
		_, message, err := wsc.webSocketConnection.ReadMessage()
		if message != nil {
			// Broadcast incomming message
			fmt.Println(string(message))
		}
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		// wsc.webSocketConnection.WriteMessage(mt, message)
	}
}

func (wsc *webSocketClient) CloseConnection() {
	wsc.webSocketConnection.Close()
	wsc.setStatus(CLOSED)
}

func (wsc *webSocketClient) setStatus(webSocketStatus WebSocketStatus) {
	fmt.Println("Client ", wsc.ClientId, " status has been updated to", webSocketStatus)
	wsc.status = webSocketStatus
	// Notify the observer
	wsc.wscHub.updateClientState(wsc.ClientId, webSocketStatus)
}

func (wsc *webSocketClient) SendMessage(message []byte) {
	if wsc.status != CONNECTED {
		return
	}
	fmt.Println("Sending message to client ", wsc.ClientId)
	wsc.webSocketConnection.WriteMessage(1, message)
}
