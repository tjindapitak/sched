package external

import (
	"time"
)

func CurrentTime() int64 {
	return (time.Now().UnixNano() / (int64(time.Second) / int64(time.Nanosecond)))
}

func GetEpocSecond(in time.Time) int64 {
	return (in.UnixNano() / (int64(time.Second) / int64(time.Nanosecond)))
}
