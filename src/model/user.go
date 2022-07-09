package model

import (
	"time"
	"vkma2022-finale/types"
)

type User struct {
	id        types.Id
	firstName string
	lastName  string
	vkId      types.Id
	created   time.Time
}
