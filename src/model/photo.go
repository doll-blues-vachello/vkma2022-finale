package model

import (
	"time"
	"vkma2022-finale/types"
)

type Photo struct {
	albumId    types.Id
	uploaderId types.Id
	path       string
	created    time.Time
}
