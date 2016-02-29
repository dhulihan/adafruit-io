package aio

import (
	log "github.com/Sirupsen/logrus"
	"testing"
)

func Setup() Context {
	log.SetLevel(log.DebugLevel)
	return Context{api_key: "fake-key-here"}
}

func TestFeeds(t *testing.T) {
	a := Setup()
	_, err := Feeds(&a)
	if err != nil {
		t.Error(err)
	}
}
