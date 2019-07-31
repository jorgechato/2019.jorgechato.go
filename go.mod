module github.com/jorgechato/api.jorgechato.com

go 1.12

require (
	github.com/gin-gonic/gin v1.4.0
	jorgechato.com/api v0.0.0-00010101000000-000000000000
	jorgechato.com/utils v0.0.0-00010101000000-000000000000
)

replace jorgechato.com/utils => ./utils

replace jorgechato.com/api => ./api
