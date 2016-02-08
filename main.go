package main

import (
	"github.com/dhulihan/adafruit-io-go/aio"
	"log"
)

func main() {
	log.Printf("Starting...")
	if aio_key := aio.GetKey(); aio_key == "" {
		log.Fatal()
	} else {
		log.Printf("adafruit.io key found via $AIO_KEY\n")
		log.Printf("Feeds: %s", aio.Feeds())
	}
}
