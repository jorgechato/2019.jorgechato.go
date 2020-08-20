package goodreads

import (
	"net/url"
	"os"

	"jorgechato.com/pkg/shared/infrastructure/api"
)

func FetchBooks() ([]byte, error) {
	q := url.Values{}
	q.Add("key", os.Getenv(API_KEY))
	q.Add("v", API_V)
	q.Add("sort", SORT)

	u := url.URL{
		Scheme:   api.HTTPS,
		Host:     BASE_API,
		Path:     REVIEW_LIST + os.Getenv(USER),
		RawQuery: q.Encode(),
	}

	return api.Get(u.String())
}
