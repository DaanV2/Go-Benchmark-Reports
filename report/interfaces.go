package report

import "io"

type StringWriterCloser interface {
	io.StringWriter
	io.Closer
}
