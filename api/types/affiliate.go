package types

import (
	"strings"

	"github.com/jinzhu/gorm"
)

type Affiliate struct {
	Model
	Name      string `json:"name"`
	Url       string `json:"url"`
	Public    bool   `json:"public"`
	Thumbmail string `json:"thumbnail"`
	Preview   string `json:"preview"`
	Bucket    string
	Score     int    `json:"score"`
	Tags      []*Tag `gorm:"many2many:affiliate_tags" json:"tags"`
}

type Bucket struct {
	Model
	Name       string       `gorm:"primary_key" json:"name"`
	Affiliates []*Affiliate `gorm:"foreignkey:Bucket" json:"affiliates"`
}

func (bucket *Bucket) BeforeSave(scope *gorm.Scope) error {
	scope.SetColumn(
		"name",
		strings.ToLower(bucket.Name),
	)

	return nil
}
