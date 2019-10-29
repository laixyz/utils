package utils

import (
	"math/rand"
	"time"
)

// RandInt64 随机数 int64
func RandInt64(min, max int64) int64 {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Int63n(max-min+1)
}
