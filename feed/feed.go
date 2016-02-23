package feed

import (
	"encoding/json"
	// "fmt"
	log "github.com/Sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	API_URL_BASE = "https://io.adafruit.com/api"
)

type AdafruitIOContext struct {
	key string
}

type Feed struct {
	ID                                          rune
	Key, Name, Description, Status              string
	History, Enabled                            bool
	Unit_Type, Unit_Symbol, License, Visibility string
	Last_Value                                  string
	Created_At, Updated_At                      time.Time
}

// Get slice of feeds
func Feeds(aio_key string) []Feed {
	url := API_URL_BASE + "/feeds"
	log.Debug("GET ", url)

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("x-aio-key", aio_key)
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.WithField("error", err).Fatal("Reponse error")
	} else {
		// close Body after everything is done
		defer resp.Body.Close()
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.WithField("error", err).Fatal("Error reading response body")
		}
		log.WithField("response", string(b)).Debug("Response:")
		var feeds []Feed
		err = json.Unmarshal(b, &feeds)
		if err != nil {
			log.Fatal("Could not unmarshal:", err)
			return nil
		}
		log.WithField("feeds", feeds).Debug("Found feeds")
		return feeds
	}
	return nil
}

func (f *Feed) Send(value interface{}) bool {
	return true
}

func (f *Feed) Get(key string) interface{} {
	return "NOT IMPLEMENTED"
}

func FeedInfo(a *AdafruitIOContext, id string) map[string]interface{} {
	url := API_URL_BASE + "/feeds/" + string(id)
	log.Debug("GET ", url)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("x-aio-key", a.key)
	client := &http.Client{}
	resp, _ := client.Do(req)

	log.WithField("status", resp.Status).Debug("response.Status")
	switch {
	case resp.StatusCode == 404:
		log.Fatal("404: Feed not found")
		return nil
	case resp.StatusCode == 200:
		b, _ := ioutil.ReadAll(resp.Body)

		var f interface{}
		json.Unmarshal(b, &f)
		log.WithField("response", string(b)).Debug("Response:")

		attrs := f.(map[string]interface{})
		log.WithField("len(attrs)", len(attrs)).Debug("Found some attributes")
		return attrs
	}
	return nil
}
