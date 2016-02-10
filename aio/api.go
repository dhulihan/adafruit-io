package aio

import (
	"net/http"
	"github.com/Sirupsen/logrus"
	"io/ioutil"
	"fmt"
	"bytes"
	"encoding/json"
	"reflect"
	"time"
)

var log = logrus.New()

const (
 	api_url_base = "https://io.adafruit.com/api"
)

type Feed struct {
	id string
	name string
	key string
	description string
	unit_type string
	unit_symbol string
	status string
	visibility string
	enabled string
	created_at time.Time
	updated_at time.Time
}

func Feeds() []string {	
	url := api_url_base + "/feeds"
	log.Debug("GET", url)

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("x-aio-key", GetKey())
	client := &http.Client{}
	resp, err := client.Do(req)

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

		log.Debug("response: ", string(b))

		var f interface{}
		if json.Unmarshal(b, &f) != nil {
			log.Fatal(err)
		}
		fmt.Printf("f is %s\n\n", reflect.TypeOf(f))

		// someone call Tom Hanks, we gunna cast away
		feeds := f.([]interface {})
		log.WithField("feeds", len(feeds)).Debug("Found feeds")
		for _, feed_iface := range feeds {
			feed := feed_iface.(map[string]interface{})
			log.WithField("feed", feed["name"]).Debug("Found Feed")
			for attr := range feed {
				log.WithField(attr, feed[attr]).Debug("Attr")
			} 
		} 
	}
	return []string{""}
}

func UpdateFeed(feed_id string) (bool, error) {	
	url := api_url_base + "/api/feeds"
	fmt.Println("URL: ", url)

	// Set body data
	// var jsonStr = []byte(`{"title":"Buy cheese and bread for breakfast."}`)
	var jsonStr = []byte(`{"last_value":10`)
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("x-aio-key", GetKey())
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	return true, nil
}