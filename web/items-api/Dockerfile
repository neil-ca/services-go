FROM golang:1.15.5-alpine as builder
COPY go.mod go.sum /go/src/github.com/Neil-Uli/items-api/

WORKDIR /go/src/github.com/Neil-Uli/items-api

RUN go mod download

COPY . /go/src/github.com/Neil-Uli/items-api
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/items-api github.com/Neil-Uli/items-api

FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /go/src/github.com/Neil-Uli/build/items-api /usr/bin/items-api

EXPOSE 8080 8080
ENTRYPOINT ["/usr/bin/items-api"]