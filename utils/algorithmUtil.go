package utils

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

// RandString generates a random string with length 'n'
func RandString(n int) string {
	rand.Seed(time.Now().UnixNano())

	data := make([]byte, n)
	for idx := range data {
		data[idx] = letters[rand.Int63()%int64(len(letters))]
	}
	return string(data)
}

// RandomUUIDString helps to generate a random UUID-V4 string
func RandomUUIDString() string {
	return uuid.New().String()
}
