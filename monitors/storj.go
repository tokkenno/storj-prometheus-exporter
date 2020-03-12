package monitors

import (
	"encoding/json"
	"fmt"
	"github.com/prometheus/common/log"
	"github.com/tokkenno/storj-prometheus-exporter/models"
	"github.com/tokkenno/storj-prometheus-exporter/sensors"
	"io/ioutil"
	"net/http"
	"time"
)

type StorjMonitor struct {
	url string
	Refresh time.Duration
}

func NewStorjMonitor(url string) *StorjMonitor {
	return &StorjMonitor{
		url: url,
		Refresh: 5 * time.Second,
	}
}


func (mon *StorjMonitor) Start() {
	go func() {
		for {
			mon.update()
			time.Sleep(mon.Refresh)
		}
	}()
}

func (mon *StorjMonitor) update() {
	err := mon.UpdateDashboard(mon.url)
	if err != nil {
		log.Warn(err)
	}
}

func (mon *StorjMonitor) UpdateDashboard(url string) error {
	start := time.Now()
	resp, err := http.Get(fmt.Sprintf("http://%s/api/dashboard", url))
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	requestInfo := new(models.RequestInfo)
	requestInfo.Duration = time.Since(start)

	var dashboard models.DashboardRequest
	err = json.Unmarshal(body, &dashboard)
	if err != nil {
		return err
	}

	sensors.SetDashboardInfo(url, &dashboard.Data)
	sensors.SetRequestInfo(url, dashboard.Data.NodeID, requestInfo)

	for _, satellite := range dashboard.Data.Satellites {
		err := mon.UpdateSatellite(url, dashboard.Data.NodeID, satellite.ID)
		if err != nil {
			log.Warnf("The satellite %s can't be updated for node %s", satellite.ID, url)
		}
	}

	return nil
}

func (mon *StorjMonitor) UpdateSatellite(url string, nodeId string, satelliteId string) error {
	start := time.Now()
	resp, err := http.Get(fmt.Sprintf("http://%s/api/satellite/%s", url, satelliteId))
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	requestInfo := new(models.RequestInfo)
	requestInfo.Duration = time.Since(start)

	satellite := new(models.SatelliteRequest)
	err = json.Unmarshal(body, satellite)
	if err != nil {
		return err
	}

	sensors.SetSatelliteInfo(url, nodeId, satelliteId, satellite.Data)

	return nil
}
