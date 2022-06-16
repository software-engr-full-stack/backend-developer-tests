package main

import (
	"net/http"
	"time"
	"log"

	"fmt"

	// "strconv"
	// "github.com/prometheus/client_golang/prometheus"
	// "github.com/prometheus/client_golang/prometheus/promhttp"
	// // "github.com/prometheus/client_golang/prometheus/promauto"

	"github.com/software-engr-full-stack/backend-developer-tests/rest-service/pkg/handlers"
	"github.com/software-engr-full-stack/backend-developer-tests/rest-service/pkg/metrics"
)

func main() {
	fmt.Println("SP// Backend Developer Test - RESTful Service")
	fmt.Println()

	var err = metrics.Init()
	// var err = metricsInit()
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()

	// Notes: I wanted to go for the minimal approach hence this ugly hack.
	// Using tools like gorilla mux will make these 2 lines look a lot better.
	mux.HandleFunc("/people", handlers.People)
	mux.HandleFunc("/people/", handlers.People)

	mux.HandleFunc("/metrics", metrics.HandleFunc)
	// mux.HandleFunc("/metrics", metricsHandleFunc)

	port := "8000"

	// Notes: written this way so I can easily add middleware in the future.
	// For example: "Handler: cors.Handler(middleware.WithLogging(mux))"
    srv := &http.Server{
    	Handler:      metrics.Middleware(mux),
    	// Handler:      metricsMiddleware(mux),
        Addr:         "0.0.0.0:" + port,
        WriteTimeout: 15 * time.Second,
        ReadTimeout:  15 * time.Second,
    }

    fmt.Println("Server started on PORT " + port)
    log.Fatal(srv.ListenAndServe())
}

// func metricsHandleFunc(w http.ResponseWriter, r *http.Request) {
//     promhttp.Handler().ServeHTTP(w, r)
// }

// type responseWriter struct {
// 	http.ResponseWriter
// 	statusCode int
// }

// func NewResponseWriter(w http.ResponseWriter) *responseWriter {
// 	return &responseWriter{w, http.StatusOK}
// }

// func (rw *responseWriter) WriteHeader(code int) {
// 	rw.statusCode = code
// 	rw.ResponseWriter.WriteHeader(code)
// }

// var totalRequests = prometheus.NewCounterVec(
// 	prometheus.CounterOpts{
// 		Name: "http_requests_total",
// 		Help: "Number of get requests.",
// 	},
// 	[]string{"path"},
// )

// var responseStatus = prometheus.NewCounterVec(
// 	prometheus.CounterOpts{
// 		Name: "response_status",
// 		Help: "Status of HTTP response",
// 	},
// 	[]string{"status"},
// )

// var httpDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
// 	Name: "http_response_time_seconds",
// 	Help: "Duration of HTTP requests.",
// }, []string{"path"})
// // var httpDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
// // 	Name: "http_response_time_seconds",
// // 	Help: "Duration of HTTP requests.",
// // }, []string{"path"})

// func metricsMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		path := r.URL.Path

// 		timer := prometheus.NewTimer(httpDuration.WithLabelValues(path))
// 		rw := NewResponseWriter(w)
// 		next.ServeHTTP(rw, r)

// 		statusCode := rw.statusCode

// 		responseStatus.WithLabelValues(strconv.Itoa(statusCode)).Inc()
// 		totalRequests.WithLabelValues(path).Inc()

// 		timer.ObserveDuration()
// 	})
// }

// func metricsInit() error {
//     err := prometheus.Register(totalRequests)
//     if err != nil {
//         return err
//     }
//     err = prometheus.Register(responseStatus)
//     if err != nil {
//         return err
//     }
//     err = prometheus.Register(httpDuration)
//     if err != nil {
//         return err
//     }

//     return nil
// }
