package articles

import (
	. "github.com/jorgechato/api.jorgechato.com/api/types"
)

func getArticles(first, offset int) []Article {
	var articles []Article

	Session.Order("published_at desc").
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
