package types

import (
	"strings"

	"github.com/jinzhu/gorm"
)

type Bucket struct {
	Model
	Name        string       `gorm:"primary_key" json:"name"`
	Description string       `json:"description"`
	Affiliates  []*Affiliate `gorm:"foreignkey:Bucket" json:"affiliates"`
}

func (bucket *Bucket) BeforeSave(scope *gorm.Scope) error {
	scope.SetColumn(
		"name",
		strings.ToLower(bucket.Name),
	)

	return nil
}
