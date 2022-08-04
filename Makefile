SRCROOT ?= $(realpath .)

APP_DOCKER_LABEL = dbg
APP_DOCKER_TAG = 0.1

default: build

build: prepare
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -gcflags "all=-N -l" -o ./bin/app

prepare: clean
	 go mod tidy && go mod vendor

clean:
	rm -rf vendor
	rm -rf bin

test:
	go test -coverprofile=coverage.out

docker.build:
	docker build \
		-t $(APP_DOCKER_LABEL):$(APP_DOCKER_TAG) $(SRCROOT)

docker.run:
	docker run -p 8080:8080 -p 30123:30123 $(APP_DOCKER_LABEL):$(APP_DOCKER_TAG)

d: docker.build


