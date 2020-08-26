package common

type MonitorConfig struct {
	Api      string `json:"api"`
	Url      string `json:"url"`
	Interval int    `json:"interval"`
}

type Config struct {
	Port     int             `json:"port"`
	Monitors []MonitorConfig `json:"monitors"`
}
