package service

import (
	. "github.com/jorgechato/api.jorgechato.com/api/types"
)

func DeleteArticle(article *Article) {
	Session.
		Delete(article)
}

func GetArticles(first, offset int) []Article {
	var articles []Article

	Session.
		Order("published_at desc").
		Offset(offset).
		Limit(first).
		Find(&articles)

	return articles
}

func GetArticleByID(id string) Article {
	var article Article

	Session.First(&article, "id = ?", id)

	return article
}

func CreateArticle(article *Article) {
	Session.Create(article)
}

func UpdateArticle(article *Article) {
	Session.
		Model(article).
		Omit("created_at").
		Omit("id").
		Update(*article)
}
