package report

import (
	"os"
	"strings"
)

//MarkdownWriter is a markdown writer
type MarkdownWriter struct {
	file StringWriterCloser
}

//NewMarkdownWriter creates a new markdown writer
func NewMarkdownWriter(filepath string) (*MarkdownWriter, error) {
	file, err := os.Create(filepath)
	if err != nil {
		return nil, err
	}

	return &MarkdownWriter{file: file}, nil
}

//WriteString writes a string to the markdown file
func (w *MarkdownWriter) WriteLine(text string) error {
	_, err := w.file.WriteString(text + "\n")
	return err
}

//Close closes the markdown file
func (w *MarkdownWriter) Close() error {
	return w.file.Close()
}

//WriteHeader writes a header to the markdown file
func (w *MarkdownWriter) WriteHeader(header string, level int) error {
	return w.WriteLine(strings.Repeat("#", level) + " " + header)
}

//WriteReport writes a report to the markdown file
func (w *MarkdownWriter) WriteReport(r *Reporter) error {
	//Write title
	w.WriteHeader(r.Name, 1)
	w.WriteLine("")

	//Information
	w.WriteHeader("Information", 2)
	w.WriteAttributes(r.Attributes)
	w.WriteLine("")

	//Results
	w.WriteHeader("Results", 2)
	table := CreateTable(r)
	return w.WriteTable(table)
}

func (w *MarkdownWriter) WriteAttributes(attributes map[string]string) error {
	for key, value := range attributes {
		w.WriteLine("```ini")
		if err := w.WriteLine(key + "=" + value); err != nil {
			return err
		}
		w.WriteLine("```")
	}

	return nil
}

func (w *MarkdownWriter) WriteTable(table *Table) error {
	w.WriteLine("|" + strings.Join(table.Headers, "|") + "|")
	w.WriteLine("|" + strings.Repeat("---|", len(table.Headers)))
	for _, row := range table.Rows {
		w.WriteLine("|" + strings.Join(row.Values, "|") + "|")
	}

	return nil
}
