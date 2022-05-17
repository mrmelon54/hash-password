.PHONY: build

build:
	mkdir dist/ -p
	go build -o ./dist/hash-password .
