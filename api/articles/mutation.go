package articles

import (
	"github.com/graphql-go/graphql"

	. "github.com/jorgechato/api.jorgechato.com/api/types"
)

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
				Type: graphql.Boolean,
			},
			"slug": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"created_at": &graphql.ArgumentConfig{
				Type: graphql.DateTime,
			},
			"published_at": &graphql.ArgumentConfig{
				Type: graphql.DateTime,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			var article Article
			article.Title = p.Args["title"].(string)
			article.Content = p.Args["content"].(string)

			a := createArticle(article)

			return a, nil
		},
	}

	return
}
