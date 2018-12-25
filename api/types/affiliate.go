package types

import (
	"github.com/graphql-go/graphql"
)

type Affiliate struct {
	Model
	Name      string  `json:"name"`
	Url       string  `json:"url"`
	Public    bool    `json:"public"`
	Thumbmail string  `json:"thumbnail"`
	Preview   string  `json:"preview"`
	Bucket    *Bucket `gorm:"foreignkey:bucket" json:"bucket"`
	Tags      []*Tag  `gorm:"many2many:affiliate_tags" json:"tags"`
}

var AffiliateType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Affiliate",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.ID),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if obj, ok := p.Source.(Article); ok == true {
					return obj.ID, nil
				}
				return nil, nil
			},
		},
		"name": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
		"url": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
		"public": &graphql.Field{
			Type: graphql.Boolean,
		},
		"thumbnail": &graphql.Field{
			Type: graphql.String,
		},
		"preview": &graphql.Field{
			Type: graphql.String,
		},
		"bucket": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
		"tags": &graphql.Field{
			Type: graphql.NewList(TagType),
		},
	},
})
