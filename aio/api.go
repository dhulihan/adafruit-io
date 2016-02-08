package aio

import (
	"os"
	"net/http"
	"log"
	"io/ioutil"
	"fmt"
	"bytes"
)

const (
 	api_url_base = "https://io.adafruit.com/api"
)

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
		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Printf("%s", err)
			os.Exit(1)
		}
		// return string[]{contents}
		feeds := []string {string(contents)}
		return feeds
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