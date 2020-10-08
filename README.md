# auth-server-go #

[![Gitpod ready-to-code](https://img.shields.io/badge/Gitpod-ready--to--code-blue?logo=gitpod)](https://gitpod.io/#https://bitbucket.org/Eldius/auth-server-go)

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
go run user add -u "eldius" -W "MyStrongPass@1" -a
```

```shell
curl -i localhost:8000/login -d '{"user": "eldius", "pass": "MyStrongPass@1"}'
# returns something like this
# {"token":"header.payload.sign"}
```

```shell
# the "header.payload.sign" value is acquired in the last snippet call
curl -i localhost:8000/admin -H "Authorization: Bearer header.payload.sign"
```

```shell
curl -i localhost:8000/admin -H "Authorization: Bearer $( curl --fail localhost:8000/login -d '{"user": "eldius", "pass": "pass@1"}' 2>/dev/null | jq -r '. | .token' )" 2>/dev/null
```