# lightweight container for go
FROM golang:1.16 as builder

COPY go.mod go.sum /go/src/github.com/febrielven/go-crudapi/
WORKDIR /go/src/github.com/febrielven/go-crudapi
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/go-crudapi github.com/febrielven/go-crudapi


# Start a new stage from alpine
FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /go/src/github.com/febrielven/go-crudapi/build/go-crudapi /usr/bin/go-crudapi


EXPOSE 8081 8081
ENTRYPOINT ["/usr/bin/go-crudapi"]
