package mutation

import (
	"github.com/graphql-go/graphql"

	"github.com/jorgechato/api.jorgechato.com/api/graphql/controller"
	. "github.com/jorgechato/api.jorgechato.com/api/graphql/schema"
)

// DeleteBucket delete an Bucket
func DeleteBucket() (field *graphql.Field) {
	field = &graphql.Field{
		Type: BucketType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.ID),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return controller.DeleteBucket(p)
		},
	}

	return
}

// CleanBucket unattach all affiliates from a bucket
func CleanBucket() (field *graphql.Field) {
	field = &graphql.Field{
		Type: BucketType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.ID),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return controller.CleanBucket(p)
		},
	}

	return
}

// CreateBucket unattach all affiliates from a bucket
func CreateBucket() (field *graphql.Field) {
	// TODO
	field = &graphql.Field{
		Type: BucketType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.ID),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return controller.CreateBucket(p)
		},
	}

	return
}
