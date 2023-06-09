package ptime

import (
	"time"
)

func SetDefaultTimeToUTC() {
	time.Local = time.UTC
}

func Today() time.Time {
	return time.Now()
}
