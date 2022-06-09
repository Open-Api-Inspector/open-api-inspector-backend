package requesthub

import (
	"io"
	websockethub "open-api-inspector-backend/websocket-hub"

	"github.com/gin-gonic/gin"
)

type RequestHub interface {
	AddRequest(c *gin.Context)
}

type requestHub struct {
	apiRequests     []ApiRequest
	WsConnectionHub websockethub.WSConnection
}

func NewRequestHub(wsConnectionHub websockethub.WSConnection) requestHub {
	return requestHub{
		WsConnectionHub: wsConnectionHub,
	}
}

func (r requestHub) AddRequest(c *gin.Context) {
	requestBody, _ := io.ReadAll(c.Request.Body)
	apiRequest := NewApiRequest(c.Request.Header, requestBody)
	r.apiRequests = append(r.apiRequests, apiRequest)

	// TODO: Broadcast new request.
	r.WsConnectionHub.Broadcast()

}
