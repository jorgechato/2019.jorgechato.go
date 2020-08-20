package api

import (
	"github.com/gin-gonic/gin"

	"jorgechato.com/pkg/location/application"
)

func Router(BASE string, router *gin.Engine) {
	var service application.LocationUseCase
	handler := NewHandler(service)

	v := router.Group(BASE + "/location")
	{
		v.GET("/current", handler.GetCurrent)
		v.GET("/next", handler.GetNext)
	}
}
