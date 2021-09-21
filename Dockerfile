# lightweight container for go
FROM golang:1.16

COPY go.mod go.sum /go/src/github.com/febrielven/go-crudapi/
WORKDIR /go/src/github.com/febrielven/go-crudapi
RUN go mod download
COPY .  /go/src/github.com/febrielven/go-crudapi
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/go-crudapi github.com/febrielven/go-crudapi


