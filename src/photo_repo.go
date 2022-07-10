package main

import (
	"fmt"
	"time"
)

type PhotoRepo Repository

func (repo PhotoRepo) Add(photo Photo) int64 {
	var query = fmt.Sprintf(`
		INSERT INTO
			Photos(AlbumID, UploaderID, Url)
		VALUES(%d, %d, '%s')
	`, photo.AlbumID, photo.UploaderID, photo.Url)

	var res, _ = repo.DB.Exec(query)
	var id, _ = res.LastInsertId()

	return id
}

func (repo PhotoRepo) GetByID(id int64) (Photo, error) {
	var photo Photo
	var created time.Time
	var query = fmt.Sprintf("SELECT * FROM Photos WHERE ID = %d", id)

	var e = repo.DB.QueryRow(query).Scan(
		&photo.ID,
		&photo.AlbumID,
		&photo.UploaderID,
		&photo.Url,
		&created)

	if e != nil {
		return Photo{}, e
	}

	photo.Created = created.Unix()
	return photo, nil
}

func (repo PhotoRepo) Update(photo Photo) {
	query := fmt.Sprintf(`
		UPDATE
			Photos
		SET
			AlbumID = %d,
		WHERE
			ID = %d
	`, photo.AlbumID, photo.ID)

	repo.DB.Exec(query)
}

func (repo PhotoRepo) DeleteByID(id int64) {
	query := fmt.Sprintf("DELETE FROM Photos WHERE ID = %d", id)
	repo.DB.Exec(query)
}
