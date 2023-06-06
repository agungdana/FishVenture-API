package client

type Message struct {
	Subject string
	Reply   string
	Data    []byte
}

type NatsClientMesage func(ncm *Message)

type WsRequest struct {
	Type string `json:"type,omitempty"`
	Data []byte `json:"data,omitempty"`
}

type WsResponse struct {
	Type string `json:"type,omitempty"`
	Data any    `json:"data,omitempty"`
}
