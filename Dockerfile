FROM alpine:3.12

WORKDIR /app

COPY ./bin/message-server-linux-amd64 /app/message-server

EXPOSE 8000

ENTRYPOINT [ "/app/message-server", "start" ]
