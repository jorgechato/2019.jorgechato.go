package github

import (
	"context"
	"os"

	"github.com/google/go-github/v27/github"
	"golang.org/x/oauth2"

	. "jorgechato.com/utils"
)

var Client *github.Client
var Ctx context.Context

func init() {
	Ctx = context.Background()

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv(GITHUB_TOKEN)},
	)
	httpClient := oauth2.NewClient(Ctx, ts)

	Client = github.NewClient(httpClient)
}
