

default: build_docker
	echo build complete

deps:
	@echo deps: "\n\n**** Downloading deps ****\n\n"
	go get github.com/op/go-logging
	go get github.com/garyburd/redigo/redis
	go get github.com/gorilla/mux

build: clean deps
	#go build -o bin/cart src/cart/main.go
	@echo build: "\n\n**** Building ****\n\n"
	env GOOS=linux GOARCH=amd64 go build -o bin/cart  src/cart/main.go
	shasum --algorithm 256 bin/* > bin/sha256sums

build_docker:
	@echo build_docker: "\n\n**** Calling docker golang image ****\n\n"
	docker run --rm -it --mount type=bind,src=$(PWD)/,dst=/cart --workdir=/cart golang make build

clean:
	@echo build: "\n\n**** Removing old artefacts ****\n\n"
	go clean
	-rm -rf ./bin/*
	echo clean: make complete

docker: build
	@#@echo docker: "-- Removing all Cart images --"
	@#-docker rmi -f $(docker images cart -q) 
	@echo docker: "\n\n**** Creating a new Cart image ****\n\n"

	docker build -t cart:0.0.1 .

install: docker
	echo install: make complete

.PHONY: build clean docker install build_docker deps
