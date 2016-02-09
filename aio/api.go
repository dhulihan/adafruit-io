package aio

import (
	"os"
	"net/http"
	"log"
	"io/ioutil"
	"fmt"
	"bytes"
	"encoding/json"
	"reflect"
	"time"
)

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
	log.Println("GET", url)

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("x-aio-key", GetKey())
	client := &http.Client{}
	resp, err := client.Do(req)

	// response, err := http.Get(url)
	if err != nil {
		log.Printf("%s", err)
		os.Exit(1)
	} else {
		// close Body after everything is done
		defer resp.Body.Close()
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Printf("%s", err)
			os.Exit(1)
		}

		fmt.Printf("response: %s\n", b)

		var f interface{}
		if json.Unmarshal(b, &f) != nil {
			log.Fatal(err)
		}
		fmt.Printf("f is %s\n\n", reflect.TypeOf(f))

		// someone call Tom Hanks, we gunna cast away
		feeds := f.([]interface {})
		log.Println("Found", len(feeds), "feeds")
		for _, feed_int := range feeds {
			feed := feed_int.(map[string]interface{})
			fmt.Printf("\t%s\n", feed["name"])
			// delete(feed, "name")
			for attr := range feed {
				fmt.Printf("\t\t%s: %s\n", attr, feed[attr])			
			} 
			fmt.Println()
		} 
	}
	return nil
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