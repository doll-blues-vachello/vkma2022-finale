package main

type TextAutograph struct {
	ID       int64  `json:"id"`
	PhotoID  int64  `json:"photo_id"`
	AuthorID int64  `json:"author_id"`
	Text     string `json:"text"`
	Created  int64  `json:"created"`
}
