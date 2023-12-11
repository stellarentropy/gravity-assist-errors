package logging

import (
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/diode"
	"github.com/stellarentropy/gravity-assist-common/config/common"
)

// GetLogger returns a preconfigured [zerolog.Logger] for application-wide
// logging, ensuring timestamped and formatted output with efficient message
// processing, including the ability to drop messages under high load to
// maintain system performance.
func GetLogger() Logger {
	// Set the global duration field unit to milliseconds
	zerolog.DurationFieldUnit = time.Millisecond
	// Set the global time field format to RFC3339
	zerolog.TimeFieldFormat = time.RFC3339

	// Declare a variable to hold the multi-level writer
	var multi zerolog.LevelWriter

	var wr diode.Writer

	logFormat := common_config.Common.LogFormat

	switch logFormat {
	case "color", "text":
		// Create a new diode writer with a buffer size of 1000 and a flush interval of 10 milliseconds
		wr = diode.NewWriter(consoleWriter(), 1000, 10*time.Millisecond, func(missed int) {
			// If any messages are dropped, print the number of dropped messages
			fmt.Printf("Dropped %d messages", missed)
		})
	case "json":
		// Create a new diode writer with a buffer size of 1000 and a flush interval of 10 milliseconds
		wr = diode.NewWriter(os.Stdout, 1000, 10*time.Millisecond, func(missed int) {
			// If any messages are dropped, print the number of dropped messages
			fmt.Printf("Dropped %d messages", missed)
		})
	}

	// Assign the diode writer to the multi-level writer
	multi = zerolog.MultiLevelWriter(wr)

	// Create a new logger with the multi-level writer, add a timestamp to each log message
	logger := zerolog.New(multi).With().Timestamp().Logger()

	// Set the global log level to Info
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	// Return the logger
	return Logger{logger}
}

// consoleWriter creates and returns a [zerolog.ConsoleWriter] that formats log
// messages for display in the console, incorporating features like color
// coding, time stamps, caller information, log levels, and message content
// arranged in a readable tabular format with fixed message width.
func consoleWriter() zerolog.ConsoleWriter {
	logFormat := common_config.Common.LogFormat

	var color bool

	if logFormat == "color" {
		color = true
	}

	// Create a new console writer that outputs to the standard output
	writer := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "15:04:05", NoColor: !color}

	// Set the order of the parts in the log message
	writer.PartsOrder = []string{
		zerolog.TimestampFieldName,
		zerolog.CallerFieldName,
		zerolog.LevelFieldName,
		zerolog.MessageFieldName,
	}

	// Define the format of the message part
	writer.FormatMessage = func(i interface{}) string {
		// Format the message to have a fixed width of 60 characters
		return fmt.Sprintf("| %-60s|", i)
	}

	// Return the configured console writer
	return writer
}
