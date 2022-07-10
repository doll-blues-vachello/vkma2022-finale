package main

import (
	"fmt"
	"time"
)

type ImageAutographRepo Repository

func (repo ImageAutographRepo) Add(imageAutograph ImageAutograph) int64 {
	var query = fmt.Sprintf(`
		INSERT INTO
			ImageAutographs(PhotoID, AuthorID, Url)
		VALUES(%d, %d, '%s')
	`, imageAutograph.PhotoID, imageAutograph.AuthorID, imageAutograph.Url)

	var res, _ = repo.DB.Exec(query)
	var id, _ = res.LastInsertId()

	return id
}

func (repo ImageAutographRepo) GetByID(id int64) (ImageAutograph, error) {
	var imageAutograph ImageAutograph
	var created time.Time
	var query = fmt.Sprintf("SELECT * FROM ImageAutographs WHERE ID = %d", id)

	var e = repo.DB.QueryRow(query).Scan(
		&imageAutograph.ID,
		&imageAutograph.PhotoID,
		&imageAutograph.AuthorID,
		&imageAutograph.Url,
		&created)

	if e != nil {
		return ImageAutograph{}, e
	}

	imageAutograph.Created = created.Unix()
	return imageAutograph, nil
}
