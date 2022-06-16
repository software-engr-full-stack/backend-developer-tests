package metrics

import (
    "net/http"
    "strconv"

    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
    namespace = "stackpath_dev_tests_rest_service"
    subsystem = "development"
)

func Init() error {
    err := prometheus.Register(totalRequests)
    if err != nil {
        return err
    }
    err = prometheus.Register(responseStatus)
    if err != nil {
        return err
    }
    err = prometheus.Register(httpDuration)
    if err != nil {
        return err
    }

    return nil
}

func Middleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        path := r.URL.Path

        timer := prometheus.NewTimer(httpDuration.WithLabelValues(path))
        rw := newResponseWriter(w)
        next.ServeHTTP(rw, r)

        statusCode := rw.statusCode

        responseStatus.WithLabelValues(strconv.Itoa(statusCode)).Inc()
        totalRequests.WithLabelValues(path).Inc()

        timer.ObserveDuration()
    })
}

func HandleFunc(w http.ResponseWriter, r *http.Request) {
    promhttp.Handler().ServeHTTP(w, r)
}

type responseWriter struct {
    http.ResponseWriter
    statusCode int
}

func newResponseWriter(w http.ResponseWriter) *responseWriter {
    return &responseWriter{w, http.StatusOK}
}

// Notes: I commented this out by mistake because in the process of editing, my editor
// marked this method as unused. The problem manifested itself in the Grafana
// stackpath_dev_tests_rest_service_development_http_response_status query.
// It would not show requests that resulted in 404 status being returned.
func (rw *responseWriter) WriteHeader(code int) {
    rw.statusCode = code
    rw.ResponseWriter.WriteHeader(code)
}

var totalRequests = prometheus.NewCounterVec(
    prometheus.CounterOpts{
        Namespace: namespace,
        Subsystem: subsystem,
        Name: "http_requests_total",
        Help: "Number of requests",
    },
    []string{"path"},
)

var responseStatus = prometheus.NewCounterVec(
    prometheus.CounterOpts{
        Namespace: namespace,
        Subsystem: subsystem,
        Name: "http_response_status",
        Help: "Status of HTTP response",
    },
    []string{"status"},
)

var httpDuration = prometheus.NewHistogramVec(
    prometheus.HistogramOpts{
        Namespace: namespace,
        Subsystem: subsystem,
        Name: "http_response_time_seconds",
        Help: "Duration of HTTP requests",
    },
    []string{"path"},
)
