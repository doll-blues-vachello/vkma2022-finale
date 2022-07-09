package repository

import (
	"fmt"

	"github.com/doll-blues-vachello/vkma2022-finale/src/types"
)

type AlbumAccessRepository struct {
	db                   DB
	albumAccessTableName string
	//
}

func (repo *AlbumAccessRepository) Add(albumAccess AlbumAccess) types.Id {
	var query = fmt.Sprintf(
		"INSERT INTO %s(photoId, userId, accessLevel) VALUES(%d, %d, %d)",
		repo.albumAccessTableName,
		albumAccess.albumId,
		albumAccess.userId,
		albumAccess.accessLevel,
	)

	res, _ = repo.db.Exec(query)
	return res.LastInsertId()
}

func (repo *AlbumAccessRepository) GetByAlbumId(albumId types.Id) AlbumAccess {
	var albumAccess AlbumAccess
	var query = fmt.Sprintf(
		"SELECT * FROM %s WHERE albumId = %d",
		repo.albumAccessTableName,
		albumId,
	)

	row, err = repo.db.QueryRow(query)

	if err != nil {
		panic(err)
	}

	row.Scan(
		&albumAccess.photoId,
		&albumAccess.userId,
		&albumAccess.accessLevel,
	)

	return albumAccess
}
