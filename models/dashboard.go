package models

import "time"

type Satellite struct {
	ID           string `json:"id"`
	Url          string `json:"url"`
	Disqualified string `json:"disqualified"`
}

type SizeQuota struct {
	Used      float64 `json:"used"`
	Available float64 `json:"available"`
}

type DashboardData struct {
	NodeID                string      `json:"nodeID"`
	Wallet                string      `json:"wallet"`
	Satellites            []Satellite `json:"satellites"`
	DiskSpace             SizeQuota   `json:"diskSpace"`
	Bandwidth             SizeQuota   `json:"bandwidth"`
	LastPinged            time.Time   `json:"lastPinged"`
	LastPingedFromID      time.Time   `json:"lastPingedFromID"`
	LastPingedFromAddress time.Time   `json:"lastPingedFromAddress"`
	Version               string      `json:"version"`
	AllowedVersion        string      `json:"allowedVersion"`
	UpToDate              bool        `json:"upToDate"`
	StartedAt             time.Time   `json:"startedAt"`
}

type DashboardRequest struct {
	Data  DashboardData `json:"data"`
	Error string        `json:"error"`
}