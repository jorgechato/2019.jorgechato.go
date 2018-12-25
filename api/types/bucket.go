package types

import (
	"github.com/graphql-go/graphql"
)

type Bucket struct {
	Name      string `json:"name"`
	Thumbmail string `json:"thumbnail"`
	Preview   string `json:"preview"`
}

var BucketType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Bucket",
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
			Type: graphql.String,
		},
		"thumbnail": &graphql.Field{
			Type: graphql.String,
		},
		"preview": &graphql.Field{
			Type: graphql.String,
		},
	},
})
