package repository

import (
	"github.com/doll-blues-vachello/vkma2022-finale/src/model"
	"github.com/doll-blues-vachello/vkma2022-finale/src/types"
)

type PhotoRepository struct {
	photosTableName string
	usersTableName  string
}

func GetById(id types.Id) model.Photo {
	//
}
