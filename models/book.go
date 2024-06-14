package models

type Book struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	ISBN    string `json:"isbn"`
	Author  string `json:"author"`
	Release int    `json:"release"`
}
