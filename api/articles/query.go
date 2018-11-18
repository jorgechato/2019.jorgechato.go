package articles

import (
	"github.com/graphql-go/graphql"

	. "github.com/jorgechato/api.jorgechato.com/api/types"
)

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

func GetArticleByID() (field *graphql.Field) {
	field = &graphql.Field{
		Type: ArticleType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.ID),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			var article Article
			article.Title = "hello world"

			return article, nil
		},
	}

	return
}
