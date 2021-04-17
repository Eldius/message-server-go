
.EXPORT_ALL_VARIABLES:

#MESSENGER_APP_LOG_FORMAT=
MESSENGER_APP_DATABASE_URL=messenger.db
MESSENGER_APP_DATABASE_ENGINE=sqlite3
MESSENGER_APP_DATABASE_LOG=true
MESSENGER_AUTH_USER_PATTERN="^[a-zA-Z0-9\\._-]*$"
MESSENGER_AUTH_PASS_PATTERN="^[a-zA-Z0-9\\._-]*$"
MESSENGER_AUTH_JWT_SECRET="MyStr@ngeP2ss"
MESSENGER_AUTH_USER_DEFAULT_ACTIVE=true

clean:
	-rm *.log*
	-rm **/*.log*
	-rm *.db*
	-rm **/*.db*

start:
	go run main.go start

test: clean
	go test ./... -cover


createtestusers:
	go run main.go user add -u "eldius" -W "MyStrongAdminPass@1" -a
	go run main.go user add -u "testUser" -W "MyStrongPass@1"
	go run main.go user add -u "test1" -W "MyStrongPass@1"
	go run main.go user add -u "test2" -W "MyStrongPass@2"
