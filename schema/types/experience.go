package types

import "time"

type Experience struct {
	ID          uint      `json:"id,omitempty"`
	Company     string    `json:"company"`
	Position    string    `json:"position"`
	Thumbmail   string    `json:"thumbnail"`
	Url         string    `json:"url"`
	Location    string    `json:"location"`
	Start_at    time.Time `json:"start_at"`
	End_at      time.Time `json:"end_at"`
	Description string    `json:"description"`
}

func (e Experience) Save() {
	// TODO: save model
}
