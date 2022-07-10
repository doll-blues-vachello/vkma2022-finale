EXEC=vkma2022-finale

server:
	rm -f bin/${EXEC}
	go build -o bin/${EXEC} src/*.go

database:
	sqlite3 storage.db < init/schema.sql

