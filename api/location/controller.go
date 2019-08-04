package location

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"

	. "jorgechato.com/utils"
)

func Today(c *gin.Context) {
	var w Where

	if entry, err := Cache.Get(CACHE_KEY_LOCATION); err == nil {
		json.Unmarshal(entry, &w)
	} else {
		res, _ := http.Get(LOCATION_API)

		defer res.Body.Close()

		var polarsteps Polarsteps

		json.NewDecoder(res.Body).Decode(&polarsteps)

		w.build(polarsteps)
		cached, _ := json.Marshal(&w)

		Cache.Set(
			CACHE_KEY_LOCATION,
			cached,
		)
	}

	c.JSON(
		http.StatusOK,
		w,
	)
}
