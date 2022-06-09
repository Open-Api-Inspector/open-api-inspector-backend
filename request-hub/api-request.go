package requesthub

import (
	"net/http"

	"github.com/google/uuid"
)

type ApiRequest interface {
}

type apiRequest struct {
	RequestId          string
	RequestHeader      http.Header
	RequestBody        []byte
	ResponseStatusCode int
}

func NewApiRequest(header http.Header, body []byte) *apiRequest {
	return &apiRequest{
		RequestId:     uuid.New().String(),
		RequestHeader: header,
		RequestBody:   body,
	}
}
