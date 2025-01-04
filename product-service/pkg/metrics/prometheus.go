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
		[]string{"method", "handler"},
	)

	ServiceDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "business_logic_duration_seconds",
			Help:    "Duration of business logic execution.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"operation"},
	)

	RepositoryDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "db_query_duration_seconds",
			Help:    "Duration of database queries.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"query"},
	)
)

func init() {
	prometheus.MustRegister(ControllersDuration)
	prometheus.MustRegister(ServiceDuration)
	prometheus.MustRegister(RepositoryDuration)
}
