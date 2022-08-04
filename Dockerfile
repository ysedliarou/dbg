FROM golang:latest

USER root

RUN go install github.com/go-delve/delve/cmd/dlv@master

WORKDIR /go/src/github.com/ysed/myapp
COPY . .
RUN make

EXPOSE 8080 30123

ENTRYPOINT ["sh", "entrypoint.sh"]