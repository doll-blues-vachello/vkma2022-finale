package model

import "vkma2022-finale/types"

type AlbumAccess struct {
	albumId     types.Id
	userId      types.Id
	accessLevel types.AccessLevel
}
