package report

import (
	"os"
	"strings"
)

type CommaSeperateValuesWriter struct {
	file StringWriterCloser
}

//NewCommaSeperateValuesWriter
func NewCommaSeperateValuesWriter(filepath string) (*CommaSeperateValuesWriter, error) {
	file, err := os.Create(filepath)
	if err != nil {
		return nil, err
	}

	return &CommaSeperateValuesWriter{file: file}, nil
}

//Close closes the csv file
func (w *CommaSeperateValuesWriter) Close() error {
	return w.file.Close()
}

//WriteString writes a string to the markdown file
func (w *CommaSeperateValuesWriter) WriteLine(text string) error {
	_, err := w.file.WriteString(text + "\n")
	return err
}

func (w *CommaSeperateValuesWriter) WriteRow(items ...string) error {
	values := make([]string, len(items))

	for i, value := range items {
		if strings.Contains(value, ",") {
			values[i] = "\"" + value + "\""
		} else {
			values[i] = value
		}
	}

	return w.WriteLine(strings.Join(values, ",") + "\n")
}

//WriteHeader writes a header to the markdown file
func (w *CommaSeperateValuesWriter) WriteHeader(header string, level int) error {
	return w.WriteRow(strings.Repeat("#", level) + " " + header)
}

//WriteReport writes a report to the markdown file
func (w *CommaSeperateValuesWriter) WriteReport(r *Reporter) error {
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

func (w *CommaSeperateValuesWriter) WriteAttributes(attributes map[string]string) error {
	for key, value := range attributes {
		if err := w.WriteRow(key, value); err != nil {
			return err
		}
	}

	return nil
}

func (w *CommaSeperateValuesWriter) WriteTable(table *Table) error {
	w.WriteRow(table.Headers...)
	for _, row := range table.Rows {
		w.WriteRow(row.Values...)
	}

	return nil
}
