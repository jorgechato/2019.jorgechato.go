package articles

import (
	. "github.com/jorgechato/api.jorgechato.com/api/types"
)

func deleteArticle(article *Article) {
	Session.
		Delete(article)
}

func getArticles(first, offset int) []Article {
	var articles []Article

	Session.
		Order("published_at desc").
		Offset(offset).
		Limit(first).
		Find(&articles)

	return articles
}

func getArticleByID(id string) Article {
	var article Article

	Session.First(&article, "id = ?", id)

	return article
}

func createArticle(article *Article) {
	Session.Create(article)
}

func updateArticle(article *Article) {
	Session.
		Model(article).
		Omit("created_at").
		Omit("id").
		Update(*article)
}

func getTags(article Article) []*Tag {
	var tags []*Tag

	Session.
		Model(&article).
		Related(&tags, "Tags")

	return tags
}
