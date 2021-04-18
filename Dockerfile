FROM golang:1.15.3-alpine3.12 AS builder
WORKDIR /app
COPY . /app
RUN apk add build-base git
RUN go test ./... -cover
RUN CGO_ENABLED=1 \
    go build \
    -v \
    -o bin/message-server-go \
    -a \
    -ldflags '-extldflags "-static"' \
    -ldflags "-X 'github.com/eldius/message-server-go/config.buildDate=$(date +"%Y-%m-%dT%H:%M:%S%:z")' -X 'github.com/eldius/message-server-go/config.version=$(git rev-parse --abbrev-ref HEAD)' -X 'github.com/eldius/message-server-go/config.branchName=$(git branch --show-current)'" \
        .

FROM alpine:3.12

WORKDIR /app

COPY --from=builder /app/bin/message-server-go /app
RUN apk add --no-cache sqlite && \
    addgroup -S messengergroup && \
    adduser -S messenger -G messengergroup && \
    mv /app/message-server-go /app/message-server-go && \
    chown messenger. -R /app

EXPOSE 8000

USER messenger

ENTRYPOINT [ "/app/message-server-go", "start" ]
