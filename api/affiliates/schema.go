package affiliates

import (
	"github.com/graphql-go/graphql"

	"github.com/jorgechato/api.jorgechato.com/api/buckets"
	. "github.com/jorgechato/api.jorgechato.com/api/tags"
	. "github.com/jorgechato/api.jorgechato.com/api/types"
)

var AffiliateType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Affiliate",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.ID),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if obj, ok := p.Source.(Affiliate); ok == true {
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
			Type: graphql.String,
		},
		"score": &graphql.Field{
			Type: graphql.Int,
		},
	},
})

func init() {
	AffiliateType.AddFieldConfig(
		"tags",
		&graphql.Field{
			Type: graphql.NewList(TagType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if obj, ok := p.Source.(Affiliate); ok == true {
					obj.Tags = getTags(obj)

					return obj.Tags, nil
				}
				return nil, nil
			},
		})

	AffiliateType.AddFieldConfig(
		"bucket",
		&graphql.Field{
			Type: buckets.BucketType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if obj, ok := p.Source.(Affiliate); ok == true {
					return buckets.GetBucketByID(obj.Bucket), nil
				}
				return nil, nil
			},
		})
}
