package requestmanager

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiRequest struct {
	RequestId          string      `json:"requestId"`
	Url                string      `json:"url"`
	RequestHeader      http.Header `json:"requestHeader"`
	RequestBody        []byte      `json:"requestBody"`
	ResponseStatusCode int         `json:"responseStatusCode"`
}

func NewApiRequest(c *gin.Context, requestId string) *ApiRequest {
	requestBody, _ := io.ReadAll(c.Request.Body)
	return &ApiRequest{
		RequestId:     requestId,
		Url:           c.Request.RequestURI,
		RequestHeader: c.Request.Header,
		RequestBody:   requestBody,
	}
}
