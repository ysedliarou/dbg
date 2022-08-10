SRC_ROOT ?= $(realpath .)

COMMIT = $(shell git log --pretty=format:'%h' -n 1)
APP_DOCKER_TAG = ysedcoupa/dbg:$(COMMIT)

default: build

fmt:
	go fmt ./...

build: prepare
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -gcflags "all=-N -l" -o ./bin/app

prepare: clean fmt
	 go mod tidy && go mod vendor

clean:
	rm -rf vendor
	rm -rf bin

test:
	go test -coverprofile=coverage.out

docker.build:
	docker image build -t $(APP_DOCKER_TAG) $(SRC_ROOT)

docker.run:
	docker run -p 8080:8080 -p 30123:30123 $(APP_DOCKER_TAG)

docker.push:
	docker image push $(APP_DOCKER_TAG)

d: docker.build


