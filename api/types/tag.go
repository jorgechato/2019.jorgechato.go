package types

import (
	"strings"

	"github.com/graphql-go/graphql"
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

var TagType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Tag",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.ID),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if obj, ok := p.Source.(Tag); ok == true {
					return obj.ID, nil
				}
				return nil, nil
			},
		},
		"name": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
})
