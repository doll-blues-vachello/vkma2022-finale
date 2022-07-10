package main

type ImageAutographs struct {
	ID       int64  `json:"id"`
	PhotoID  int64  `json:"photo_id"`
	AuthorID int64  `json:"author_id"`
	Url      string `json:"url"`
	Created  int64  `json:"created"`
}
