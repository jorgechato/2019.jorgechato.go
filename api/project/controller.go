package project

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/go-github/v27/github"
	"golang.org/x/oauth2"

	. "jorgechato.com/utils"
)

var client *github.Client
var ctx context.Context

func init() {
	ctx = context.Background()

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv(GITHUB_TOKEN)},
	)
	httpClient := oauth2.NewClient(ctx, ts)

	client = github.NewClient(httpClient)
}

func Repos(c *gin.Context) {
	var repos Repositories

	res, _, _ := client.Search.Repositories(
		ctx,
		fmt.Sprintf("topic:%v user:%v", GITHUB_TOPIC, GITHUB_USER),
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
