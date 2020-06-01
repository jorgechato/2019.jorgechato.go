package polarsteps

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetTripByID(id int) ([]byte, error) {
	res, e := http.Get(
		fmt.Sprintf(
			TRIPS_API,
			id,
		),
	)

	defer res.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(res.Body)

	return bodyBytes, e
}
