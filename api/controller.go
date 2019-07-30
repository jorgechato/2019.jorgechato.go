package api

import (
	"net/http"
	// "context"
	"encoding/json"
)

func healthHandler() http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")

		res.WriteHeader(http.StatusOK)

		d := map[string]bool{"alive": true}
		json.NewEncoder(res).Encode(d)
	})
}
