package articles

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	. "github.com/jorgechato/api.jorgechato.com/api/types"
	. "github.com/jorgechato/api.jorgechato.com/utils"
)

func getArticles() []Article {
	db, _ := gorm.Open("postgres", DB)
	defer db.Close()

	var articles []Article
	db.Find(&articles)

	return articles
}

func getArticleByID(id string) Article {
	db, _ := gorm.Open("postgres", DB)
	defer db.Close()

	var article Article
	db.First(&article, "id = ?", id)

	return article
}

func createArticle(article *Article) {
	db, _ := gorm.Open("postgres", DB)
	defer db.Close()

	db.Create(article)
}
