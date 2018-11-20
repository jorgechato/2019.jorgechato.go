package articles

import (
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
				Type:         graphql.Boolean,
				DefaultValue: false,
			},
			"published_at": &graphql.ArgumentConfig{
				Type:         graphql.DateTime,
				DefaultValue: time.Now(),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			var article Article

			article.Title = p.Args["title"].(string)
			article.Content = p.Args["content"].(string)
			article.Public = p.Args["public"].(bool)

			time, _ := time.Parse(time.RFC3339, p.Args["published_at"].(string))
			article.PublishedAt = time

			if obj, ok := p.Args["thumbnail"].(string); ok == true {
				article.Thumbmail = obj
			}

			createArticle(&article)

			return article, nil
		},
	}

	return
}
