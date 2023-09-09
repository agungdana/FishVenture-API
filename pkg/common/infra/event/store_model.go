package event

import "time"

type Payload struct {
	Key         string
	Data        any
	ExpiredTime time.Duration
}
