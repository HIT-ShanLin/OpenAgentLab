// Package domain defines the core entities for the observability module.
//
// Observability provides metrics, tracing, and structured logging
// across all OpenAgentLab services.
package domain

// MetricPoint represents a single observation for Prometheus-compatible metrics.
type MetricPoint struct {
	Name   string
	Labels map[string]string
	Value  float64
}

// TraceSpan represents a span in a distributed trace (OpenTelemetry-compatible).
type TraceSpan struct {
	TraceID    string
	SpanID     string
	ParentID   string
	Name       string
	Attributes map[string]string
}

// LogEntry represents a structured log record.
type LogEntry struct {
	Level   string
	Message string
	Fields  map[string]interface{}
	TraceID string
}
