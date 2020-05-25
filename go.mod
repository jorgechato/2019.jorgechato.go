module github.com/jorgechato/api.jorgechato.com

go 1.14

replace jorgechato.com => ./

require (
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.6.3
	github.com/json-iterator/go v1.1.9
	github.com/pkg/errors v0.9.1
	jorgechato.com v0.0.0-00010101000000-000000000000
)
