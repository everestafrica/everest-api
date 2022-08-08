package models

import "time"

type News struct {
	Title       string `json:"title"`
	Img         string `json:"img"`
	Author      string `json:"author"`
	Link        string `json:"link"`
	Description string `json:"description"`
	Date        string `json:"date"`
}

type T struct {
	Author      string    `json:"author"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Url         string    `json:"url"`
	UrlToImage  string    `json:"urlToImage"`
	PublishedAt time.Time `json:"publishedAt"`
	Content     string    `json:"content"`
}
