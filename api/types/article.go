package types

import (
	"time"

	"github.com/graphql-go/graphql"
)

type Article struct {
	Model
	Title       string    `json:"title"`
	Tags        []*Tag    `gorm:"many2many:article_tags" json:"tags"`
	Slug        string    `json:"slug"`
	Content     string    `json:"content"`
	PublishedAt time.Time `json:"published_at"`
	Public      bool      `json:"public"`
	Thumbmail   string    `json:"thumbnail"`
}

var ArticleType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Article",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.ID),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if obj, ok := p.Source.(Article); ok == true {
					return obj.ID, nil
				}
				return nil, nil
			},
		},
		"title": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
		"tags": &graphql.Field{
			// TODO: change type or not?
			Type: graphql.NewList(graphql.String),
		},
		"slug": &graphql.Field{
			Type: graphql.String,
		},
		"content": &graphql.Field{
			Type: graphql.String,
		},
		"created_at": &graphql.Field{
			Type: graphql.DateTime,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if obj, ok := p.Source.(Article); ok == true {
					return obj.CreatedAt, nil
				}
				return nil, nil
			},
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
