package main

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/tokkenno/storj-prometheus-exporter/monitors"
	"net/http"
)

func main() {
	mon := monitors.NewStorjMonitor("false.com")
	mon.Start()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}
