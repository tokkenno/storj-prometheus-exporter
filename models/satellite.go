package models

import "time"

type SatellitesData struct {
}

type SatellitesRequest struct {
	Data  SatellitesData `json:"data"`
	Error string         `json:"error"`
}

type DailyStorage struct {
	AtRestTotal   float64   `json:"atRestTotal"`
	IntervalStart time.Time `json:"intervalStart"`
}

type IngressBandwidth struct {
	Repair float64 `json:"repair"`
	Usage  float64 `json:"usage"`
}

type EgressBandwidth struct {
	Repair float64 `json:"repair"`
	Usage  float64 `json:"usage"`
	Audit  float64 `json:"audit"`
}

type DailyBandwidth struct {
	Egress        EgressBandwidth `json:"egress"`
	Ingress       EgressBandwidth `json:"ingress"`
	Delete        float64         `json:"delete"`
	IntervalStart time.Time       `json:"intervalStart"`
}

type SatelliteCheck struct {
	TotalCount   float64 `json:"totalCount"`
	SuccessCount float64 `json:"successCount"`
	Alpha        float64 `json:"alpha"`
	Beta         float64 `json:"beta"`
	Score        float64 `json:"score"`
}

type SatelliteData struct {
	Id               string           `json:"id"`
	StorageDaily     []DailyStorage   `json:"storageDaily"`
	BandwidthDaily   []DailyBandwidth `json:"bandwidthDaily"`
	StorageSummary   float64          `json:"storageSummary"`
	BandwidthSummary float64          `json:"bandwidthSummary"`
	EgressSummary    float64          `json:"egressSummary"`
	IngressSummary   float64          `json:"ingressSummary"`
	Audit            SatelliteCheck   `json:"audit"`
	Uptime           SatelliteCheck   `json:"uptime"`
}

type SatelliteRequest struct {
	Data  SatelliteData `json:"data"`
	Error string        `json:"error"`
}
