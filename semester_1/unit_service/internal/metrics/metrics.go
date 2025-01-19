package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	TotalUnits = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "total_units",
		Help: "total number of units",
	})

	Requests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "requests",
			Help: "incomming requests",
		},
		[]string{"handler"},
	)
)

func InitPrometheusMetrics() {
	prometheus.MustRegister(TotalUnits)
	prometheus.MustRegister(Requests)
}

func ServeMetrics(w http.ResponseWriter, r *http.Request) {
	promhttp.Handler().ServeHTTP(w, r)
}
