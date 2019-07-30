package schema

import (
	"github.com/graphql-go/graphql"

	// "github.com/jorgechato/api.jorgechato.com/api/graphql/controller"
	"github.com/jorgechato/api.jorgechato.com/api/service"
	. "github.com/jorgechato/api.jorgechato.com/api/types"
)

var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
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
		"thumbnail": &graphql.Field{
			Type: graphql.String,
		},
	},
})

func init() {
	// UserType.AddFieldConfig(
	// "location",
	// &graphql.Field{
	// Type: LocationType,
	// Resolve: func(p graphql.ResolveParams) (interface{}, error) {
	// return controller.GetLocation(p)
	// },
	// })

	// UserType.AddFieldConfig(
	// "todos",
	// &graphql.Field{
	// Type: graphql.NewList(TodoType),
	// Args: graphql.FieldConfigArgument{
	// "first": &graphql.ArgumentConfig{
	// Type:         graphql.Int,
	// DefaultValue: -1,
	// },
	// "offset": &graphql.ArgumentConfig{
	// Type:         graphql.Int,
	// DefaultValue: -1,
	// },
	// },
	// Resolve: func(p graphql.ResolveParams) (interface{}, error) {
	// return controller.GetTodos(), nil
	// },
	// })

	UserType.AddFieldConfig(
		"video",
		&graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return service.GetLastVideo(), nil
			},
		})
}
