package polarsteps

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func GetProfile() ([]byte, error) {
	res, e := http.Get(
		fmt.Sprintf(
			PROFILE_API,
			os.Getenv(POLARSTEP_USER),
		),
	)

	defer res.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(res.Body)

	return bodyBytes, e
}
