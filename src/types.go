package main

import "database/sql"

type Repository struct {
	DB *sql.DB
}

type NotFoundError error
