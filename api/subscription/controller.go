package subscription

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	. "jorgechato.com/utils"
)

func List(c *gin.Context) {
	fmt.Sprintf(
		MAILCHIMP_API,
		os.Getenv(MAILCHIMP_DS),
		os.Getenv(MAILCHIMP_LIST),
	)

	c.JSON(
		http.StatusOK,
		"",
	)
}
