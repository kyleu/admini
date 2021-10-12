package cutil

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"os"

	"go.uber.org/zap"

	"github.com/kyleu/admini/app/util"
)

var key string

func EncryptMessage(message string, logger *zap.SugaredLogger) (string, error) {
	byteMsg := []byte(message)
	block, err := aes.NewCipher(getKey(logger))
	if err != nil {
		return "", fmt.Errorf("could not create new cipher: %v", err)
	}

	cipherText := make([]byte, aes.BlockSize+len(byteMsg))
	iv := cipherText[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return "", fmt.Errorf("could not encrypt: %v", err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], byteMsg)

	return base64.StdEncoding.EncodeToString(cipherText), nil
}

func DecryptMessage(message string, logger *zap.SugaredLogger) (string, error) {
	cipherText, err := base64.StdEncoding.DecodeString(message)
	if err != nil {
		return "", fmt.Errorf("could not base64 decode: %v", err)
	}

	block, err := aes.NewCipher(getKey(logger))
	if err != nil {
		return "", fmt.Errorf("could not create new cipher: %v", err)
	}

	if len(cipherText) < aes.BlockSize {
		return "", fmt.Errorf("invalid ciphertext block size")
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)

	return string(cipherText), nil
}

func getKey(logger *zap.SugaredLogger) []byte {
	if key == "" {
		env := util.AppKey + "_encryption_key"
		key = os.Getenv(env)
		if key == "" {
			logger.Warnf("using default encryption key\nset environment variable [%s] to save sessions between restarts", env)
			key = util.AppKey + "_secret"
		}
		for i := len(key); i < 16; i++ {
			key += " "
		}
		key = key[:16]
	}
	return []byte(key)
}
