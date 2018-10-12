package types

import "time"

type Event struct {
	ID          uint      `json:"id,omitempty"`
	Title       string    `json:"title"`
	Thumbmail   string    `json:"thumbnail"`
	Url         string    `json:"url"`
	Location    string    `json:"location"`
	Start_at    time.Time `json:"start_at"`
	End_at      time.Time `json:"end_at"`
	Description string    `json:"description"`
}

func (e Event) Save() {
	// TODO: save model
}
