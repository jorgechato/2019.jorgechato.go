package api

import (
	"github.com/gin-gonic/gin"

	"jorgechato.com/api/location"
	. "jorgechato.com/utils"
)

func Build() (router *gin.Engine) {
	router = gin.Default()

	v := router.Group(APIBASE)
	{
		v.GET("/where", location.Today)
	}

	return
}
