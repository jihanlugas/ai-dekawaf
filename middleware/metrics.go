package middleware

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "app_http_requests_total",
			Help: "Total HTTP requests processed, partitioned by method, path, status, consumer_type, and consumer_id",
		},
		[]string{"method", "path", "status", "consumer_type", "consumer_id"},
	)

	httpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "app_http_request_duration_seconds",
			Help:    "HTTP request duration in seconds, partitioned by method, path, consumer_type, and consumer_id",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "path", "consumer_type", "consumer_id"},
	)
)

func init() {
	prometheus.MustRegister(httpRequestsTotal)
	prometheus.MustRegister(httpRequestDuration)
}

func PrometheusMetricsHandler() gin.HandlerFunc {
	h := promhttp.Handler()
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func PrometheusHTTPMetrics() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		duration := time.Since(start).Seconds()
		status := strconv.Itoa(c.Writer.Status())

		consumerType := c.Writer.Header().Get("X-Consumer-Type")
		if consumerType == "" {
			consumerType = "unknown"
		}

		consumerID := c.Writer.Header().Get("X-Consumer-ID")
		if consumerID == "" {
			consumerID = "anonymous"
		}

		path := c.FullPath()
		if path == "" {
			path = c.Request.URL.Path
		}

		httpRequestsTotal.WithLabelValues(c.Request.Method, path, status, consumerType, consumerID).Inc()
		httpRequestDuration.WithLabelValues(c.Request.Method, path, consumerType, consumerID).Observe(duration)
	}
}
