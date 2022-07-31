SRCROOT ?= $(realpath .)


VERSION = $$(git name-rev --tags --name-only $$(git rev-parse HEAD))
REVISION = $$(git rev-parse --short HEAD)

APP_DOCKER_LABEL = dbg
# APP_DOCKER_TAG = $(VERSION).$(REVISION)
APP_DOCKER_TAG = 0.1

d: docker.build

docker.build:
	docker build \
		-t $(APP_DOCKER_LABEL):$(APP_DOCKER_TAG) $(SRCROOT)

app.build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -gcflags "all=-N -l" -o ./app
