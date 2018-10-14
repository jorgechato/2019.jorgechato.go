package types

import "time"

type Article struct {
	ID           uint      `json:"id,omitempty"`
	Created_at   time.Time `json:"created_at"`
	Title        string    `json:"title"`
	Tags         []Tag     `json:"tags"`
	Slug         string    `json:"slug"`
	Content      string    `json:"content"`
	Published_at time.Time `json:"published_at"`
	Public       bool      `json:"public"`
	Thumbmail    string    `json:"thumbnail"`
}

type Tag struct {
	ID   uint   `json:"id,omitempty"`
	Name string `json:"tag"`
}

func (a Article) Save() {
	a.Slug = slugify(a.Title)
	// TODO: save model
}
