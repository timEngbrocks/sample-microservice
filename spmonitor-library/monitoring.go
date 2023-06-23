package monitoring

import (
	"io"
	"net/http"
)

var handler = MetricHandler{
	registry: MetricRegistry{
		metrics: make([]Metric, 0),
	},
}

func GetMetricHandler() *MetricHandler {
	return &handler
}

func (h *MetricHandler) Metrics(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, h.registry.ExportAll())
}

type MetricHandler struct {
	registry MetricRegistry
}

type MetricDescriptor struct {
	Name string
	Help string
}

type Metric interface {
	Export() string
}

type MetricRegistry struct {
	metrics []Metric
}

func (r *MetricRegistry) AddMetric(metric Metric) {
	r.metrics = append(r.metrics, metric)
}

func (r MetricRegistry) ExportAll() string {
	var result = ""
	for _, metric := range r.metrics {
		result += metric.Export() + "\n"
	}
	if len(result) > 0 {
		result = result[:len(result)-1]
	}
	return result
}
