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
    // https://pkg.go.dev/github.com/prometheus/client_golang/prometheus#Registerer
    // Register registers a new Collector to be included in metrics
    // collection. It returns an error if the descriptors provided by the
    // Collector are invalid or if they — in combination with descriptors of
    // already registered Collectors — do not fulfill the consistency and
    // uniqueness criteria described in the documentation of metric.Desc.
    //
    // If the provided Collector is equal to a Collector already registered
    // (which includes the case of re-registering the same Collector), the
    // returned error is an instance of AlreadyRegisteredError, which
    // contains the previously registered Collector.
    //
    // A Collector whose Describe method does not yield any Desc is treated
    // as unchecked. Registration will always succeed. No check for
    // re-registering (see previous paragraph) is performed. Thus, the
    // caller is responsible for not double-registering the same unchecked
    // Collector, and for providing a Collector that will not cause
    // inconsistent metrics on collection. (This would lead to scrape
    // errors.)
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
