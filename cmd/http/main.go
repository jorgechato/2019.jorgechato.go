package main

import (
	"fmt"
	"io"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	location "jorgechato.com/pkg/location/infrastructure/api"
)

var (
	PORT    string = "5000"
	HOST    string = "0.0.0.0"
	APIBASE string = "/v2"

	LOGPATH string = "out/api.jorgechato.com.log"

	address string = fmt.Sprintf("%v:%v", HOST, PORT)
)

func main() {
	// Logging to a file.
	f, _ := os.Create(LOGPATH)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{
		"https://jorgechato.com",
		"https://whereisjorge.today",
	}
	router.Use(cors.New(config))

	location.Router(APIBASE, router)

	router.Run(address)
}
