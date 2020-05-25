package api

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"jorgechato.com/pkg/location/application"
)

func Router(APIBASE string, router *gin.Engine) {
	var service application.LocationUseCase
	handler := NewHandler(service)

	v := router.Group(fmt.Sprintf("%s/location", APIBASE))
	{
		v.GET("/current", handler.GetCurrent)
		v.GET("/next", handler.GetNext)
	}
}
