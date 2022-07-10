package main

type Album struct {
	ID      int64  `json:"id"`
	UserID  int64  `json:"user_id"`
	Title   string `json:"title"`
	Created int64  `json:"created"`
}
