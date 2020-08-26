package storj

import (
	"encoding/json"
	"fmt"
	"github.com/prometheus/common/log"
	"storj-prometheus-exporter/miners/storj/models"
	"storj-prometheus-exporter/miners/storj/sensors"
	"io/ioutil"
	"net/http"
	"storj-prometheus-exporter/common"
	"strings"
	"time"
)

type Monitor struct {
	config common.MonitorConfig
}

func (mon Monitor) GetApiName() string {
	return "storj"
}

func (mon Monitor) Update() {
	err := mon.UpdateSNO(mon.config.Url)
	if err != nil {
		log.Warn(fmt.Sprintf("Error while update node <%s>:", mon.config.Url))
		log.Warn(err)
	}
}

func (mon *Monitor) UpdateSNO(url string) error {
	start := time.Now()
	resp, err := http.Get(fmt.Sprintf("http://%s/api/sno/", url))
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if strings.Contains(string(body), "<!DOCTYPE html>") {
		// Fallback to pre-1.0.0 API
		return mon.UpdateDashboard(url)
	}

	requestInfo := new(models.RequestInfo)
	requestInfo.Duration = time.Since(start)

	var dashboard models.DashboardData
	err = json.Unmarshal(body, &dashboard)
	if err != nil {
		return err
	}

	sensors.SetDashboardInfo(url, &dashboard)
	sensors.SetRequestInfo(url, dashboard.NodeID, requestInfo)

	for _, satellite := range dashboard.Satellites {
		err := mon.UpdateSNOSatellite(url, dashboard.NodeID, satellite.ID)
		if err != nil {
			log.Warnf("The satellite %s can't be updated for node %s", satellite.ID, url)
		}
	}

	return nil
}

func (mon *Monitor) UpdateSNOSatellite(url string, nodeId string, satelliteId string) error {
	start := time.Now()
	resp, err := http.Get(fmt.Sprintf("http://%s/api/sno/satellite/%s", url, satelliteId))
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	requestInfo := new(models.RequestInfo)
	requestInfo.Duration = time.Since(start)

	satellite := new(models.SatelliteData)
	err = json.Unmarshal(body, satellite)
	if err != nil {
		return err
	}

	sensors.SetSatelliteInfo(url, nodeId, satelliteId, *satellite)

	return nil
}

func (mon *Monitor) UpdateDashboard(url string) error {
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

func (mon *Monitor) UpdateSatellite(url string, nodeId string, satelliteId string) error {
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
