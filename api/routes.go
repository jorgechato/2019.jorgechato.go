package api

import (
	"net/http"
	// "context"
	"encoding/json"

	"github.com/graphql-go/graphql"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Build the needed endpoints
func Build() (mux *http.ServeMux) {
	// Register handlers to routes.
	mux = http.NewServeMux()

	mux.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		result := graphql.Do(graphql.Params{
			Schema:        buildSchema(),
			RequestString: request.URL.Query().Get("query"),
			// Context: context.WithValue(
			// context.Background(),
			// "token",
			// request.URL.Query().Get("token"),
			// ),
		})

		json.NewEncoder(response).Encode(result)
	})

	mux.Handle("/metrics", promhttp.Handler())

	return
}
