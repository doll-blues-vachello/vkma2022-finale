package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

type Id uint64

type AccessLevel uint8

const (
	READ   AccessLevel = 1
	CREATE AccessLevel = 2
	DELETE AccessLevel = 4
)

type AlbumAccess struct {
	albumId     Id
	userId      Id
	accessLevel AccessLevel
}

func main() {
	var r = gin.Default()
	// var dbContext DbContext
	// dbContext.Open("storage.db")
	// defer dbContext.Close()

	var db, err = sql.Open("sqlite3", "storage.db")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	r.GET("/albumAccess/add", func(c *gin.Context) {
		Add(db, AlbumAccess{1, 1, 7})
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	r.GET("/albumAccess/get", func(c *gin.Context) {
		var aa = GetByAlbumId(db, 1)

		c.JSON(http.StatusOK, gin.H{
			"list": aa,
		})
	})

	r.Run()
}

func Add(db *sql.DB, albumAccess AlbumAccess) {
	var query = fmt.Sprintf(
		"INSERT INTO %s(albumId, userId, accessLevel) VALUES(%d, %d, %d)",
		"AlbumAccess",
		albumAccess.albumId,
		albumAccess.userId,
		albumAccess.accessLevel,
	)

	var _, err = db.Exec(query)

	if err != nil {
		panic(err)
	}
}

func GetByAlbumId(db *sql.DB, albumId Id) []AlbumAccess {
	var albumAccessList []AlbumAccess
	var query = fmt.Sprintf(
		"SELECT * FROM %s WHERE albumId = %d",
		"AlbumAccess",
		albumId,
	)

	var rows, err = db.Query(query)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var albumAccess AlbumAccess

		rows.Scan(
			&albumAccess.albumId,
			&albumAccess.userId,
			&albumAccess.accessLevel,
		)

		albumAccessList = append(albumAccessList, albumAccess)
	}

	return albumAccessList
}
