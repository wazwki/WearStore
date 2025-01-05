package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	ControllersDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"handler"},
	)

	ServiceDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "business_logic_duration_seconds",
			Help:    "Duration of business logic execution.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"operation"},
	)
)

func init() {
	prometheus.MustRegister(ControllersDuration)
	prometheus.MustRegister(ServiceDuration)
}
