package util

import (
	"math/rand"
	crypto "crypto/rand"
	"io"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz0123456789"

var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

func RandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func RandomInt(maxInclusive int) int {
	return seededRand.Intn(maxInclusive)
}

func RandomBytes(size int) []byte {
	b := make([]byte, size)
	_, err := io.ReadFull(crypto.Reader, b)
	if err != nil {
		panic("source of randomness unavailable: " + err.Error())
	}
	return b
}
