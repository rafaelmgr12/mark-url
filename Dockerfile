FROM golang:alpine AS build-env
ENV GOPATH /go
WORKDIR /go/src/github.com/rafaelmgr12/Mark-URL
COPY . /go/src/github.com/rafaelmgr12/Mark-URL
RUN cd /go/src/github.com/rafaelmgr12/Mark-URL && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/mark github.com/rafaelmgr12/Mark-URL/cmd/main

FROM alpine

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk*
WORKDIR /app
COPY --from=build-env /go/src/github.com/rafaelmgr12/Mark-URL/build/mark /app
COPY .env /app

EXPOSE 8080

ENTRYPOINT [ "./mark" ]
