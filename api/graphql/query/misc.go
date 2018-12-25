package query

import (
	"github.com/graphql-go/graphql"

	"github.com/jorgechato/api.jorgechato.com/api/graphql/controller"
	. "github.com/jorgechato/api.jorgechato.com/api/types"
)

// GetMisc get all misc
func GetMisc() (field *graphql.Field) {
	field = &graphql.Field{
		Type: graphql.NewList(MiscType),
		Args: graphql.FieldConfigArgument{
			"first": &graphql.ArgumentConfig{
				Type:         graphql.Int,
				DefaultValue: -1,
			},
			"offset": &graphql.ArgumentConfig{
				Type:         graphql.Int,
				DefaultValue: -1,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return controller.GetMiscs(p)
		},
	}

	return
}

// GetMiscByID get misc by ID
func GetMiscByID() (field *graphql.Field) {
	field = &graphql.Field{
		Type: MiscType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.ID),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return controller.GetMisc(p)
		},
	}

	return
}
