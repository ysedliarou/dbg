FROM golang:latest

USER root

RUN go install github.com/go-delve/delve/cmd/dlv@master

WORKDIR /go/src/github.com/ysed/myapp
COPY . .
RUN make app.build

EXPOSE 8080
EXPOSE 30123

ENTRYPOINT ["sh", "entrypoint.sh"]