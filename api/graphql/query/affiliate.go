package query

import (
	"github.com/graphql-go/graphql"

	"github.com/jorgechato/api.jorgechato.com/api/graphql/controller"
	. "github.com/jorgechato/api.jorgechato.com/api/graphql/schema"
)

// GetAffiliates get all Affiliates
func GetAffiliates() (field *graphql.Field) {
	field = &graphql.Field{
		Type: graphql.NewList(AffiliateType),
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
			return controller.GetAffiliates(p)
		},
	}

	return
}

// GetAffiliateByID get Affiliate by ID
func GetAffiliateByID() (field *graphql.Field) {
	field = &graphql.Field{
		Type: AffiliateType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.ID),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return controller.GetAffiliate(p)
		},
	}

	return
}
