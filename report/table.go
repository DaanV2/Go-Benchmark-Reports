package report

import "strconv"

type Table struct {
	Headers []string
	Rows    []*TableRow
}

func NewTable() *Table {
	return &Table{
		Headers: make([]string, 0),
		Rows:    make([]*TableRow, 0),
	}
}

func (t *Table) SetHeader(header string, index int) *Table {
	for len(t.Headers) <= index {
		t.Headers = append(t.Headers, "")
	}
	t.Headers[index] = header
	return t
}

func (t *Table) GetOrAddHeader(value string) int {
	for i, header := range t.Headers {
		if header == value {
			return i
		}
	}

	t.Headers = append(t.Headers, value)
	return len(t.Headers) - 1
}

func (t *Table) AddRow(row *TableRow) *Table {
	t.Rows = append(t.Rows, row)
	return t
}

func (t *Table) SetRow(row *TableRow, index int) *Table {
	for len(t.Rows) <= index {
		t.Rows = append(t.Rows, nil)
	}
	t.Rows[index] = row
	return t
}

func (t *Table) GetRow(index int) *TableRow {
	for len(t.Rows) <= index {
		t.Rows = append(t.Rows, nil)
	}

	return t.Rows[index]
}

//SetCell sets a cell at the given index
func (t *Table) SetCell(cell string, row, column int) *Table {
	t.GetRow(row).SetValue(cell, column)
	return t
}

type TableRow struct {
	Values []string
}

func NewTableRow() *TableRow {
	return &TableRow{
		Values: make([]string, 0),
	}
}

//AddValue adds a value to the row
func (r *TableRow) AddValue(value string) *TableRow {
	r.Values = append(r.Values, value)
	return r
}

//SetValue sets a value at the given index
func (r *TableRow) SetValue(value string, index int) *TableRow {
	for len(r.Values) <= index {
		r.Values = append(r.Values, "")
	}
	r.Values[index] = value
	return r
}

//GetValue gets a value at the given index
func (r *TableRow) GetValue(index int) string {
	for len(r.Values) <= index {
		r.Values = append(r.Values, "")
	}

	return r.Values[index]
}

func CreateTable(report *Reporter) *Table {
	table := &Table{}
	table.SetHeader("Name", 0)
	table.SetHeader("Operations", 1)

	for _, benchMark := range report.Benchmarks {
		row := NewTableRow()
		table.AddRow(row)
		row.SetValue(benchMark.Name, 0)
		row.SetValue(strconv.FormatUint(benchMark.Operations, 10), 1)

		for _, metric := range benchMark.Metrics {
			cell := table.GetOrAddHeader(metric.Unit)
			row.SetValue(metric.Value, cell)
		}
	}

	return table
}
