package zeroconf

import (
	"log"
	"os"
)

var DEBUG = os.Getenv("DEBUG") != ""

func Warn(message string, err error) {
	if DEBUG {
		log.Printf("[WARN] zeroconf: %s: %v", message, err)
	}
}
