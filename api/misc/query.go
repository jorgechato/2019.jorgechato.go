package misc

import (
	"github.com/graphql-go/graphql"

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

			misc := getMiscList(
				p.Args["first"].(int),
				p.Args["offset"].(int),
			)

			return misc, nil
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

			misc := getMiscByID(p.Args["id"].(string))

			return misc, nil
		},
	}

	return
}
