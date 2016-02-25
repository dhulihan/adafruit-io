package aio

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	API_URL_BASE = "https://io.adafruit.com/api"
)

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

func Find(id string, a *Context) (Feed, error) {
	url := API_URL_BASE + "/feeds/" + id
	log.Debug("GET ", url)

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("x-aio-key", a.api_key)
	client := &http.Client{}
	resp, err := client.Do(req)

	var feed Feed
	if err != nil {
		return feed, err
	} else {
		// close Body after everything is done
		defer resp.Body.Close()
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.WithField("error", err).Fatal("Error reading response body")
		}
		log.WithField("response", string(b)).Debug("Response:")
		err = json.Unmarshal(b, &feed)
		if err != nil {
			return feed, err
		}
		log.WithField("feed", feed).Debug("Found feed")
		return feed, nil
	}
}

func Send(id string, value string, a *Context) error {
	url := API_URL_BASE + "/feeds/" + id
	log.Debug("PUT ", url)

	var jsonStr = []byte(fmt.Sprintf(`{"last_value": %s}`, value))
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonStr))
	log.WithField("req body", string(jsonStr)).Debug()
	req.Header.Set("x-aio-key", a.api_key)
	req.Header.Set("Content-type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return err
	} else {
		// close Body after everything is done
		defer resp.Body.Close()
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		log.WithField("response", string(b)).Debug("Response:")
		if resp.StatusCode != 200 {
			msg := fmt.Sprintf("Got HTTP Status: %s", resp.Status)
			return errors.New(msg)
		}
		return nil

	}
}

func (f *Feed) Send(value string, a *Context) error {
	url := API_URL_BASE + "/feeds/" + string(f.ID)
	log.Debug("PUT ", url)

	req, err := http.NewRequest("PUT", url, nil)
	req.Header.Set("x-aio-key", a.api_key)
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return err
	} else {
		// close Body after everything is done
		defer resp.Body.Close()
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		log.WithField("response", string(b)).Debug("Response:")

		if resp.StatusCode != 200 {
			msg := fmt.Sprintf("Got HTTP Status: %s", resp.StatusCode)
			return errors.New(msg)
		}
		return nil

	}
}
