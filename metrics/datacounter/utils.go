package datacounter

import (
	"io"
)

// CloseWriters closes all provided [io.Closer] instances and returns the first
// encountered error, if any.
func CloseWriters(writers []io.Closer) error {
	for _, w := range writers {
		if err := w.Close(); err != nil {
			logger.Error().Err(err).Msg("failed to close writer")
			return err
		}
	}

	return nil
}
