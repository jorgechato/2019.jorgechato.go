package types

import (
	"time"

	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
)

type Misc struct {
	gorm.Model
	ID           string    `json:"id,omitempty"`
	Created_at   time.Time `json:"created_at"`
	Title        string    `json:"title"`
	Url          string    `json:"url"`
	Content      string    `json:"content"`
	Published_at time.Time `json:"published_at"`
	Public       bool      `json:"public"`
	Thumbmail    string    `json:"thumbnail"`
}

var MiscType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Misc",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.ID,
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
