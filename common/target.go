package common

type Target interface {
	GetApiName() string
	GenMonitor(config MonitorConfig) Monitor
}
