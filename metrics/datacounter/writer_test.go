package datacounter

import (
	"bytes"
	"io"
	"testing"
)

type nopWriteCloser struct {
	io.Writer
}

func (nopWriteCloser) Close() error { return nil }

func TestWriterCounter(t *testing.T) {
	data := []byte("Hello, World!")
	writer := NewWriterCounter(nopWriteCloser{bytes.NewBuffer(data)})

	n, err := writer.Write(data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if n != len(data) {
		t.Fatalf("expected to write %d bytes, got %d", len(data), n)
	}

	if writer.Count() != uint64(len(data)) {
		t.Fatalf("expected count to be %d, got %d", len(data), writer.Count())
	}

	if err := writer.Close(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
