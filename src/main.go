package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var httpRequestsTotal = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total number of HTTP requests to monitoring project",
	},
	[]string{"path"},
)

func main() {
	prometheus.MustRegister(httpRequestsTotal)

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		httpRequestsTotal.WithLabelValues(r.URL.Path).Inc()
		w.Write([]byte("Service is up!"))
	})

	http.Handle("/metrics", promhttp.Handler())

	http.ListenAndServe(":3425", nil)
}
