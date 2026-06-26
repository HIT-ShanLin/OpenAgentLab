package service

import (
	"context"

	"github.com/HIT-ShanLin/OpenAgentLab/server/internal/observability/domain"
)

// MetricsService provides Prometheus-compatible metric collection.
type MetricsService struct {
	collector domain.Collector
}

// New creates a new MetricsService.
func New(c domain.Collector) *MetricsService {
	return &MetricsService{collector: c}
}

// Record stores a metric observation.
func (s *MetricsService) Record(ctx context.Context, name string, labels map[string]string, value float64) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	point := domain.MetricPoint{Name: name, Labels: labels, Value: value}
	return s.collector.Collect(ctx, point)
}
