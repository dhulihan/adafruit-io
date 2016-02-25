MAIN=adafruit-io
VAL:=$(shell echo $$RANDOM | head -c 3)
OPTS=-d
all: clean build test

build:
	@echo "+ $@"
	@go build -v 

test:
	@echo "+ $@"
	go test ./...

test-send: build
	@echo "+ $@"
	./adafruit-io -d send foo $(VAL)

bats: build
	@echo "+ $@"
	bats *.bats

clean:
	@echo "+ $@"
	@rm $(MAIN)

env:
	@echo "+ $@"
	@echo $(VAL) 

watch:
	    (while true; do make build.log; sleep 2; done) | grep -v 'make\[1\]'

build.log: #./src/*
	    go run main.go | tee build.log
