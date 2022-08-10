FROM golang:1.18 as builder

COPY go.mod go.sum /go/src/github.com/rafaelmgr12/Mark-URL/
WORKDIR /go/src/github.com/rafaelmgr12/Mark-URL
RUN go mod download
COPY . /go/src/github.com/rafaelmgr12/Mark-URL
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/mark github.com/rafaelmgr12/Mark-URL/cmd/main


FROM alpine

RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /go/src/github.com/rafaelmgr12/Mark-URL/build/mark /usr/bin/mark

EXPOSE 8080 8080

ENTRYPOINT ["/usr/bin/mark"]
