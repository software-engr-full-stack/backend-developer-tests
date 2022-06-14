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

	mux.HandleFunc("/people", handlers.People)
	mux.HandleFunc("/people/", handlers.People)

	port := "8000"

    srv := &http.Server{
    	Handler:      mux,
        Addr:         "0.0.0.0:" + port,
        WriteTimeout: 15 * time.Second,
        ReadTimeout:  15 * time.Second,
    }

    fmt.Println("Server started on PORT " + port)
    log.Fatal(srv.ListenAndServe())
}
