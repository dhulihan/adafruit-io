package aio

import (
	"os"
	"log"
)

// Get secret adafruit.io key. First try envvar $AIO_KEY.
func GetKey() (string) {
	envkey := os.Getenv("AIO_KEY")
	if envkey == "" {
		log.Fatal("$AIO_KEY undefined")
		os.Exit(1)
	}
	return envkey
}
