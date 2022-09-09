package report

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ParseBenchmark(t *testing.T) {
	t.Run("With metrics", func(t *testing.T) {
		parser := &TextParser{}
		benchmark := parser.ParseBenchmark("Benchmark_Name/Subname         	  260517	     23746 ns/op	       100.0 N	         0.0003839 N/op	     616 B/op	      11 allocs/op")

		assert.NotNil(t, benchmark)
		assert.Equal(t, "Benchmark_Name/Subname", benchmark.Name)
		assert.Equal(t, uint64(260517), benchmark.Operations)
		assert.Equal(t, 5, len(benchmark.Metrics))

		if len(benchmark.Metrics) < 5 {
			return
		}

		assert.Equal(t, "23746", benchmark.Metrics[0].Value)
		assert.Equal(t, "ns/op", benchmark.Metrics[0].Unit)

		assert.Equal(t, "100.0", benchmark.Metrics[1].Value)
		assert.Equal(t, "N", benchmark.Metrics[1].Unit)

		assert.Equal(t, "0.0003839", benchmark.Metrics[2].Value)
		assert.Equal(t, "N/op", benchmark.Metrics[2].Unit)

		assert.Equal(t, "616", benchmark.Metrics[3].Value)
		assert.Equal(t, "B/op", benchmark.Metrics[3].Unit)

		assert.Equal(t, "11", benchmark.Metrics[4].Value)
		assert.Equal(t, "allocs/op", benchmark.Metrics[4].Unit)
	})
}

func Test_ParseMetric(t *testing.T) {
	t.Run("With unit", func(t *testing.T) {
		parser := &TextParser{}
		metric := parser.ParseMetric("123 ns/op")

		assert.NotNil(t, metric)
		assert.Equal(t, "123", metric.Value)
		assert.Equal(t, "ns/op", metric.Unit)
	})

	t.Run("Without unit", func(t *testing.T) {
		parser := &TextParser{}
		metric := parser.ParseMetric("123")

		assert.NotNil(t, metric)
		assert.Equal(t, "123", metric.Value)
		assert.Equal(t, "", metric.Unit)
	})
}
