package metrics

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	_ "github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	HttpRequestCountTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "http_requests_count_total",
			Help: "Number of HTTP requests in total.",
		})
	HttpRequestCountSuccessful = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "http_requests_count_successful",
			Help: "Number of HTTP successful requests.",
		})
	HttpRequestCountError = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "http_requests_count_error",
			Help: "Number of HTTP error requests.",
		})
	HttpRequestCountNotFound = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "http_requests_count_not_found",
			Help: "Number of HTTP NOT FOUND requests.",
		})
	HttpRequestDuration = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Name: "http_request_duration",
			Help: "Tracks a time server took to handle a request.",
		})
)

func init() {
	prometheus.MustRegister(HttpRequestCountTotal)
	prometheus.MustRegister(HttpRequestCountSuccessful)
	prometheus.MustRegister(HttpRequestCountError)
	prometheus.MustRegister(HttpRequestCountNotFound)
	prometheus.MustRegister(HttpRequestDuration)
}

func PrometheusMiddleware(ctx *gin.Context) {
	timer := prometheus.NewTimer(HttpRequestDuration)
	ctx.Next()
	timer.ObserveDuration()
}
