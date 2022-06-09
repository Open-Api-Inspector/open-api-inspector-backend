package requestmanager

import (
	"fmt"
	"io"
	websockethub "open-api-inspector-backend/websocket-hub"

	"github.com/gin-gonic/gin"
)

type RequestManager interface {
	AddRequest(c *gin.Context)
}

type requestManager struct {
	apiRequests     []ApiRequest
	WsConnectionHub websockethub.WSConnectionHub
}

func NewRequestManager(wsConnectionHub websockethub.WSConnectionHub) *requestManager {
	return &requestManager{
		WsConnectionHub: wsConnectionHub,
	}
}

func (r *requestManager) AddRequest(c *gin.Context) {
	requestBody, _ := io.ReadAll(c.Request.Body)
	apiRequest := NewApiRequest(c.Request.Header, requestBody)
	r.apiRequests = append(r.apiRequests, apiRequest)
	fmt.Println("Request Num: ", len(r.apiRequests))
}
