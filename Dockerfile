FROM alpine:3.12

WORKDIR /app

COPY static/bundle /app/static/bundle
COPY static/css /app/static/css
COPY static/assets /app/static/assets
COPY templates /app/templates
COPY bin/auth-server /app/auth-server

EXPOSE 8000

ENTRYPOINT [ "/app/auth-server", "start" ]
