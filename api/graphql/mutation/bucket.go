package mutation

import (
	"github.com/graphql-go/graphql"

	"github.com/jorgechato/api.jorgechato.com/api/graphql/controller"
	. "github.com/jorgechato/api.jorgechato.com/api/graphql/schema"
)

var BucketArg = graphql.FieldConfigArgument{
	"name": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"thumbnail": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"description": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"preview": &graphql.ArgumentConfig{
		Type:         graphql.String,
		DefaultValue: "",
	},
	"id": {
		Type:         graphql.ID,
		DefaultValue: "",
	},
}

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
	field = &graphql.Field{
		Type: BucketType,
		Args: BucketArg,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return controller.CreateBucket(p)
		},
	}

	return
}
