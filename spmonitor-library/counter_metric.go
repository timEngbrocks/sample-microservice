package monitoring

import "fmt"

type CounterMetric struct {
	descriptor MetricDescriptor
	value      int
}

func NewCounterMetric(descriptor MetricDescriptor) *CounterMetric {
	metric := CounterMetric{
		descriptor: descriptor,
		value:      0,
	}
	handler.registry.AddMetric(&metric)
	return &metric
}

func (m CounterMetric) Export() string {
	return fmt.Sprintf(
		"# HELP %s %s\n# TYPE %s counter\n%s %d\n",
		m.descriptor.Name,
		m.descriptor.Help,
		m.descriptor.Name,
		m.descriptor.Name,
		m.value,
	)
}

func (m *CounterMetric) Increment() {
	m.value++
}
