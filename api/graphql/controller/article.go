package controller

import (
	"strings"
	"time"

	"github.com/graphql-go/graphql"

	"github.com/jorgechato/api.jorgechato.com/api/service"
	. "github.com/jorgechato/api.jorgechato.com/api/types"
)

func DeleteArticle(p graphql.ResolveParams) (interface{}, error) {
	article := service.GetArticleByID(p.Args["id"].(string))

	if article.ID != "" {
		service.DeleteArticle(&article)

		return article, nil
	}

	return nil, nil
}

func GetArticle(p graphql.ResolveParams) (interface{}, error) {
	article := service.GetArticleByID(p.Args["id"].(string))

	return article, nil
}

func GetArticles(p graphql.ResolveParams) (interface{}, error) {
	articles := service.GetArticles(
		p.Args["first"].(int),
		p.Args["offset"].(int),
	)

	return articles, nil
}

func CreateArticle(p graphql.ResolveParams) (interface{}, error) {
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
			tag := service.GetTagByName(strings.ToLower(obj.(string)))
			article.Tags = append(article.Tags, &tag)
		}
	}

	service.CreateArticle(&article)

	return article, nil
}

func UpdateArticle(p graphql.ResolveParams) (interface{}, error) {
	var article Article

	article.ID = p.Args["id"].(string)
	if article.ID == "" {
		return nil, nil
	}

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
			tag := service.GetTagByName(strings.ToLower(obj.(string)))
			article.Tags = append(article.Tags, &tag)
		}
	}

	service.UpdateArticle(&article)

	return article, nil
}
