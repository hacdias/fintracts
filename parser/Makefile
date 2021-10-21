.PHONY: clean docker

clean:
	rm -f parser

build:
	go build -o parser cmd/*.go

docker:
	docker build .. -t parser -f Dockerfile
