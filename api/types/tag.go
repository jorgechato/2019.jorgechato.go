package types

import (
	"strings"

	"github.com/jinzhu/gorm"
)

type Tag struct {
	Model
	Name       string       `json:"name"`
	Articles   []*Article   `gorm:"many2many:article_tags" json:"-"`
	Affiliates []*Affiliate `gorm:"many2many:affiliate_tags" json:"-"`
}

func (tag *Tag) BeforeSave(scope *gorm.Scope) error {
	scope.SetColumn(
		"name",
		strings.ToLower(tag.Name),
	)

	return nil
}
