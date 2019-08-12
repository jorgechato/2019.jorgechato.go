package profile

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	g "jorgechato.com/github"
	. "jorgechato.com/utils"
)

func Get(c *gin.Context) {
	var profile Profile

	res, _, _ := g.Client.Users.Get(
		g.Ctx,
		os.Getenv(GITHUB_USER),
	)

	r, _ := json.Marshal(res)

	json.Unmarshal(r, &profile)

	c.JSON(
		http.StatusOK,
		profile,
	)
}
