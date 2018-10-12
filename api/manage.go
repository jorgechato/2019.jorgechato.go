package api

import (
	"net/http"

	"github.com/graph-gophers/graphql-go/relay"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/jorgechato/2019.jorgechato.go/schema"
)

// Build the needed endpoints
func Build() (mux *http.ServeMux) {
	// Register handlers to routes.
	mux = http.NewServeMux()
	mux.Handle("/", &relay.Handler{Schema: schema.Get()})
	mux.Handle("/metrics", promhttp.Handler())

	return
}
