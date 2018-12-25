package schema

import (
	"github.com/graphql-go/graphql"

	"github.com/jorgechato/api.jorgechato.com/api/service"
	. "github.com/jorgechato/api.jorgechato.com/api/types"
)

var BucketType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Bucket",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.ID),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if obj, ok := p.Source.(Bucket); ok == true {
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

func init() {
	BucketType.AddFieldConfig(
		"affiliates",
		&graphql.Field{
			Type: graphql.NewList(AffiliateType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if obj, ok := p.Source.(Bucket); ok == true {
					return service.GetAffiliatesByBucket(obj), nil
				}
				return nil, nil
			},
		})
	BucketType.AddFieldConfig(
		"bestSeller",
		&graphql.Field{
			Type: AffiliateType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if obj, ok := p.Source.(Bucket); ok == true {
					return service.GetBestSellerByBucket(obj), nil
				}
				return nil, nil
			},
		})
}
