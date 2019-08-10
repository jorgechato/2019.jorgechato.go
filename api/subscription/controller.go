package subscription

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"

	. "jorgechato.com/utils"
)

func List(c *gin.Context) {
	client := &http.Client{}

	var s Subscritor
	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	s.Status = "pending"

	data := new(bytes.Buffer)
	json.NewEncoder(data).Encode(s)

	query := fmt.Sprintf(
		MAILCHIMP_API,
		os.Getenv(MAILCHIMP_DS),
		os.Getenv(MAILCHIMP_LIST),
		fmt.Sprintf(
			"%x",
			md5.Sum(
				[]byte(
					strings.ToLower(s.EmailAddress),
				),
			),
		),
	)

	req, _ := http.NewRequest(
		http.MethodPut,
		query,
		data,
	)

	req.SetBasicAuth("", os.Getenv(MAILCHIMP_KEY))
	res, _ := client.Do(req)
	defer res.Body.Close()

	c.JSON(
		res.StatusCode,
		res.Body,
	)
}
