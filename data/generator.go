package data

import (
	"math/rand"
	"time"
)

const letters = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomText(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func RepetitiveText(length int, c byte) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = c
	}
	return string(b)
}
