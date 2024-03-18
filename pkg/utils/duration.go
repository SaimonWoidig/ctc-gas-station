package utils

import (
	"math/rand"
	"time"
)

func RandBetween(min int, max int) int {
	return rand.Intn(max-min+1) + min
}
func RandDurationBetween(min time.Duration, max time.Duration) time.Duration {
	return time.Duration(RandBetween(int(min.Milliseconds()), int(max.Milliseconds()))) * time.Millisecond
}
