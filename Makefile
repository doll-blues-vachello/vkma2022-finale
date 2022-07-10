DB=storage.db
EXEC=vkma2022-finale
DEFAULT_PORT=4567


database:
	sqlite3 ${DB} < init/schema.sql

server:
	rm -f bin/${EXEC}
	go build -o bin/${EXEC} src/*.go

