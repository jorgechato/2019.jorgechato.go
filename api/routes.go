package api

import (
	"github.com/gin-gonic/gin"

	"jorgechato.com/api/article"
	"jorgechato.com/api/location"
	"jorgechato.com/api/misc"
	"jorgechato.com/api/profile"
	"jorgechato.com/api/project"
	"jorgechato.com/api/subscription"
	. "jorgechato.com/utils"
)

func Build() (router *gin.Engine) {
	router = gin.Default()

	v := router.Group(APIBASE)
	{
		v.GET("/location", location.Today)
		v.GET("/projects", project.Repos)
		v.GET("/lists", misc.Lists)
		v.GET("/articles", article.List)
		v.GET("/article/:id", article.ByID)
		v.GET("/profile", profile.Get)
		v.POST("/subscribe", subscription.List)
	}

	return
}
