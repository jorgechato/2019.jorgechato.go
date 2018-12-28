package api

import (
	"github.com/graphql-go/graphql"

	"github.com/jorgechato/api.jorgechato.com/api/graphql/mutation"
	"github.com/jorgechato/api.jorgechato.com/api/graphql/query"
)

func buildSchema() (schema graphql.Schema) {
	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"articles": query.GetArticles(),
			"article":  query.GetArticleByID(),

			"buckets": query.GetBuckets(),
			"bucket":  query.GetBucketByID(),

			"affiliates": query.GetAffiliates(),
			"affiliate":  query.GetAffiliateByID(),

			"miscs": query.GetMisc(),
			"misc":  query.GetMiscByID(),
		},
	})

	rootMutation := graphql.NewObject(graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			"createArticle": mutation.CreateArticle(),
			"updateArticle": mutation.UpdateArticle(),
			"deleteArticle": mutation.DeleteArticle(),

			"createAffiliate": mutation.CreateAffiliate(),
			"updateAffiliate": mutation.UpdateAffiliate(),
			"deleteAffiliate": mutation.DeleteAffiliate(),

			"deleteBucket": mutation.DeleteBucket(),
			"cleanBucket":  mutation.CleanBucket(),
			"CreateBucket": mutation.CreateBucket(),

			"createMisc": mutation.CreateMisc(),
			"updateMisc": mutation.UpdateMisc(),
		},
	})

	schema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query:    rootQuery,
		Mutation: rootMutation,
	})

	return
}
