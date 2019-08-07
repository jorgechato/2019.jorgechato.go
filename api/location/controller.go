package location

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	. "jorgechato.com/utils"
)

func Today(c *gin.Context) {
	var w Where

	if entry, err := Cache.Get(CACHE_KEY_LOCATION); err == nil {
		json.Unmarshal(entry, &w)
	} else {
		res, _ := http.Get(
			fmt.Sprintf(
				LOCATION_API,
				os.Getenv(LOCATION_USER),
			),
		)

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
