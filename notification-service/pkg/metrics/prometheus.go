package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
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
	prometheus.MustRegister(ServiceDuration)
}
