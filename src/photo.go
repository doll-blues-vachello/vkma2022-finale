package main

type Photo struct {
	ID         int64  `json:"id"`
	AlbumID    int64  `json:"album_id"`
	UploaderID int64  `json:"uploader_id"`
	Url        string `json:"url"`
	Created    int64  `json:"created"`
}
