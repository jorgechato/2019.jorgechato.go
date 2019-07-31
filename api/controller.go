package api

import (
	"github.com/gin-gonic/gin"
)

func pong(c *gin.Context) {
	c.String(200, "pong")
}
