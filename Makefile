BINARY_NAME=./rean

all: build test

build:
	go build -o ${BINARY_NAME} .

test:
	go test ./lib

run:
	go build -o ${BINARY_NAME} .
	./${BINARY_NAME}

clean:
	go clean
