package report

import (
	"errors"
	"path"
)

//Reporter is the main report object
type Reporter struct {
	Attributes map[string]string `json:"attributes,omitempty"`
	Benchmarks []*Benchmark      `json:"benchmarks"`
	Name       string            `json:"name"`
}

//NewReporter creates a new Reporter
func NewReport() *Reporter {
	return &Reporter{
		Attributes: make(map[string]string),
		Benchmarks: make([]*Benchmark, 0),
		Name:       "",
	}
}

//ParseFile parses a file and returns a Reporter
func (r *Reporter) ParseFile(filepath string) error {
	ext := path.Ext(filepath)
	r.Name = path.Base(filepath)

	switch ext {
	case ".text":
		parser := &TextParser{}
		return parser.ParseFile(filepath, r)
	}

	return errors.New("Unknown file extension: " + ext)
}

//AddBenchmark adds a benchmark to the report
func (r *Reporter) AddBenchmark(b *Benchmark) {
	if b != nil {
		r.Benchmarks = append(r.Benchmarks, b)
	}
}
