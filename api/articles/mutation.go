package articles

import (
	"time"

	"github.com/graphql-go/graphql"
)

// CreateArticle create an article
func CreateArticle() (field *graphql.Field) {
	field = &graphql.Field{
		Type: ArticleType,
		Args: graphql.FieldConfigArgument{
			"title": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"content": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"thumbnail": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"public": &graphql.ArgumentConfig{
				Type:         graphql.Boolean,
				DefaultValue: false,
			},
			"tags": &graphql.ArgumentConfig{
				Type: graphql.NewList(graphql.String),
			},
			"published_at": &graphql.ArgumentConfig{
				Type:         graphql.DateTime,
				DefaultValue: time.Now(),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return create(p)
		},
	}

	return
}

// UpdateArticle update an article
func UpdateArticle() (field *graphql.Field) {
	field = &graphql.Field{
		Type: ArticleType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.ID),
			},
			"title": &graphql.ArgumentConfig{
				Type:         graphql.String,
				DefaultValue: "",
			},
			"content": &graphql.ArgumentConfig{
				Type:         graphql.String,
				DefaultValue: "",
			},
			"tags": &graphql.ArgumentConfig{
				Type: graphql.NewList(graphql.String),
			},
			"thumbnail": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"public": &graphql.ArgumentConfig{
				Type:         graphql.Boolean,
				DefaultValue: false,
			},
			"published_at": &graphql.ArgumentConfig{
				Type:         graphql.DateTime,
				DefaultValue: time.Now(),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return update(p)
		},
	}

	return
}

// DeleteArticle delete an article
func DeleteArticle() (field *graphql.Field) {
	field = &graphql.Field{
		Type: ArticleType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.ID),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return delete(p)
		},
	}

	return
}
