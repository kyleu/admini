package util

import (
	"crypto/rand"
	"io"
)

func RandomBytes(size int) []byte {
	b := make([]byte, size)
	_, err := io.ReadFull(rand.Reader, b)
	if err != nil {
		panic("source of randomness unavailable: " + err.Error())
	}
	return b
}
