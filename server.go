package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/jorgechato/2019.jorgechato.go/api"
	"github.com/jorgechato/2019.jorgechato.go/metrics"
)

var (
	version           = "1.0"
	tag               = ""
	author            = "jorgechato"
	port              = "8080"
	host              = ""
	idleTimeout       = time.Second * 60
	writeTimeout      = time.Second * 10
	readHeaderTimeout = time.Second * 1
	maxHeaderBytes    = http.DefaultMaxHeaderBytes
	address           = fmt.Sprintf("%v:%v", host, port)
	info              = map[string]string{
		"Build User":   "@" + author,
		"Version":      "v" + version,
		"Version desc": "v" + tag,
		"Server":       "http://" + address,
	}
	start = time.Now()
)

func main() {
	b, _ := json.MarshalIndent(info, "", "  ")
	fmt.Println(string(b))

	server := &http.Server{
		Addr:              address,
		Handler:           api.Build(),
		ReadHeaderTimeout: readHeaderTimeout,
		WriteTimeout:      writeTimeout,
		IdleTimeout:       idleTimeout,
		MaxHeaderBytes:    maxHeaderBytes,
	}

	go func() {
		log.Fatal(server.ListenAndServe())
	}()

	metrics.RampUpTime.Set(
		float64(time.Since(start).Nanoseconds()),
	)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), 2)
	defer cancel()
	server.Shutdown(ctx)
	os.Exit(0)
}
