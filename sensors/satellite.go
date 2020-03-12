package sensors

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/tokkenno/storj-prometheus-exporter/models"
	"time"
)

var (
	satSummarySensor = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "storj_node_sat_summary",
		Help: "Storj satellite summary metrics",
	}, []string{"node", "host", "satellite", "type"})
	satAuditSensor = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "storj_node_sat_audit",
		Help: "Storj satellite audit metrics",
	}, []string{"node", "host", "satellite", "type"})
	satUptimeSensor = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "storj_node_sat_uptime",
		Help: "Storj satellite uptime metrics",
	}, []string{"node", "host", "satellite", "type"})
	satMonthEgressSensor = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "storj_node_sat_month_egress",
		Help: "Storj satellite egress since current month start",
	}, []string{"node", "host", "satellite", "type"})
	satMonthIngressSensor = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "storj_node_sat_month_ingress",
		Help: "Storj satellite ingress since current month start",
	}, []string{"node", "host", "satellite", "type"})
	satDayEgressSensor = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "storj_node_sat_day_egress",
		Help: "Storj satellite egress since current day start",
	}, []string{"node", "host", "satellite", "type"})
	satDayIngressSensor = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "storj_node_sat_day_ingress",
		Help: "Storj satellite ingress since current day start",
	}, []string{"node", "host", "satellite", "type"})
	satMonthStorageSensor = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "storj_node_sat_month_storage",
		Help: "Storj satellite data stored on disk since current month start",
	}, []string{"node", "host", "satellite"})
	satDayStorageSensor = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "storj_node_sat_day_storage",
		Help: "Storj satellite data stored on disk since current day start",
	}, []string{"node", "host", "satellite"})
)

func SetSatelliteInfo(hostId string, nodeId string, satelliteId string, satelliteInfo models.SatelliteData) {
	satSummarySensor.
		With(prometheus.Labels{"node": nodeId, "host": hostId, "satellite": satelliteId, "type": "storage"}).
		Set(satelliteInfo.StorageSummary)
	satSummarySensor.
		With(prometheus.Labels{"node": nodeId, "host": hostId, "satellite": satelliteId, "type": "bandwidth"}).
		Set(satelliteInfo.BandwidthSummary)
	satSummarySensor.
		With(prometheus.Labels{"node": nodeId, "host": hostId, "satellite": satelliteId, "type": "egress"}).
		Set(satelliteInfo.EgressSummary)
	satSummarySensor.
		With(prometheus.Labels{"node": nodeId, "host": hostId, "satellite": satelliteId, "type": "ingress"}).
		Set(satelliteInfo.IngressSummary)

	satAuditSensor.
		With(prometheus.Labels{"node": nodeId, "host": hostId, "satellite": satelliteId, "type": "alpha"}).
		Set(satelliteInfo.Audit.Alpha)
	satAuditSensor.
		With(prometheus.Labels{"node": nodeId, "host": hostId, "satellite": satelliteId, "type": "beta"}).
		Set(satelliteInfo.Audit.Beta)
	satAuditSensor.
		With(prometheus.Labels{"node": nodeId, "host": hostId, "satellite": satelliteId, "type": "score"}).
		Set(satelliteInfo.Audit.Score)
	satAuditSensor.
		With(prometheus.Labels{"node": nodeId, "host": hostId, "satellite": satelliteId, "type": "successCount"}).
		Set(satelliteInfo.Audit.SuccessCount)
	satAuditSensor.
		With(prometheus.Labels{"node": nodeId, "host": hostId, "satellite": satelliteId, "type": "totalCount"}).
		Set(satelliteInfo.Audit.TotalCount)

	satUptimeSensor.
		With(prometheus.Labels{"node": nodeId, "host": hostId, "satellite": satelliteId, "type": "alpha"}).
		Set(satelliteInfo.Uptime.Alpha)
	satUptimeSensor.
		With(prometheus.Labels{"node": nodeId, "host": hostId, "satellite": satelliteId, "type": "beta"}).
		Set(satelliteInfo.Uptime.Beta)
	satUptimeSensor.
		With(prometheus.Labels{"node": nodeId, "host": hostId, "satellite": satelliteId, "type": "score"}).
		Set(satelliteInfo.Uptime.Score)
	satUptimeSensor.
		With(prometheus.Labels{"node": nodeId, "host": hostId, "satellite": satelliteId, "type": "successCount"}).
		Set(satelliteInfo.Uptime.SuccessCount)
	satUptimeSensor.
		With(prometheus.Labels{"node": nodeId, "host": hostId, "satellite": satelliteId, "type": "totalCount"}).
		Set(satelliteInfo.Uptime.TotalCount)

	monthStorage := float64(0)
	lastDayIndex := 0
	lastDayDate := time.Time{}
	for dayIndex, day := range satelliteInfo.StorageDaily {
		monthStorage += day.AtRestTotal
		if day.IntervalStart.After(lastDayDate) {
			lastDayDate = day.IntervalStart
			lastDayIndex = dayIndex
		}
	}

	satMonthStorageSensor.
		With(prometheus.Labels{"node": nodeId, "host": hostId, "satellite": satelliteId}).
		Set(monthStorage)
	satDayStorageSensor.
		With(prometheus.Labels{"node": nodeId, "host": hostId, "satellite": satelliteId}).
		Set(satelliteInfo.StorageDaily[lastDayIndex].AtRestTotal)

	monthBandwidth := new(models.DailyBandwidth)
	lastDayIndex = 0
	lastDayDate = time.Time{}
	for dayIndex, day := range satelliteInfo.BandwidthDaily {
		monthBandwidth.Add(&day)
		if day.IntervalStart.After(lastDayDate) {
			lastDayDate = day.IntervalStart
			lastDayIndex = dayIndex
		}
	}

	satMonthEgressSensor.
		With(prometheus.Labels{"node": nodeId, "host": hostId, "satellite": satelliteId, "type": "audit"}).
		Set(monthBandwidth.Egress.Audit)
	satMonthEgressSensor.
		With(prometheus.Labels{"node": nodeId, "host": hostId, "satellite": satelliteId, "type": "repair"}).
		Set(monthBandwidth.Egress.Repair)
	satMonthEgressSensor.
		With(prometheus.Labels{"node": nodeId, "host": hostId, "satellite": satelliteId, "type": "usage"}).
		Set(monthBandwidth.Egress.Usage)
	satMonthIngressSensor.
		With(prometheus.Labels{"node": nodeId, "host": hostId, "satellite": satelliteId, "type": "repair"}).
		Set(monthBandwidth.Ingress.Repair)
	satMonthIngressSensor.
		With(prometheus.Labels{"node": nodeId, "host": hostId, "satellite": satelliteId, "type": "usage"}).
		Set(monthBandwidth.Ingress.Usage)

	satDayEgressSensor.
		With(prometheus.Labels{"node": nodeId, "host": hostId, "satellite": satelliteId, "type": "audit"}).
		Set(satelliteInfo.BandwidthDaily[lastDayIndex].Egress.Audit)
	satDayEgressSensor.
		With(prometheus.Labels{"node": nodeId, "host": hostId, "satellite": satelliteId, "type": "repair"}).
		Set(satelliteInfo.BandwidthDaily[lastDayIndex].Egress.Repair)
	satDayEgressSensor.
		With(prometheus.Labels{"node": nodeId, "host": hostId, "satellite": satelliteId, "type": "usage"}).
		Set(satelliteInfo.BandwidthDaily[lastDayIndex].Egress.Usage)
	satDayIngressSensor.
		With(prometheus.Labels{"node": nodeId, "host": hostId, "satellite": satelliteId, "type": "repair"}).
		Set(satelliteInfo.BandwidthDaily[lastDayIndex].Ingress.Repair)
	satDayIngressSensor.
		With(prometheus.Labels{"node": nodeId, "host": hostId, "satellite": satelliteId, "type": "usage"}).
		Set(satelliteInfo.BandwidthDaily[lastDayIndex].Ingress.Usage)
}