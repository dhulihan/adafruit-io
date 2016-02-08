watch:
	    (while true; do make build.log; sleep 2; done) | grep -v 'make\[1\]'

build.log: #./src/*
	    go run main.go | tee build.log
