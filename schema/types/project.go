package types

import "time"

type Project struct {
	ID          uint `json:"id,omitempty"`
	Owner       string
	Repo        string
	Thumbmail   string
	Description string
	Url         string
	Updated_at  time.Time
	Description string
}

func (p Project) Save() {
	// TODO: load required info from github
	// TODO: save model
}
