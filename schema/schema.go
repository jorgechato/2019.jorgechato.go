package schema

import graphql "github.com/graph-gophers/graphql-go"

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
