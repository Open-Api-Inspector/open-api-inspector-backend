package websockethub

type WebSocketCommand string

const (
	PURGE WebSocketCommand = "Purge"
)

type WebSocketCommandJson struct {
	Command   WebSocketCommand `json:"command"`
	RequestId string           `json:"requestId"`
}
