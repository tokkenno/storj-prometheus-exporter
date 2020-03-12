package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/tokkenno/storj-prometheus-exporter/monitors"
	"net/http"
	"os"
	"strings"
)

const (
	port = 2112
)

func main() {
	var nodes []string

	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		if pair[0] == "NODES" {
			nodes = strings.Split(pair[1], ",")
		}
	}

	if nodes != nil && len(nodes) > 0 {
		for _, nodeUrl := range nodes {
			fmt.Println(fmt.Sprintf("Starting to monitor node in %s", nodeUrl))
			mon := monitors.NewStorjMonitor(nodeUrl)
			mon.Start()
		}

		http.Handle("/metrics", promhttp.Handler())

		listenUri := fmt.Sprintf(":%d", port)
		fmt.Println(fmt.Sprintf("Listening on %s", listenUri))
		http.ListenAndServe(listenUri, nil)
	} else {
		fmt.Println("NODES env variable not set or malformed. Exiting...")
	}
}
