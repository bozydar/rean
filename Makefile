BINARY_NAME=./rean

all: build test

build:
	go build -o ${BINARY_NAME} .

test:
	go test ./lib

run:
	go build -o ${BINARY_NAME} .
	./${BINARY_NAME}

install: all
	cp ./rean /usr/local/bin/

clean:
	go clean
