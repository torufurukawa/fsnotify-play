.PHONY: all env build run test

GO = env GOPATH=$(shell pwd) go

all:
	@echo env ..... create environment
	@echo build ... build executable file
	@echo test .... test

env:
	$(GO) get "gopkg.in/fsnotify.v1"

build: main

main: main.go
	$(GO) build -o main main.go

run: build
	./main

test: build
	# setup test directories
	if [ -a _work ]; then rm -r _work; fi
	mkdir -p _work/foo
	mkdir -p _work/bar
	touch _work/foo/a
	# start monitoring
	./main & echo "$$!" > main.pid
	sleep 1
	# operate on dir
	touch _work/foo/x
	touch _work/bar/y
	# stop monitoring
	sleep 1
	kill `cat main.pid`
	rm main.pid
