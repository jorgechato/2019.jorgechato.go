package articles

import (
	"time"

	"github.com/graphql-go/graphql"

	"github.com/jorgechato/api.jorgechato.com/api/tags"
	. "github.com/jorgechato/api.jorgechato.com/api/types"
)

func delete(p graphql.ResolveParams) (interface{}, error) {
	article := getArticleByID(p.Args["id"].(string))

	if article.ID != "" {
		deleteArticle(&article)

		return article, nil
	}

	return nil, nil
}

func get(p graphql.ResolveParams) (interface{}, error) {
	article := getArticleByID(p.Args["id"].(string))

	return article, nil
}

func getList(p graphql.ResolveParams) (interface{}, error) {
	articles := getArticles(
		p.Args["first"].(int),
		p.Args["offset"].(int),
	)

	return articles, nil
}

func create(p graphql.ResolveParams) (interface{}, error) {
	var article Article

	article.Title = p.Args["title"].(string)
	article.Content = p.Args["content"].(string)
	article.Public = p.Args["public"].(bool)

	switch arg := p.Args["published_at"].(type) {
	case string:
		time, _ := time.Parse(time.RFC3339, arg)
		article.PublishedAt = time
	case time.Time:
		article.PublishedAt = arg
	}

	if obj, ok := p.Args["thumbnail"].(string); ok == true {
		article.Thumbmail = obj
	}

	if objs, ok := p.Args["tags"].([]interface{}); ok == true {
		for _, obj := range objs {
			tag := tags.GetTagByName(obj.(string))
			article.Tags = append(article.Tags, &tag)
		}
	}

	createArticle(&article)

	return article, nil
}

func update(p graphql.ResolveParams) (interface{}, error) {
	var article Article

	article.ID = p.Args["id"].(string)
	article.Title = p.Args["title"].(string)
	article.Content = p.Args["content"].(string)
	article.Public = p.Args["public"].(bool)

	switch arg := p.Args["published_at"].(type) {
	case string:
		time, _ := time.Parse(time.RFC3339, arg)
		article.PublishedAt = time
	case time.Time:
		article.PublishedAt = arg
	}

	if obj, ok := p.Args["thumbnail"].(string); ok == true {
		article.Thumbmail = obj
	}

	if objs, ok := p.Args["tags"].([]interface{}); ok == true {
		for _, obj := range objs {
			tag := tags.GetTagByName(obj.(string))
			article.Tags = append(article.Tags, &tag)
		}
	}

	updateArticle(&article)

	return article, nil
}
