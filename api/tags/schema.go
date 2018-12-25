package tags

import (
	"github.com/graphql-go/graphql"

	. "github.com/jorgechato/api.jorgechato.com/api/types"
)

var TagType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Tag",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.ID),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if obj, ok := p.Source.(Tag); ok == true {
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
