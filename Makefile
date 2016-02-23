MAIN=adafruit-io
#VAL=$(echo $RANDOM | head -c 3)
VAL?=10
OPTS=-d
all: clean build test

build:
	@echo "+ $@"
	@go build -v 

test: build
	@echo "+ $@"
	./$(MAIN) -d
	#./$(MAIN) -d feeds
	# ./$(MAIN) -d info foo
	#./$(MAIN) -d get foo
	#./$(MAIN) $(OPTS) send foo $(VAL)

clean:
	@echo "+ $@"
	@rm $(MAIN)

watch:
	    (while true; do make build.log; sleep 2; done) | grep -v 'make\[1\]'

build.log: #./src/*
	    go run main.go | tee build.log
