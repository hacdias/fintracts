.PHONY: clean docker

clean:
	rm -f validator

docker:
	docker build .. -t validator -f Dockerfile

build:
	go build -o validator cmd/*.go
