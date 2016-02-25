A cli tool for [adafruit.io](https://adafruit.io) written in go.

## UNSTABLE, CHECK BACK LATER

## Installation

	go get github.com/dhulihan/adafruit-io

### Providing your key

`adafruit-go` requires your secret AIO key. 

![](key.jpg)

You can provide it by using the `--key` flag

	adafruit-io --key 'MY_KEY' [...]

`adafruit-io` also looks for this key in the environment variable `$AIO_KEY`

	AIO_KEY='MY_KEY' adafruit-io [...]

To set it permanently, add this to `~/.bashrc|.zshrc`

	export AIO_KEY='MY_KEY'

## Usage

Get all feeds

	adafruit-io

Get latest value of a feed

	adafruit-io foo

Send a value to a feed

	adafruit-io foo 9.7	