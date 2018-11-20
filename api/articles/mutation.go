package articles

import (
	"log"
	"time"

	"github.com/graphql-go/graphql"

	. "github.com/jorgechato/api.jorgechato.com/api/types"
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
				Type: graphql.Boolean,
			},
			"slug": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"published_at": &graphql.ArgumentConfig{
				Type: graphql.DateTime,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			var article Article

			article.Title = p.Args["title"].(string)
			article.Content = p.Args["content"].(string)

			if obj, ok := p.Args["public"].(bool); ok == true {
				article.Public = obj
			}
			if obj, ok := p.Args["thumbnail"].(string); ok == true {
				article.Thumbmail = obj
			}
			if obj, ok := p.Args["published_at"].(time.Time); ok == true {
				article.PublishedAt = obj
			}
			_, ok := p.Args["published_at"].(time.Time)
			log.Println(ok)

			createArticle(&article)

			return article, nil
		},
	}

	return
}
