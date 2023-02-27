package models

type News struct {
	Base
	Source      string `json:"source"`
	Title       string `json:"title"`
	Img         string `json:"img"`
	Author      string `json:"author"`
	Link        string `json:"link"`
	Description string `json:"description"`
	Date        string `json:"date"`
	Category    string `json:"category"`
}
