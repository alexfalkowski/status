package health

import "github.com/alexfalkowski/go-service/v2/time"

// Config configures the health subsystem.
//
// Both fields are parsed as durations (for example: "250ms", "5s", "1m").
//
// The values are used by the health module when registering checks and
// observers:
//
//   - Duration: the interval between health check executions.
//   - Timeout: the maximum amount of time a health check is allowed to run.
type Config struct {
	Duration time.Duration `yaml:"duration,omitempty" json:"duration,omitempty" toml:"duration,omitempty"`
	Timeout  time.Duration `yaml:"timeout,omitempty" json:"timeout,omitempty" toml:"timeout,omitempty"`
}
