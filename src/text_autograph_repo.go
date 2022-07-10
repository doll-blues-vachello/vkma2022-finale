package main

import (
	"fmt"
	"time"
)

type TextAutographRepo Repository

func (repo TextAutographRepo) Add(textAutograph TextAutograph) int64 {
	var query = fmt.Sprintf(`
		INSERT INTO
			TextAutographs(PhotoID, AuthorID, Text)
		VALUES(%d, %d, '%s')
	`, textAutograph.PhotoID, textAutograph.AuthorID, textAutograph.Text)

	var res, _ = repo.DB.Exec(query)
	var id, _ = res.LastInsertId()

	return id
}

func (repo TextAutographRepo) GetByID(id int64) (TextAutograph, error) {
	var textAutograph TextAutograph
	var created time.Time
	var query = fmt.Sprintf("SELECT * FROM TextAutographs WHERE ID = %d", id)

	var e = repo.DB.QueryRow(query).Scan(
		&textAutograph.ID,
		&textAutograph.PhotoID,
		&textAutograph.AuthorID,
		&textAutograph.Text,
		&created)

	if e != nil {
		return TextAutograph{}, e
	}

	textAutograph.Created = created.Unix()
	return textAutograph, nil
}

func (repo TextAutographRepo) Update(autograph TextAutograph) {
	query := fmt.Sprintf(`
		UPDATE
			TextAutographs
		SET
			PhotoID  = %d,
			AuthorID = %d,
			Text     = '%s'
		WHERE
			ID = %d
	`, autograph.PhotoID, autograph.AuthorID, autograph.Text, autograph.ID)

	repo.DB.Exec(query)
}

func (repo TextAutographRepo) DeleteByID(id int64) {
	query := fmt.Sprintf("DELETE FROM TextAutographs WHERE ID = %d", id)
	repo.DB.Exec(query)
}
