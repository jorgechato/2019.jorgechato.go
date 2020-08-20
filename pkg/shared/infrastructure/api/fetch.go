package api

import (
	"io/ioutil"
	"net/http"
)

func Get(query string) ([]byte, error) {
	res, e := http.Get(query)

	defer res.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(res.Body)

	return bodyBytes, e
}
