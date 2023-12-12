package common_config

import (
	"time"

	"github.com/stellarentropy/gravity-assist-common/config"
)

// CommonConfig holds the operational settings necessary for an agent to function
// correctly in its environment. It encompasses parameters for network
// communication, management of concurrent processes, identification of
// resources, integration with services, and security protocols. The structure
// facilitates the agent's ability to adapt to various deployment contexts by
// utilizing environmental variables to customize its behavior accordingly.
type CommonConfig struct {
	ServiceName string

	EnableMetricCollection bool
	EnableTraceCollection  bool

	GoogleProjectId string

	DebugListenAddress string
	DebugListenPort    int
	DebugReadTimeout   time.Duration
	DebugWriteTimeout  time.Duration

	MetricsListenAddress string
	MetricsListenPort    int
	MetricsReadTimeout   time.Duration
	MetricsWriteTimeout  time.Duration

	HealthListenAddress string
	HealthListenPort    int
	HealthReadTimeout   time.Duration
	HealthWriteTimeout  time.Duration

	GracefulShutdownTimeout time.Duration

	LogFormat string
}

// Common holds the configuration for an agent, encapsulating settings necessary
// for network communication, worker pool management, storage access, and
// service endpoint interactions, and includes security parameters to ensure
// safe operations within its environment.
var Common = &CommonConfig{
	ServiceName: config.NewEnv("SE_GA_SERVICE_NAME").
		WithDefault("gravity-assist-common").
		WithRequired().
		GetString(),

	EnableMetricCollection: config.NewEnv("SE_GA_ENABLE_METRIC_COLLECTION").
		WithDefault("true").
		WithRequired().
		GetBool(),

	EnableTraceCollection: config.NewEnv("SE_GA_ENABLE_TRACE_COLLECTION").
		WithDefault("true").
		WithRequired().
		GetBool(),

	GoogleProjectId: config.NewEnv("SE_GA_PROJECT_ID").
		WithDefault("gravity-assist").
		WithRequired().
		GetString(),

	// region Debug
	DebugListenAddress: config.NewEnv("SE_GA_DEBUG_LISTEN_ADDRESS").
		WithDefault("127.0.0.1").
		WithRequired().
		GetAddress(),

	DebugListenPort: config.NewEnv("SE_GA_DEBUG_LISTEN_PORT").
		WithDefault("0").
		WithRequired().
		GetPort(),

	DebugReadTimeout: config.NewEnv("SE_GA_DEBUG_READ_TIMEOUT").
		WithDefault("60s").
		WithRequired().
		GetDuration(),

	DebugWriteTimeout: config.NewEnv("SE_GA_DEBUG_WRITE_TIMEOUT").
		WithDefault("60s").
		WithRequired().
		GetDuration(),
	// endregion

	// region Metrics
	MetricsListenAddress: config.NewEnv("SE_GA_METRICS_LISTEN_ADDRESS").
		WithDefault("0.0.0.0").
		WithRequired().
		GetAddress(),

	MetricsListenPort: config.NewEnv("SE_GA_METRICS_LISTEN_PORT").
		WithDefault("9090").
		WithRequired().
		GetPort(),

	MetricsReadTimeout: config.NewEnv("SE_GA_METRICS_READ_TIMEOUT").
		WithDefault("60s").
		WithRequired().
		GetDuration(),

	MetricsWriteTimeout: config.NewEnv("SE_GA_METRICS_WRITE_TIMEOUT").
		WithDefault("60s").
		WithRequired().
		GetDuration(),
	// endregion

	// region Health
	HealthListenAddress: config.NewEnv("SE_GA_HEALTH_LISTEN_ADDRESS").
		WithDefault("0.0.0.0").
		WithRequired().
		GetAddress(),

	HealthListenPort: config.NewEnv("SE_GA_HEALTH_LISTEN_PORT").
		WithDefault("1234").
		WithRequired().
		GetPort(),

	HealthReadTimeout: config.NewEnv("SE_GA_HEALTH_READ_TIMEOUT").
		WithDefault("60s").
		WithRequired().
		GetDuration(),

	HealthWriteTimeout: config.NewEnv("SE_GA_HEALTH_WRITE_TIMEOUT").
		WithDefault("60s").
		WithRequired().
		GetDuration(),
	// endregion

	GracefulShutdownTimeout: config.NewEnv("SE_GA_HEALTH_GRACEFUL_SHUTDOWN_TIMEOUT").
		WithDefault("60s").
		WithRequired().
		GetDuration(),

	LogFormat: config.NewEnv("SE_GA_LOG_FORMAT").
		WithDefault("color").
		WithOptions("text", "color", "json").
		WithRequired().
		GetString(),
}
