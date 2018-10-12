package types

import "time"

type Misc struct {
	ID           uint      `json:"id,omitempty"`
	Created_at   time.Time `json:"created_at"`
	Title        string    `json:"title"`
	Url          string    `json:"url"`
	Slug         string    `json:"slug"`
	Content      string    `json:"content"`
	Published_at time.Time `json:"published_at"`
	Public       bool      `json:"public"`
	Thumbmail    string    `json:"thumbnail"`
}

func (m Misc) Save() {
	m.Slug = slugify(m.Title)
	// TODO: save model
}
