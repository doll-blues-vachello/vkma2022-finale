BIN=vkma2022-finale

server:
	go build -o bin/${BIN} src/main.go

database:
	sqlite3 storage.db < init/schema.sql

