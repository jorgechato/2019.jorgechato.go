package query

import (
	"github.com/graphql-go/graphql"

	"github.com/jorgechato/api.jorgechato.com/api/graphql/controller"
	. "github.com/jorgechato/api.jorgechato.com/api/graphql/schema"
)

// GetBuckets get all Buckets
func GetBuckets() (field *graphql.Field) {
	field = &graphql.Field{
		Type: graphql.NewList(BucketType),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return controller.GetBuckets(p)
		},
	}

	return
}

// GetBucketsyID get Bucketsby ID
func GetBucketByID() (field *graphql.Field) {
	field = &graphql.Field{
		Type: BucketType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.ID),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return controller.GetBucket(p)
		},
	}

	return
}
