package sensors

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"storj-prometheus-exporter/miners/storj/models"
)

var (
	requestTimeSensor = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "storj_node_request_time",
		Help: "Duration of request to Storj node. Can be used to detect network problems.",
	}, []string{"node", "host"})
)

func SetRequestInfo(hostId string, nodeId string, requestInfo *models.RequestInfo) {
	requestTimeSensor.With(prometheus.Labels{"node": nodeId, "host": hostId}).Set(float64(requestInfo.Duration.Milliseconds()))
}