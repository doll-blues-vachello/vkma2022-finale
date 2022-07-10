package main

type AlbumAccess struct {
	AlbumID     int64 `json:"album_id"`
	UserID      int64 `json:"user_id"`
	AccessLevel int   `json:"access_level"`
}
