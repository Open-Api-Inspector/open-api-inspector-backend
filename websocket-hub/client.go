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
}

type webSocketClient struct {
	ginContext          *gin.Context
	webSocketConnection *websocket.Conn
	status              WebSocketStatus
}

func NewWebSocketClient(c *gin.Context) webSocketClient {

	ws, _ := wsUpgrader.Upgrade(c.Writer, c.Request, nil)
	wsc := webSocketClient{
		ginContext:          c,
		webSocketConnection: ws,
		status:              INIT,
	}
	// Handle Each client in each thread.
	go wsc.HandleSocket()

	wsc.status = CONNECTED

	return wsc
}

func (wsc webSocketClient) HandleSocket() {
	defer wsc.webSocketConnection.Close()
	for {
		mt, message, err := wsc.webSocketConnection.ReadMessage()
		if message != nil {
			fmt.Print(string(message))
		}
		if err != nil {
			fmt.Printf(err.Error())
			break
		}
		wsc.webSocketConnection.WriteMessage(mt, message)
	}
	wsc.status = CLOSED
}
