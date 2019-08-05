module github.com/jorgechato/api.jorgechato.com

go 1.12

require (
	github.com/allegro/bigcache v1.2.1
	github.com/gin-gonic/gin v1.4.0
	github.com/google/go-github/v27 v27.0.4
	golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45
	jorgechato.com v0.0.0-00010101000000-000000000000
)

replace jorgechato.com => ./
