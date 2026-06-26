package domain

import "context"

// Collector defines the metric storage backend contract.
type Collector interface {
	Collect(ctx context.Context, point MetricPoint) error
}
