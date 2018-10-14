package schema

import (
	"bytes"

	graphql "github.com/graph-gophers/graphql-go"
)

type query struct{}

func (_ *query) Hello() string { return "Hello, world!" }

// Get the schema of graphql
func Get() (schema *graphql.Schema) {
	s := `
                schema {
                        query: Query
                }
                type Query {
                        hello: String!
                }
        `

	schema = graphql.MustParseSchema(s, &query{})
	return
}

// // If this method complains about not finding functions AssetNames() or MustAsset(),
// // run `go generate` against this package to generate the functions.
// func read() string {
// buf := bytes.Buffer{}
// for _, name := range AssetNames() {
// b := MustAsset(name)
// buf.Write(b)
//
// // Add a newline if the file does not end in a newline.
// if len(b) > 0 && b[len(b)-1] != '\n' {
// buf.WriteByte('\n')
// }
// }
//
// return buf.String()
// }
