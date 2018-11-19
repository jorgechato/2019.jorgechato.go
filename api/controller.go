package api

import (
	"net/http"
	// "context"
	"encoding/json"

	"github.com/graphql-go/graphql"
)

func graphqlHandler() http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")

		var query = struct {
			Query string `json:"query,omitempty"`
		}{""}

		decoder := json.NewDecoder(req.Body)
		decoder.Decode(&query)

		result := graphql.Do(graphql.Params{
			Schema:        buildSchema(),
			RequestString: query.Query,
			// Context: context.WithValue(
			// context.Background(),
			// "token",
			// req.URL.Query().Get("token"),
			// ),
		})

		json.NewEncoder(res).Encode(result)
	})
}

func healthHandler() http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")

		res.WriteHeader(http.StatusOK)

		d := map[string]bool{"alive": true}
		json.NewEncoder(res).Encode(d)
	})
}
