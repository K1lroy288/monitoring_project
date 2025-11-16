package main

import (
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/shirou/gopsutil/v3/cpu"
)

var (
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests to monitoring project",
		},
		[]string{"method", "path", "status"},
	)

	cpuLoaded = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "CPU_loaded",
		},
	)

	httpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "http_request_duration_seconds",
			Help: "Request duration in seconds",
		},
		[]string{"path"},
	)
)

func cpuLoad(cpuL prometheus.Gauge) {
	for {
		percentages, err := cpu.Percent(200*time.Millisecond, false)
		if err != nil {
			log.Fatalf("Не удалосось получить загруженность CPU: %v", err)
		} else {
			cpuL.Set(percentages[0])
		}
		time.Sleep(10 * time.Second)
	}
}

func main() {
	prometheus.MustRegister(httpRequestsTotal, cpuLoaded, httpRequestDuration)

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		w.Write([]byte("Service is up!"))

		duration := time.Since(start).Seconds()
		httpRequestsTotal.WithLabelValues(r.Method, r.URL.Path, "200").Inc()
		httpRequestDuration.WithLabelValues(r.URL.Path).Observe(duration)
	})

	http.HandleFunc("/custom_duration", func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		w.Write([]byte("Custom duration endpoint!"))

		time.Sleep(500 * time.Millisecond)

		duration := time.Since(start).Seconds()
		httpRequestsTotal.WithLabelValues(r.Method, r.URL.Path, "200").Inc()
		httpRequestDuration.WithLabelValues(r.URL.Path).Observe(duration)
	})

	go cpuLoad(cpuLoaded)

	http.Handle("/metrics", promhttp.Handler())

	http.ListenAndServe(":3425", nil)
}
