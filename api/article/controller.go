package article

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		"",
	)
}

func ByID(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		c.Param("id"),
	)
}
