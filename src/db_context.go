package db_context

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	"github.com/doll-blues-vachello/vkma2022-finale/src/repository"
)

type DbContext struct {
	db          DB
	photos      repository.PhotoRepository
	albumAccess repository.AlbumAccessRepository
}

func (ctx *DbContext) Open(dbPath string) {
	ctx.db, err = sql.Open("sqlite3", dbPath)

	if err != nil {
		panic(err)
	}

	ctx.photos = repository.PhotoRepository{
		photosTableName: "Photos",
		usersTableName:  "Users",
	}

	ctx.albumAccess = repository.AlbumAccessRepository{
		&ctx.db,
		albumAccessTableName: "AlbumAccess",
	}
}

func (ctx *DbContext) Close() {
	ctx.db.Close()
}
