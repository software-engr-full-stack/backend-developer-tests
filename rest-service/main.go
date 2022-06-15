package main

import (
	"net/http"
	"time"
	"log"

	"fmt"

	"github.com/software-engr-full-stack/backend-developer-tests/rest-service/pkg/handlers"
)

func main() {
	fmt.Println("SP// Backend Developer Test - RESTful Service")
	fmt.Println()

	mux := http.NewServeMux()

	// Notes: I wanted to go for the minimal approach hence this ugly hack.
	// Using tools like gorilla mux will make these 2 lines look a lot better.
	mux.HandleFunc("/people", handlers.People)
	mux.HandleFunc("/people/", handlers.People)

	port := "8000"

	// Notes: written this way so I can easily add middleware in the future.
	// For example: "Handler: cors.Handler(middleware.WithLogging(mux))"
    srv := &http.Server{
    	Handler:      mux,
        Addr:         "0.0.0.0:" + port,
        WriteTimeout: 15 * time.Second,
        ReadTimeout:  15 * time.Second,
    }

    fmt.Println("Server started on PORT " + port)
    log.Fatal(srv.ListenAndServe())
}
