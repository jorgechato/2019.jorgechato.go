package affiliates

import (
	"github.com/graphql-go/graphql"
)

var arguments = graphql.FieldConfigArgument{
	"name": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"url": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"thumbnail": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"bucket": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"preview": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"public": &graphql.ArgumentConfig{
		Type:         graphql.Boolean,
		DefaultValue: false,
	},
	"tags": &graphql.ArgumentConfig{
		Type: graphql.NewList(graphql.String),
	},
	"score": {
		Type:         graphql.Int,
		DefaultValue: 0,
	},
	"id": {
		Type:         graphql.ID,
		DefaultValue: "",
	},
}

// CreateAffiliate create an Affiliate
func CreateAffiliate() (field *graphql.Field) {
	field = &graphql.Field{
		Type: AffiliateType,
		Args: arguments,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return create(p)
		},
	}

	return
}

// UpdateAffiliate update an Affiliate
func UpdateAffiliate() (field *graphql.Field) {
	field = &graphql.Field{
		Type: AffiliateType,
		Args: arguments,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return update(p)
		},
	}

	return
}

// DeleteAffiliate delete an Affiliate
func DeleteAffiliate() (field *graphql.Field) {
	field = &graphql.Field{
		Type: AffiliateType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.ID),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return delete(p)
		},
	}

	return
}
