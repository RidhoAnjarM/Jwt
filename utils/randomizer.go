package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func UUID() string {
	uuid := make([]byte, 16)
	_, err := rand.Read(uuid)
	if err != nil {
		return ""
	}

	uuid[6] = (uuid[6] & 0x0f) | 0x40
	uuid[8] = (uuid[8] & 0x3f) | 0x80

	now := time.Now().UnixNano()
	uuid[0] = byte(now >> 54)
	uuid[1] = byte(now >> 48)
	uuid[2] = byte(now >> 42)
	uuid[3] = byte(now >> 36)
	uuid[4] = byte(now >> 30)
	uuid[5] = byte(now >> 24)
	uuid[6] = byte(now >> 18)
	uuid[7] = byte(now >> 12)
	uuid[8] = byte(now >> 6)

	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:])
}

func PIN(length int) string {
	chars := "0123456789"
	pinBytes := make([]byte, length)

	for i := range pinBytes {
		pinBytes[i] = chars[rand.Intn(len(chars))]
	}

	return string(pinBytes)
}
