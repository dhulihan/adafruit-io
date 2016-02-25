package aio

import (
	"os"
)

type Context struct {
	api_key, name string
}

func NewContext(key string) Context {
	key = os.Getenv("AIO_KEY")
	if key == "" {
		panic("$AIO_KEY not specified.")
	}
	return Context{api_key: key}
}
