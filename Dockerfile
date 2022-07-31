FROM 899991151204.dkr.ecr.us-east-1.amazonaws.com/golang:amazon_linux

USER root

RUN yum update -y
RUN yum install -y gcc

RUN go install github.com/go-delve/delve/cmd/dlv@master

WORKDIR /go/src/github.com/ysed/myapp
COPY . .
RUN make app.build

EXPOSE 30123 8080

COPY entrypoint.sh .
ENTRYPOINT ["sh", "entrypoint.sh"]