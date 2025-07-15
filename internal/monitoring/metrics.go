package monitoring

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	RequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"handler", "method", "code"},
	)

	ProviderRequestFailures = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "provider_request_failures_total",
			Help: "Total number of failed requests to providers.",
		},
		[]string{"provider"},
	)
)

func init() {
	prometheus.MustRegister(RequestDuration, ProviderRequestFailures)
}
