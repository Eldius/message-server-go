# message-server-go #

![Go](https://github.com/eldius/message-server-go/workflows/Go/badge.svg)
[![Gitpod ready-to-code](https://img.shields.io/badge/Gitpod-ready--to--code-blue?logo=gitpod)](https://gitpod.io/#https://bitbucket.org/Eldius/message-server-go)

## config ##

### env vars ###

- app config:
  - MESSENGER_APP_LOG_FORMAT: log format, `text` for normal logs and anything else for JSON logs
  - MESSENGER_APP_DATABASE_URL: app database connection URL
  - MESSENGER_APP_DATABASE_ENGINE: app database engine (sqlite3 is the only one tested until now)
  - MESSENGER_APP_DATABASE_LOG: app log database queries (true|false)?
- authentication config:
  - MESSENGER_AUTH_USER_PATTERN: pattern for user validation
  - MESSENGER_AUTH_PASS_PATTERN: pattern for password validation
  - MESSENGER_AUTH_JWT_SECRET: secret used to sign JWT tokens
  - MESSENGER_AUTH_USER_DEFAULT_ACTIVE: default state for new users (true|false)
- CORS interceptor config:
  - MESSENGER_CORS_ALLOW_METHODS: HTTP request methods allowed
  - MESSENGER_CORS_ALLOW_HEADERS: HTTP headers allowed
  - MESSENGER_CORS_ALLOW_ORIGIN: origin domains allowed

## dev ##

### start database using docker ##

```shell
docker run \
    --name test-db \
    -e MYSQL_ROOT_PASSWORD=my-secret-pw \
    -e MYSQL_DATABASE=auth_db \
    -e MYSQL_USER=auth_usr \
    -e MYSQL_PASSWORD=auth_pass \
    -p 3306:3306 \
    -d \
    --rm \
    mariadb:latest

mysql -u auth_usr -p -h 127.0.0.1

```

### env vars ###

```shell
export APP_DATABASE_URL='auth_usr:auth_pass@tcp(127.0.0.1:3306)/auth_db'
export APP_DATABASE_ENGINE='mysql'

```

### select users ###

```shell
mysql -u auth_usr -h 127.0.0.1 -p'auth_pass' < repository/testqueries/auth-server.sql
```

```shell
golangci-lint run && echo "" && go test ./... -coverprofile=coverage.out && echo "" && go tool cover -func=coverage.out
```

### some tests ###

```shell
go run user add -u "eldius" -W "MyStrongAdminPass@1" -a
```

```shell
curl -i localhost:8000/login -d '{"user": "eldius", "pass": "MyStrongAdminPass@1"}'
# returns something like this
# {"token":"header.payload.sign"}
```

```shell
# the "header.payload.sign" value is acquired in the last snippet call
curl -i localhost:8000/admin -H "Authorization: Bearer header.payload.sign"
```

```shell
curl -i localhost:8000/admin -H "Authorization: Bearer $( curl --fail localhost:8000/login -d '{"user": "eldius", "pass": "MyStrongAdminPass@1"}' 2>/dev/null | jq -r '. | .token' )" 2>/dev/null
```

```shell
#send message
curl -i -XPOST localhost:8000/message -H "Authorization: Bearer $( curl --fail localhost:8000/login -d '{"user": "testUser", "pass": "MyStrongPass@1"}' 2>/dev/null | jq -r '. | .token' )" -d '{"to": "eldius","msg": "My new message 01!"}' 2>/dev/null
```

```shell
#fetch messages
curl -i -XGET localhost:8000/message -H "Authorization: Bearer $( curl --fail localhost:8000/login -d '{"user": "eldius", "pass": "MyStrongAdminPass@1"}' 2>/dev/null | jq -r '. | .token' )" 2>/dev/null
```

```shell
curl -i -XPOST http://localhost:8000/user -H "Authorization: Bearer $( curl --fail localhost:8000/login -d '{"user": "eldius", "pass": "MyStrongAdminPass@1"}' 2>/dev/null | jq -r '. | .token' )" \
-d '{
  "user": "fulanus",
  "pass": "pass",
  "name": "Fulanus",
  "active": true,
  "admin": true
}'
```