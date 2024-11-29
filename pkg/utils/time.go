package utils

import (
	"time"
)

func TimeNow() time.Time {
	return time.Now()
}

func Sleep(t int) {
	time.Sleep(time.Duration(t) * time.Millisecond)
}
