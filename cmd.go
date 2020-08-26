package main

import (
	"encoding/json"
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/robfig/cron/v3"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"storj-prometheus-exporter/common"
	"storj-prometheus-exporter/miners/storj"
)

func main() {
	configFile, err := os.Open("config.json")

	if err != nil {
		log.Fatalf("Error while open config file: %s", err.Error())
		os.Exit(1)
	}

	defer configFile.Close()

	configByteValue, _ := ioutil.ReadAll(configFile)

	var config common.Config

	err = json.Unmarshal(configByteValue, &config)

	if err != nil {
		log.Fatalf("Error while load config file: %s", err.Error())
		os.Exit(2)
	}

	targets := []common.Target{
		new(storj.Target),
	}

	cronScheduler := cron.New()

	for _, monitor := range config.Monitors {
		targetFound := false
		for _, target := range targets {
			if target.GetApiName() == monitor.Api {
				log.Printf("Scheduling update for %s on: %s (Each %ds)", monitor.Api, monitor.Url, monitor.Interval)
				monitorInstance := target.GenMonitor(monitor)
				_, err := cronScheduler.AddFunc(fmt.Sprintf("@every %ds", monitor.Interval), func() {
					monitorInstance.Update()
				})
				if err != nil {
					log.Fatal(err.Error())
				} else {
					targetFound = true
				}
				break
			}
		}
		if !targetFound {
			log.Printf("Not api of type %s found.", monitor.Api)
		}
	}

	cronScheduler.Start()
	defer cronScheduler.Stop()

	http.Handle("/metrics", promhttp.Handler())

	if config.Port < 1 || config.Port > int(math.MaxUint16) {
		config.Port = 2112
	}
	listenUri := fmt.Sprintf(":%d", config.Port)
	log.Print(fmt.Sprintf("Listening on %s", listenUri))

	http.ListenAndServe(listenUri, nil)
}
