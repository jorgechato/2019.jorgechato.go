package main

import (
	"fmt"
	"io"
	"os"

	"github.com/gin-gonic/gin"

	"jorgechato.com/api"
	. "jorgechato.com/utils"
)

var (
	address = fmt.Sprintf("%v:%v", HOST, PORT)
)

func main() {
	// Logging to a file.
	f, _ := os.Create(LOGPATH)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	api.Build().Run(address)
}
