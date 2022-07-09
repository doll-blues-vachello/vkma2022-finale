package model

import (
	"time"
	"vkma2022-finale/types"
)

type Album struct {
	id      types.Id
	userId  types.Id
	created time.Time
}
