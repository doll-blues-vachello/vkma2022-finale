package model

import (
	"time"
	"vkma2022-finale/types"
)

type ImageAutographs struct {
	id       types.Id
	photoId  types.Id
	authorId types.Id
	content  string
	created  time.Time
}
