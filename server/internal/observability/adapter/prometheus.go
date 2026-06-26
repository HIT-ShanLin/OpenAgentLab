package adapter

import (
	"context"
	"log"

	"github.com/HIT-ShanLin/OpenAgentLab/server/internal/observability/domain"
)

// PrometheusCollector implements domain.Collector using Prometheus.
// Stub: will register counters/gauges/histograms via prometheus client.
type PrometheusCollector struct {
	Port int
}

// NewPrometheusCollector creates a new Prometheus-backed collector.
func NewPrometheusCollector(port int) *PrometheusCollector {
	return &PrometheusCollector{Port: port}
}

// Collect records a metric point.
func (c *PrometheusCollector) Collect(ctx context.Context, point domain.MetricPoint) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	log.Printf("[observability] metric: %s{%v} = %f", point.Name, point.Labels, point.Value)
	// TODO: register and update Prometheus metrics
	return nil
}
