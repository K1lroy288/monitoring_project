package main

import (
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/shirou/gopsutil/v3/cpu"
)

var httpRequestsTotal = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total number of HTTP requests to monitoring project",
	},
	[]string{"path"},
)

var cpuLoaded = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Name: "CPU_loaded",
	},
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
	prometheus.MustRegister(httpRequestsTotal, cpuLoaded)

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		httpRequestsTotal.WithLabelValues(r.URL.Path).Inc()
		w.Write([]byte("Service is up!"))
	})

	go cpuLoad(cpuLoaded)

	http.Handle("/metrics", promhttp.Handler())

	http.ListenAndServe(":3425", nil)
}
