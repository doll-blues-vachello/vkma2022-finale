package main

import (
	"fmt"
	"time"
)

type AlbumRepo Repository

func (repo AlbumRepo) Add(album Album) int64 {
	var query = fmt.Sprintf(`
		INSERT INTO
			Albums(UserID, Title)
		VALUES(%d, '%s')
	`, album.UserID, album.Title)

	var res, _ = repo.DB.Exec(query)
	var id, _ = res.LastInsertId()

	return id
}

func (repo AlbumRepo) GetByID(id int64) (Album, error) {
	var album Album
	var created time.Time
	var query = fmt.Sprintf("SELECT * FROM Albums WHERE ID = %d", id)

	var e = repo.DB.QueryRow(query).Scan(
		&album.ID,
		&album.UserID,
		&album.Title,
		&created)

	if e != nil {
		return Album{}, e
	}

	album.Created = created.Unix()
	return album, nil
}
