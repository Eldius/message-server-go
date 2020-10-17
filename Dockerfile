FROM golang:alpine AS builder
WORKDIR /app
COPY . /app
RUN apk add build-base
RUN go test ./... -cover
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -v -o bin/message-server-linux-amd64 -a -ldflags '-extldflags "-static"'  -ldflags "-X 'github.com/Eldius/message-server-go/config.buildDate=$(date +"%Y-%m-%dT%H:%M:%S%:z")' -X 'github.com/Eldius/message-server-go/config.version=123' -X 'github.com/Eldius/message-server-go/config.branchName=ABC'" .

FROM alpine:3.12

RUN apk add sqlite

WORKDIR /app

COPY --from=builder /app/bin/message-server-linux-amd64 /app
RUN mv /app/message-server-linux-amd64 /app/message-server

EXPOSE 8000

ENTRYPOINT [ "/app/message-server", "start" ]
