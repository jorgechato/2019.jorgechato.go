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
			"articles": articles.GetArticles(),
			"article":  articles.GetArticleByID(),
			"miscs":    misc.GetMisc(),
			"misc":     misc.GetMiscByID(),
		},
	})

	rootMutation := graphql.NewObject(graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			"createArticle": articles.CreateArticle(),
			"updateArticle": articles.UpdateArticle(),
			"deleteArticle": articles.DeleteArticle(),
			"createMisc":    misc.CreateMisc(),
			"updateMisc":    misc.UpdateMisc(),
		},
	})

	schema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query:    rootQuery,
		Mutation: rootMutation,
	})

	return
}
