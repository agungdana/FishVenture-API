package event

type ClientMessages struct {
	Topic  string         `json:"topic"`
	Action string         `json:"action"`
	CtxMap map[string]any `json:"ctxMap"`
	Data   any            `json:"data"`
}

type ClientMessaging func(m *ClientMessages)
