package schema

import (
	"github.com/graphql-go/graphql"

	"github.com/jorgechato/api.jorgechato.com/api/service"
	. "github.com/jorgechato/api.jorgechato.com/api/types"
)

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

func init() {
	ArticleType.AddFieldConfig(
		"tags",
		&graphql.Field{
			Type: graphql.NewList(TagType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if obj, ok := p.Source.(Article); ok == true {
					obj.Tags = service.GetTagsByArticle(obj)

					return obj.Tags, nil
				}
				return nil, nil
			},
		})
}
