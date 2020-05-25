package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"jorgechato.com/pkg/location/application/json"
	"jorgechato.com/pkg/location/infrastructure/polarsteps"
)

func (h handler) GetNext(ctx *gin.Context) {
	raw, err := polarsteps.Fetch()

	if err != nil {
		ctx.AbortWithError(
			http.StatusInternalServerError,
			err,
		)
	}

	jsonLocation := &json.Location{}
	location, _ := jsonLocation.Decode(raw)

	ctx.JSON(
		http.StatusOK,
		h.locationUseCase.FetchNext(location),
	)
}
