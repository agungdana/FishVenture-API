package event

type redisChanel string

const (
	WS_REPLY   redisChanel = "reply"   //for the server to send data to the client using ws
	WS_COLLECT redisChanel = "collect" //for the server to receive data from ws
)
