package report

// Benchmark represents a benchmark
type Benchmark struct {
	// Name of the benchmark
	Name string `json:"name"`
	// Number of operations executed
	Operations uint64 `json:"operations"`
	// The metrics
	Metrics []*Metric `json:"metrics"`
}

// Metric represents a metric of a benchmark
type Metric struct {
	Value string `json:"value"`
	Unit  string `json:"unit,omitempty"`
}

func (b *Benchmark) AddMetric(metric *Metric) {
	if metric != nil {
		b.Metrics = append(b.Metrics, metric)
	}
}
