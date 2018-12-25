package mutation

import (
	"time"

	"github.com/graphql-go/graphql"

	"github.com/jorgechato/api.jorgechato.com/api/graphql/controller"
	. "github.com/jorgechato/api.jorgechato.com/api/types"
)

var miscArg = graphql.FieldConfigArgument{
	"title": &graphql.ArgumentConfig{
		Type:         graphql.String,
		DefaultValue: "",
	},
	"content": &graphql.ArgumentConfig{
		Type:         graphql.String,
		DefaultValue: "",
	},
	"url": &graphql.ArgumentConfig{
		Type:         graphql.String,
		DefaultValue: "",
	},
	"thumbnail": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"public": &graphql.ArgumentConfig{
		Type:         graphql.Boolean,
		DefaultValue: false,
	},
	"published_at": &graphql.ArgumentConfig{
		Type:         graphql.DateTime,
		DefaultValue: time.Now(),
	},
	"id": {
		Type:         graphql.ID,
		DefaultValue: "",
	},
}

// CreateMisc create a misc
func CreateMisc() (field *graphql.Field) {
	field = &graphql.Field{
		Type: MiscType,
		Args: miscArg,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return controller.CreateMisc(p)
		},
	}

	return
}

// UpdateMisc create a misc
func UpdateMisc() (field *graphql.Field) {
	field = &graphql.Field{
		Type: MiscType,
		Args: miscArg,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return controller.UpdateMisc(p)
		},
	}

	return
}
