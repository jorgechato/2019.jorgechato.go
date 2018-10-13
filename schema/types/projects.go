package types

import "time"

type Project struct {
	ID          uint      `json:"id,omitempty"`
	Owner       string    `json:"-"`
	Repo        string    `json:"-"`
	Thumbmail   string    `json:"thumbnail"`
	Description string    `json:"description"`
	Url         string    `json:"url"`
	Updated_at  time.Time `json:"updated_at"`
}

func (p Project) Save() {
	// TODO: load required info from github Updated_at, Description, Url
	// TODO: save model
}
