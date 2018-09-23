

default: build_docker
	echo build complete

deps:
	@echo deps: "\n\n**** Downloading deps ****\n\n"
	go get github.com/op/go-logging
	#go get github.com/garyburd/redigo/redis
	# go get github.com/go-redis/redis
	# go get github.com/vmihailenco/msgpack
	go get github.com/gorilla/mux


build: clean deps
	#go build -o bin/postbar src/main.go
	@echo build: "\n\n**** Building ****\n\n"
	env GOOS=linux GOARCH=amd64 go build -o bin/postbar src/main.go
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
	@-docker rmi -f $(docker images postbar -q) 
	@echo docker: "\n\n**** Creating a new Cart image ****\n\n"

	docker build -t postbar:0.0.1 .

deploy: build
	@echo deploy: "\n\n**** Deploy ****\n\n"
	docker tag postbar:0.0.1 alexgamas/postbar:0.0.1
	docker push alexgamas/postbar:0.0.1

install: docker
	echo install: make complete

.PHONY: build clean docker install build_docker deps
