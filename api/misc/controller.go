package misc

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	g "jorgechato.com/github"
	. "jorgechato.com/utils"
)

func Lists(c *gin.Context) {
	res, _, _ := g.Client.Gists.Get(
		g.Ctx,
		os.Getenv(GIST_LIST),
	)

	c.JSON(
		http.StatusOK,
		res.Files,
	)
}
