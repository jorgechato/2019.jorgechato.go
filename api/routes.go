package api

import (
	"github.com/gin-gonic/gin"
)

func Build() (router *gin.Engine) {
	router = gin.Default()

	router.GET("/", pong)

	return
}
