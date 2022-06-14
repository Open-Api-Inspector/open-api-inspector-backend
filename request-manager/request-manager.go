package requestmanager

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type RequestManager interface {
	AddRequest(c *gin.Context)
}

type requestManager struct {
	ApiRequests     map[string]ApiRequest
	ApiRequestOrder []ApiRequest
}

func NewRequestManager() *requestManager {
	return &requestManager{
		ApiRequests:     make(map[string]ApiRequest, 0),
		ApiRequestOrder: make([]ApiRequest, 0),
	}
}

func (r *requestManager) AddRequest(c *gin.Context) ApiRequest {
	requestId := uuid.New().String()
	apiRequest := NewApiRequest(c, requestId)
	r.ApiRequests[requestId] = *apiRequest
	r.ApiRequestOrder = append(r.ApiRequestOrder, *apiRequest)
	fmt.Println("Adding new request to request pool with id", requestId)
	return *apiRequest
}
