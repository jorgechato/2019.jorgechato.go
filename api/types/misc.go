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
			Type: graphql.NewNonNull(graphql.ID),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if obj, ok := p.Source.(Misc); ok == true {
					return obj.ID, nil
				}
				return nil, nil
			},
		},
		"created_at": &graphql.Field{
			Type: graphql.DateTime,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if obj, ok := p.Source.(Misc); ok == true {
					return obj.CreatedAt, nil
				}
				return nil, nil
			},
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"url": &graphql.Field{
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
