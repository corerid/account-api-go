FROM golang:1.12.9-alpine3.10
WORKDIR /go
COPY . .

RUN go build /go/src/go-api/main.go /go/src/go-api/handlers.go /go/src/go-api/routes.go /go/src/go-api/structs.go