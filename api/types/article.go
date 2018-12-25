package types

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Article struct {
	Model
	Title       string    `json:"title"`
	Tags        []*Tag    `gorm:"many2many:article_tags" json:"tags"`
	Slug        string    `json:"slug"`
	Content     string    `json:"content"`
	PublishedAt time.Time `json:"published_at"`
	Public      bool      `json:"public"`
	Thumbmail   string    `json:"thumbnail"`
}

func (article *Article) BeforeSave(scope *gorm.Scope) error {
	scope.SetColumn(
		"slug",
		slugify(article.Title),
	)

	return nil
}
