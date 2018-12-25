package mutation

import (
	"time"

	"github.com/graphql-go/graphql"

	"github.com/jorgechato/api.jorgechato.com/api/graphql/controller"
	. "github.com/jorgechato/api.jorgechato.com/api/graphql/schema"
)

var articleArg = graphql.FieldConfigArgument{
	"title": {
		Type: graphql.NewNonNull(graphql.String),
	},
	"content": {
		Type: graphql.NewNonNull(graphql.String),
	},
	"thumbnail": {
		Type: graphql.String,
	},
	"public": {
		Type:         graphql.Boolean,
		DefaultValue: false,
	},
	"tags": {
		Type: graphql.NewList(graphql.String),
	},
	"published_at": {
		Type:         graphql.DateTime,
		DefaultValue: time.Now(),
	},
	"id": {
		Type:         graphql.ID,
		DefaultValue: "",
	},
}

// CreateArticle create an article
func CreateArticle() (field *graphql.Field) {
	field = &graphql.Field{
		Type: ArticleType,
		Args: articleArg,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return controller.CreateArticle(p)
		},
	}

	return
}

// UpdateArticle update an article
func UpdateArticle() (field *graphql.Field) {
	field = &graphql.Field{
		Type: ArticleType,
		Args: articleArg,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return controller.UpdateArticle(p)
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
			return controller.DeleteArticle(p)
		},
	}

	return
}
