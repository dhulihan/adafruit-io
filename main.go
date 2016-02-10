package main

import (
	"github.com/codegangsta/cli"
	"os"
	log "github.com/Sirupsen/logrus"
	"net/http"
	"io/ioutil"
	"encoding/json"
	// "reflect"
	"fmt"
)

func init() {
	log.SetLevel(log.InfoLevel)
}

var (
	api_url_base = "https://io.adafruit.com/api"
)

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
		cli.StringFlag{
			Name: "format, f",
			Value: "text",
			Usage: "Desired output format. Options: json, text (default)",
		}, cli.StringFlag{
			Name: "key, k",
			Usage: "Your adafruit.io secret key. $AIO_KEY tried first",
			EnvVar: "AIO_KEY",
		},		
		cli.BoolFlag{
			Name:  "debug, d",
			Usage: "Enable to see debug messages",
		},
	}
	// app.Action = MainAction
	app.Commands = []cli.Command {
		{
			Name:      "feeds",
			Aliases:   []string{"f"},
			Usage:     "Get Feeds",
			Action:    FeedsAction,
		},		
		{
			Name:      "key",
			Aliases:   []string{"k"},
			Usage:     "print AIO key",
			Action:    KeyAction,
		},
	}	
	Run(app)
}

func Run(app *cli.App) {
	app.Before = func(c *cli.Context) error {
		if c.GlobalString("key") == "" {
			log.Fatal("No aio key provided. Use --key KEY_HERE or export AIO_KEY=KEY_HERE")
		}

		if c.GlobalBool("debug") {
			log.SetLevel(log.DebugLevel)
			log.Debug("Debug Mode ON")		
		}
		return nil
	}
	app.Run(os.Args)
}

func MainAction(c *cli.Context) {
	log.Debug("Starting...")
	log.Debug("using adafruit.io key ", c.GlobalString("key"))
	log.Debug("Args: ", c.Args())

	if len(c.Args()) == 0 {
		log.Debug("No action specified")
		fmt.Println("Please provide a subcommand. Run --help for more. ")
	}
}

func KeyAction(c *cli.Context) {
	fmt.Println(c.String("key"))
}

func FeedsAction(c *cli.Context) {
	feeds := Feeds(c)
	if len(feeds) > 0 {
		for _, feed := range feeds  {
			fmt.Println(feed)
		}	
		 
	} else {
		fmt.Println("No feeds found.")
	}	
}

func Feeds(c *cli.Context) []string {
	url := api_url_base + "/feeds"
	log.Debug("GET", url)

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("x-aio-key", c.GlobalString("key"))
	client := &http.Client{}
	resp, err := client.Do(req)
	var feeds_sl []string

	// response, err := http.Get(url)
	if err != nil {
		log.WithField("error", err).Fatal("Reponse error")
	} else {
		// close Body after everything is done
		defer resp.Body.Close()
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.WithField("error", err).Fatal("Error reading response body")
		}
		log.WithField("response", string(b)).Debug("Repsonse:")

		var f interface{}
		if json.Unmarshal(b, &f) != nil {
			log.Fatal(err)
		}
		//log.WithField("refled.TypeOf(f)", reflect.TypeOf(f)).Debug()

		// someone call Tom Hanks, we gunna cast away
		feeds := f.([]interface {})
		log.WithField("feeds", len(feeds)).Debug("Found feeds")
		for _, feed_iface := range feeds {
			feed := feed_iface.(map[string]interface{})
			feeds_sl = append(feeds_sl, feed["name"].(string))
			log.WithField("feed", feed["name"]).Debug("Found Feed")
			// for attr := range feed {
			// 	log.WithField(attr, feed[attr]).Debug("Attr")
			// } 
		} 
	}		
	return feeds_sl
}