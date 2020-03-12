package sensors

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/tokkenno/storj-prometheus-exporter/models"
	"time"
)

var (
	nodeIdSensor = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "storj_node_id",
		Help: "The client ID",
	}, []string{"node", "host"})
	isUpdateSensor = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "storj_node_updated",
		Help: "The client is updated",
	}, []string{"node", "host"})
	usedDiskspaceSensor = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "storj_node_used_diskspace",
		Help: "Storj used diskspace metrics",
	}, []string{"node", "host"})
	totalDiskspaceSensor = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "storj_node_total_diskspace",
		Help: "Storj total diskspace metrics",
	}, []string{"node", "host"})
	usedBandwidthSensor = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "storj_node_used_bandwidth",
		Help: "Storj used bandwidth metrics",
	}, []string{"node", "host"})
	totalBandwidthSensor = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "storj_node_total_bandwidth",
		Help: "Storj total bandwidth metrics",
	}, []string{"node", "host"})
	uptimeSensor = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "storj_node_uptime",
		Help: "Storj node uptime in seconds",
	}, []string{"node", "host"})
)

func SetDashboardInfo(hostId string, dashboardInfo *models.DashboardData) {
	nodeIdSensor.With(prometheus.Labels{"node": dashboardInfo.NodeID, "host": hostId}).Set(1.0)

	if dashboardInfo.UpToDate {
		isUpdateSensor.With(prometheus.Labels{"node": dashboardInfo.NodeID, "host": hostId}).Set(1.0)
	} else {
		isUpdateSensor.With(prometheus.Labels{"node": dashboardInfo.NodeID, "host": hostId}).Set(0.0)
	}

	usedBandwidthSensor.With(prometheus.Labels{"node": dashboardInfo.NodeID, "host": hostId}).Set(dashboardInfo.Bandwidth.Used)
	totalBandwidthSensor.With(prometheus.Labels{"node": dashboardInfo.NodeID, "host": hostId}).Set(dashboardInfo.Bandwidth.Available)
	usedDiskspaceSensor.With(prometheus.Labels{"node": dashboardInfo.NodeID, "host": hostId}).Set(dashboardInfo.DiskSpace.Used)
	totalDiskspaceSensor.With(prometheus.Labels{"node": dashboardInfo.NodeID, "host": hostId}).Set(dashboardInfo.DiskSpace.Available)
	uptimeSensor.With(prometheus.Labels{"node": dashboardInfo.NodeID, "host": hostId}).Set(time.Now().Sub(dashboardInfo.StartedAt).Seconds())
}

func SetDashboardOffline(hostId string, nodeId string) {
	nodeIdSensor.With(prometheus.Labels{"node": nodeId, "host": hostId}).Set(0.0)
	usedBandwidthSensor.With(prometheus.Labels{"node": nodeId, "host": hostId}).Set(0)
	totalBandwidthSensor.With(prometheus.Labels{"node": nodeId, "host": hostId}).Set(0)
	usedDiskspaceSensor.With(prometheus.Labels{"node": nodeId, "host": hostId}).Set(0)
	totalDiskspaceSensor.With(prometheus.Labels{"node": nodeId, "host": hostId}).Set(0)
	uptimeSensor.With(prometheus.Labels{"node": nodeId, "host": hostId}).Set(0)
}