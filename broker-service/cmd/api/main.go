package main

import (
	"fmt"
	"log"
	"net/http"
)

const webPort = "8000"

type Config struct {
}

func main() {
	app := Config{}

	log.Printf("Starting the broker service on %s\n", webPort)

	// define web server
	s := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	// start the server
	err := s.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
