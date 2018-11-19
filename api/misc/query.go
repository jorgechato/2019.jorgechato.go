package misc

import (
	"github.com/graphql-go/graphql"

	. "github.com/jorgechato/api.jorgechato.com/api/types"
)

// GetMisc get all misc
func GetMisc() (field *graphql.Field) {
	field = &graphql.Field{
		Type: MiscType,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			var misc Misc
			misc.Title = "hello world"

			return misc, nil
		},
	}

	return
}
