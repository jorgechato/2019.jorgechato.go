package types

import (
	"time"

	"github.com/graphql-go/graphql"
)

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

var EventType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Event",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.ID,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"thumbnail": &graphql.Field{
			Type: graphql.String,
		},
		"url": &graphql.Field{
			Type: graphql.String,
		},
		"location": &graphql.Field{
			Type: graphql.String,
		},
		"start_at": &graphql.Field{
			Type: graphql.DateTime,
		},
		"end_at": &graphql.Field{
			Type: graphql.DateTime,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
	},
})
