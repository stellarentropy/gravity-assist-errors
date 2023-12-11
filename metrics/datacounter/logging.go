package datacounter

import (
	"github.com/stellarentropy/gravity-assist-common/logging"

	"github.com/rs/zerolog"
)

// logger provides event logging capabilities with enhanced context specifically
// for the datacounter component, including preset contextual information to
// facilitate log categorization and filtering.
var logger zerolog.Logger

// init prepares the [logger] variable with a custom [zerolog.Logger] tailored
// for the datacounter component, appending a 'datacounter' identifier for event
// filtering. This setup occurs automatically before main execution.
func init() {
	logger = logging.GetLogger().With().Str("component", "datacounter").Logger()
}
