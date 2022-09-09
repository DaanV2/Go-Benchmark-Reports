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
func (w *MarkdownWriter) WriteString(text string) error {
	_, err := w.file.WriteString(text)
	return err
}

//Close closes the markdown file
func (w *MarkdownWriter) Close() error {
	return w.file.Close()
}

//WriteHeader writes a header to the markdown file
func (w *MarkdownWriter) WriteHeader(header string, level int) error {
	return w.WriteString(strings.Repeat("#", level) + " " + header)
}

//WriteReport writes a report to the markdown file
func (w *MarkdownWriter) WriteReport(r *Reporter) error {
	//Write title
	w.WriteHeader(r.Name, 1)

	//Information
	w.WriteHeader("Information", 2)
	for key, value := range r.Attributes {
		w.WriteString("```ini")
		if err := w.WriteString(key + "=" + value); err != nil {
			return err
		}
		w.WriteString("```")
	}

	//Results
	w.WriteHeader("Results", 2)
	table := CreateTable(r)

	w.WriteString("|" + strings.Join(table.Headers, "|") + "|")
	w.WriteString("|" + strings.Repeat("---|", len(table.Headers)))
	for _, row := range table.Rows {
		w.WriteString("|" + strings.Join(row.Values, "|") + "|")
	}

	return nil
}
