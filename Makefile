
.EXPORT_ALL_VARIABLES:

#MESSENGER_APP_LOG_FORMAT=
#MESSENGER_APP_DATABASE_URL=$(HOME)/messenger.db
#MESSENGER_APP_DATABASE_ENGINE=sqlite3
MESSENGER_APP_DATABASE_URL=msg_usr:msg_pass@tcp(127.0.0.1:3306)/messages?charset=utf8mb4&parseTime=True&loc=Local
MESSENGER_APP_DATABASE_ENGINE=mysql
MESSENGER_APP_DATABASE_LOG=true
MESSENGER_AUTH_USER_PATTERN=^[a-zA-Z0-9\\._-]*$
MESSENGER_AUTH_PASS_PATTERN=.+
MESSENGER_AUTH_JWT_SECRET="MyStr@ngeP2ss"
MESSENGER_AUTH_USER_DEFAULT_ACTIVE=true

clean:
	-rm *.log*
	-rm **/*.log*
	-rm *.db*
	-rm **/*.db*

startwsgbd:
	go run main.go start

start:
	go run main.go start

test: clean
	go test ./... -cover


createtestusers:
	go run main.go user add -n "Eldius" -u "eldius" -W "MyStrongAdminPass@1" -a
	go run main.go user add -n "Test User" -u "testUser" -W "MyStrongPass@1"
	go run main.go user add -n "Ciclano da Silva" -u "test1" -W "MyStrongPass@1"
	go run main.go user add -n "Fulano Santos" -u "test2" -W "MyStrongPass@2"

sendmessages:
	$(eval ELDIUS_BEARER = $(shell curl --fail http://localhost:8000/login -d '{"user": "eldius", "pass": "MyStrongAdminPass@1"}' 2>/dev/null | jq -r '. | .token'))
	$(eval TEST_BEARER = $(shell curl --fail http://localhost:8000/login -d '{"user": "testUser", "pass": "MyStrongPass@1"}' 2>/dev/null | jq -r '. | .token'))
	$(eval TEST1_BEARER = $(shell curl --fail http://localhost:8000/login -d '{"user": "test1", "pass": "MyStrongPass@1"}' 2>/dev/null | jq -r '. | .token'))
	$(eval TEST2_BEARER = $(shell curl --fail http://localhost:8000/login -d '{"user": "test2", "pass": "MyStrongPass@2"}' 2>/dev/null | jq -r '. | .token'))
	@echo ELDIUS_BEARER: $(ELDIUS_BEARER)
	@echo TEST_BEARER: $(TEST_BEARER)
	@echo TEST1_BEARER: $(TEST1_BEARER)
	@echo TEST2_BEARER: $(TEST2_BEARER)
	curl -i -XPOST http://localhost:8000/message -H "Authorization: Bearer $(ELDIUS_BEARER)" -d '{"to": "eldius","msg": "My new message 01!"}' 2>/dev/null
	sleep 2
	curl -i -XPOST http://localhost:8000/message -H "Authorization: Bearer $(ELDIUS_BEARER)" -d '{"to": "test1","msg": "My new message 02!"}' 2>/dev/null
	sleep 2
	curl -i -XPOST http://localhost:8000/message -H "Authorization: Bearer $(ELDIUS_BEARER)" -d '{"to": "test2","msg": "My new message 03!"}' 2>/dev/null
	sleep 2
	curl -i -XPOST http://localhost:8000/message -H "Authorization: Bearer $(ELDIUS_BEARER)" -d '{"to": "testUser","msg": "My new message 04!"}' 2>/dev/null
	sleep 2
	curl -i -XPOST http://localhost:8000/message -H "Authorization: Bearer $(ELDIUS_BEARER)" -d '{"to": "testUser","msg": "My new message 05!"}' 2>/dev/null
	sleep 2
	curl -i -XPOST http://localhost:8000/message -H "Authorization: Bearer $(ELDIUS_BEARER)" -d '{"to": "testUser","msg": "My new message 06!"}' 2>/dev/null
	sleep 2
	curl -i -XPOST http://localhost:8000/message -H "Authorization: Bearer $(ELDIUS_BEARER)" -d '{"to": "testUser","msg": "My new message 07!"}' 2>/dev/null
	sleep 2
	curl -i -XPOST http://localhost:8000/message -H "Authorization: Bearer $(TEST_BEARER)" -d '{"to": "eldius","msg": "My new message 08!"}' 2>/dev/null
	sleep 2
	curl -i -XPOST http://localhost:8000/message -H "Authorization: Bearer $(TEST1_BEARER)" -d '{"to": "eldius","msg": "My new message 09!"}' 2>/dev/null
	sleep 2
	curl -i -XPOST http://localhost:8000/message -H "Authorization: Bearer $(TEST2_BEARER)" -d '{"to": "eldius","msg": "My new message 10!"}' 2>/dev/null
	sleep 2
	curl -i -XPOST http://localhost:8000/message -H "Authorization: Bearer $(TEST_BEARER)" -d '{"to": "eldius","msg": "My new message 11!"}' 2>/dev/null
	sleep 2
	curl -i -XPOST http://localhost:8000/message -H "Authorization: Bearer $(TEST_BEARER)" -d '{"to": "eldius","msg": "My new message 12!"}' 2>/dev/null

build: clean
	$(eval GIT_COMMIT_HASH = $(shell git rev-parse --short HEAD))
	$(eval GIT_BRANCH = $(shell git branch --show-current))
	$(eval BUILD_DATE = $(shell date +"%Y-%m-%dT%H:%M:%S%:z"))
	git branch --show-current
	CGO_ENABLED=0 \
	go build \
		-v \
		-o bin/message-server-go \
		-a \
		-ldflags '-extldflags "-static"' \
		-ldflags "-X 'github.com/eldius/message-server-go/config.buildDate=$(BUILD_DATE)' -X 'github.com/eldius/message-server-go/config.version=$(GIT_COMMIT_HASH)' -X 'github.com/eldius/message-server-go/config.branchName=$(GIT_BRANCH)'" \
		.

builddocker:
	$(eval GIT_COMMIT_HASH = $(shell git rev-parse --short HEAD))
	$(eval GIT_BRANCH = $(shell git branch --show-current))
	docker build \
		--no-cache \
		-t eldius/message-server-go:$(GIT_BRANCH)-$(GIT_COMMIT_HASH) \
		.

startmariadb:
	-docker kill db
	docker run \
		--name db \
		-e MYSQL_ROOT_PASSWORD=my-secret-pw \
		-e MYSQL_DATABASE=messages \
		-e MYSQL_USER=msg_usr \
		-e MYSQL_PASSWORD=msg_pass \
		-p 3306:3306 \
		-d \
		--rm \
		mariadb:latest

dbconsole:
	docker exec -it db mysql -umsg_usr -pmsg_pass

startdocker:
	$(eval GIT_COMMIT_HASH = $(shell git rev-parse --short HEAD))
	$(eval GIT_BRANCH = $(shell git branch --show-current))
	docker run \
		-it \
		--name messenger \
		--rm \
		-e MESSENGER_APP_DATABASE_URL=messenger.db \
		-e MESSENGER_APP_DATABASE_ENGINE=sqlite3 \
		-e MESSENGER_APP_DATABASE_LOG=true \
		-e MESSENGER_AUTH_USER_PATTERN=^[a-zA-Z0-9\\._-]*$ \
		-e MESSENGER_AUTH_PASS_PATTERN=.+ \
		-e MESSENGER_AUTH_JWT_SECRET="MyStr@ngeP2ss" \
		-e MESSENGER_AUTH_USER_DEFAULT_ACTIVE=true \
		-p 8000:8000 \
		-m 32m \
		--cpus=0.1 \
		eldius/message-server-go:$(GIT_BRANCH)-$(GIT_COMMIT_HASH)

startfront:
	cd static ; yarn start
