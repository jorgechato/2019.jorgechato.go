package project

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/go-github/v27/github"

	g "jorgechato.com/github"
	. "jorgechato.com/utils"
)

func Repos(c *gin.Context) {
	var repos Repositories

	res, _, _ := g.Client.Search.Repositories(
		g.Ctx,
		fmt.Sprintf(
			"topic:%v user:%v",
			os.Getenv(GITHUB_TOPIC),
			os.Getenv(GITHUB_USER),
		),
		&github.SearchOptions{
			Sort: "updated",
		},
	)

	r, _ := json.Marshal(res)

	json.Unmarshal(r, &repos)

	c.JSON(
		http.StatusOK,
		repos,
	)
}
