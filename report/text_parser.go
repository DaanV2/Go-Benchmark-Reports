package report

import (
	"os"
	"strconv"
	"strings"
)

type TextParser struct {
}

//ParseFile parses a file and returns a Reporter
func (p *TextParser) ParseFile(filepath string, report *Reporter) error {
	dat, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}

	text := string(dat)
	lines := strings.Split(text, "\n")
	startedBenchmark := false

	for _, line := range lines {
		if strings.HasPrefix(line, "Benchmark") {
			startedBenchmark = true
			report.AddBenchmark(p.ParseBenchmark(line))
		} else if !startedBenchmark {
			if strings.Contains(line, ":") {
				index := strings.Index(line, ":")
				attr := strings.TrimSpace(line[:index])
				value := strings.TrimSpace(line[index+1:])
				report.Attributes[attr] = value
			}
		}
	}

	return nil
}

//ParseBenchmark parses a benchmark line and returns a Benchmark
func (p *TextParser) ParseBenchmark(line string) *Benchmark {
	parts := strings.Split(line, "\t")
	//0 Is the name, 1 the operations
	if len(parts) < 2 {
		return nil
	}

	result := &Benchmark{
		Name: strings.TrimSpace(parts[0]),
	}

	if value, err := strconv.ParseUint(strings.TrimSpace(parts[1]), 10, 64); err != nil {
		return nil
	} else {
		result.Operations = value
	}

	for _, part := range parts[2:] {
		result.AddMetric(p.ParseMetric(strings.TrimSpace(part)))
	}

	return result
}

//ParseMetric parses a metric line and returns a Metric
func (p *TextParser) ParseMetric(line string) *Metric {
	parts := strings.Split(line, " ")
	result := &Metric{
		Value: strings.TrimSpace(parts[0]),
		Unit:  "",
	}

	if len(parts) > 1 {
		result.Unit = strings.TrimSpace(parts[1])
	}

	return result
}
