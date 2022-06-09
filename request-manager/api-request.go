package requestmanager

import (
	"net/http"
)

type ApiRequest interface {
}

type apiRequest struct {
	RequestId          string
	RequestHeader      http.Header
	RequestBody        []byte
	ResponseStatusCode int
}

func NewApiRequest(requestId string, header http.Header, body []byte) *apiRequest {
	return &apiRequest{
		RequestId:     requestId,
		RequestHeader: header,
		RequestBody:   body,
	}
}
