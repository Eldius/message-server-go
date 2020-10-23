FROM golang:1.15.3-alpine3.12 AS builder
WORKDIR /app
COPY . /app
RUN apk add build-base
RUN go test ./... -cover
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -v -o bin/message-server-linux-amd64 -a -ldflags '-extldflags "-static"'  -ldflags "-X 'github.com/Eldius/message-server-go/config.buildDate=$(date +"%Y-%m-%dT%H:%M:%S%:z")' -X 'github.com/Eldius/message-server-go/config.version=123' -X 'github.com/Eldius/message-server-go/config.branchName=ABC'" .

FROM alpine:3.12

WORKDIR /app

COPY --from=builder /app/bin/message-server-linux-amd64 /app
RUN apk add --no-cache sqlite && \
    addgroup -S messengergroup && \
    adduser -S messenger -G messengergroup && \
    mv /app/message-server-linux-amd64 /app/message-server && \
    chown messenger. -R /app

EXPOSE 8000

USER messenger

ENTRYPOINT [ "/app/message-server", "start" ]
