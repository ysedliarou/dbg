FROM golang:latest

USER root

RUN go install github.com/go-delve/delve/cmd/dlv@master

WORKDIR /go/src/github.com/ysed/myapp
COPY . .
RUN make app.build

EXPOSE 30123 8080

COPY entrypoint.sh .
ENTRYPOINT ["sh", "entrypoint.sh"]