package metrics

import "github.com/prometheus/client_golang/prometheus"

const (
	Prefix = "darmstadt_"
)

func init() {
	prometheus.MustRegister(connDurationsHisto)
	prometheus.MustRegister(connGauge)
	prometheus.MustRegister(errorCounterVec)
}
