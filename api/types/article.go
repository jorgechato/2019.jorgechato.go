package types

import (
	"time"

	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
)

type Article struct {
	gorm.Model
	ID           string    `json:"id,omitempty"`
	Created_at   time.Time `json:"created_at"`
	Title        string    `json:"title"`
	Tags         []*Tag    `gorm:"many2many:article_tags json:"tags"`
	Slug         string    `json:"slug"`
	Content      string    `json:"content"`
	Published_at time.Time `json:"published_at"`
	Public       bool      `json:"public"`
	Thumbmail    string    `json:"thumbnail"`
}

type Tag struct {
	gorm.Model
	ID       string     `json:"id,omitempty"`
	Name     string     `json:"tag"`
	Articles []*Article `gorm:"many2many:article_tags json:"articles"`
}

var ArticleType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Article",
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
		"tags": &graphql.Field{
			// TODO: change type or not?
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
