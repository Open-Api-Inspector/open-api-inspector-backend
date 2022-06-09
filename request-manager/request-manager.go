package requestmanager

import (
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type RequestManager interface {
	AddRequest(c *gin.Context)
}

type requestManager struct {
	ApiRequests map[string]ApiRequest
}

func NewRequestManager() *requestManager {
	return &requestManager{
		ApiRequests: make(map[string]ApiRequest, 0),
	}
}

func (r *requestManager) AddRequest(c *gin.Context) {
	requestBody, _ := io.ReadAll(c.Request.Body)
	requestId := uuid.New().String()
	apiRequest := NewApiRequest(requestId, c.Request.Header, requestBody)
	r.ApiRequests[requestId] = apiRequest
	fmt.Println("Request Num: ", len(r.ApiRequests))
}
