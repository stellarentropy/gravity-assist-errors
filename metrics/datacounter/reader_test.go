package datacounter

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestReaderCounter(t *testing.T) {
	data := []byte("Hello, World!")
	reader := NewReaderCounter(ioutil.NopCloser(bytes.NewReader(data)))

	buf := make([]byte, len(data))
	n, err := reader.Read(buf)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if n != len(data) {
		t.Fatalf("expected to read %d bytes, got %d", len(data), n)
	}

	if reader.Count() != uint64(len(data)) {
		t.Fatalf("expected count to be %d, got %d", len(data), reader.Count())
	}

	if err := reader.Close(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
