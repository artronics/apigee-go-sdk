.SILENT:

-include .env

token:
	get_token -p $(username)

build:
	go build -o build/apigee

run: build
	./build/apigee $(opt) -t ${APIGEE_TOKEN}

dev: build
	./build/apigee $(opt_dev) -t ${APIGEE_TOKEN}

.PHONY: build run
