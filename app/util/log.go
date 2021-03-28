package util

import (
	"log"
)

func LogInfo(msg string, args ...interface{}) {
	log.Printf(msg, args...)
}

func LogWarn(msg string, args ...interface{}) {
	log.Printf("WARN: "+msg, args...)
}
