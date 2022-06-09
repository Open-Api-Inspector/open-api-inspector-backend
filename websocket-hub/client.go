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
}

type webSocketClient struct {
	ginContext          *gin.Context
	webSocketConnection *websocket.Conn
	status              WebSocketStatus
}

func NewWebSocketClient(c *gin.Context) *webSocketClient {

	ws, _ := wsUpgrader.Upgrade(c.Writer, c.Request, nil)
	wsc := &webSocketClient{
		ginContext:          c,
		webSocketConnection: ws,
		status:              INIT,
	}
	// Handle Each client in each thread.
	go wsc.HandleSocket()

	wsc.status = CONNECTED

	return wsc
}

func (wsc *webSocketClient) HandleSocket() {
	defer wsc.webSocketConnection.Close()
	for {
		_, message, err := wsc.webSocketConnection.ReadMessage()
		if message != nil {
			fmt.Println(string(message))
		}
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		// wsc.webSocketConnection.WriteMessage(mt, message)
	}
	wsc.status = CLOSED
}

func (wsc *webSocketClient) SendMessage(message []byte) {
	if wsc.status != CONNECTED {
		return
	}
	fmt.Println("Sending message")
	wsc.webSocketConnection.WriteMessage(1, message)
}
