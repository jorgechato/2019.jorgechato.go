package articles

import (
	"github.com/graphql-go/graphql"

	. "github.com/jorgechato/api.jorgechato.com/api/types"
)

// GetArticles get all articles
func GetArticles() (field *graphql.Field) {
	field = &graphql.Field{
		Type: graphql.NewList(ArticleType),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			articles := getArticles()

			return articles, nil
		},
	}

	return
}

// GetArticleByID get article by ID
func GetArticleByID() (field *graphql.Field) {
	field = &graphql.Field{
		Type: ArticleType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			article := getArticleByID(p.Args["id"].(string))

			return article, nil
		},
	}

	return
}
