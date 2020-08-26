package storj

import "storj-prometheus-exporter/common"

type Target struct {
}

func (t Target) GetApiName() string {
	return "storj"
}

func (t Target) GenMonitor(config common.MonitorConfig) common.Monitor {
	return &Monitor{
		config: config,
	}
}
