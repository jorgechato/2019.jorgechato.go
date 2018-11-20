package types

import (
	"time"

	"github.com/graphql-go/graphql"
)

type Misc struct {
	Model
	Title       string    `json:"title"`
	Url         string    `json:"url"`
	Content     string    `json:"content"`
	PublishedAt time.Time `json:"published_at"`
	Public      bool      `json:"public"`
	Thumbmail   string    `json:"thumbnail"`
}

var MiscType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Misc",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"created_at": &graphql.Field{
			Type: graphql.DateTime,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"url": &graphql.Field{
			Type: graphql.String,
		},
		"slug": &graphql.Field{
			Type: graphql.String,
		},
		"content": &graphql.Field{
			Type: graphql.String,
		},
		"published_at": &graphql.Field{
			Type: graphql.DateTime,
		},
		"public": &graphql.Field{
			Type: graphql.Boolean,
		},
		"thumbnail": &graphql.Field{
			Type: graphql.String,
		},
	},
})
