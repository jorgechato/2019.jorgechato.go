package misc

import (
	"github.com/graphql-go/graphql"

	"github.com/jorgechato/api.jorgechato.com/api/types"
)

func GetMisc() (field *graphql.Field) {
	field = &graphql.Field{
		Type: types.MiscType,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			var misc types.Misc
			misc.Title = "hello world"

			return misc, nil
		},
	}

	return
}
