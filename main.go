package main

import (
	"github.com/dhulihan/adafruit-io-go/aio"
	"github.com/codegangsta/cli"
	"os"
	log "github.com/Sirupsen/logrus"
)

func init() {
	log.SetLevel(log.InfoLevel)
}

func main() {
	app := cli.NewApp()
	app.Name = "adafruit-io"
	app.Version = "1.0.0"
	app.Usage = "Send data to your adafruit.io dashboard"

	app.Flags = []cli.Flag {
		// cli.StringFlag{
		// 	Name: "verbose",
		// 	Value: "true",
		// 	Usage: "Enable to see debug messages",
		// }
		cli.BoolFlag{
			Name:  "debug, d",
			Usage: "Enable to see debug messages",
		},
	}

	app.Action = func(c *cli.Context) {
		if c.Bool("debug") {
			log.SetLevel(log.DebugLevel)
		}
		log.Debug("Starting...")

		if aio_key := aio.GetKey(); aio_key == "" {
			log.Fatal(`$AIO_KEY not set. Try export AIO_KEY="KEY GOES HERE"`)
		} else {
			log.Debug("adafruit.io key found via $AIO_KEY\n")
			feeds := aio.Feeds()
			log.WithField("feeds", len(feeds)).Debug("Found feeds")
		}
	}
	app.Run(os.Args)
}
