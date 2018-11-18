package api

import (
	"github.com/graphql-go/graphql"

	"github.com/jorgechato/api.jorgechato.com/api/articles"
	"github.com/jorgechato/api.jorgechato.com/api/misc"
)

func buildSchema() (schema graphql.Schema) {
	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"misc":     misc.GetMisc(),
			"articles": articles.GetArticles(),
			"article":  articles.GetArticleByID(),
		},
	})

	rootMutation := graphql.NewObject(graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			"createArticle": articles.CreateArticle(),
		},
	})

	schema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query:    rootQuery,
		Mutation: rootMutation,
	})

	return
}
