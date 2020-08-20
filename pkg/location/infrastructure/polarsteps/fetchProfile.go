package polarsteps

import (
	"net/url"
	"os"

	"jorgechato.com/pkg/shared/infrastructure/api"
)

func FetchProfile() ([]byte, error) {
	u := url.URL{
		Scheme: api.HTTPS,
		Host:   BASE_API,
		Path:   PROFILE + os.Getenv(USER),
	}

	return api.Get(u.String())
}
