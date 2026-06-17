package config

import (
	"github.com/alexfalkowski/go-service/v2/config"
	"github.com/alexfalkowski/go-service/v2/runtime"
	"github.com/alexfalkowski/go-service/v2/time"
	"github.com/alexfalkowski/status/internal/health"
)

// MaxSleepLimit is the largest configured status response delay.
const MaxSleepLimit = 5 * time.Minute

// Config is the service configuration loaded by the server command.
//
// Health and the embedded shared service config are required. MaxSleep is
// optional; zero falls back to MaxSleepLimit, and positive values must not
// exceed MaxSleepLimit.
type Config struct {
	Health         *health.Config `yaml:"health,omitempty" json:"health,omitempty" toml:"health,omitempty" validate:"required"`
	*config.Config `yaml:",inline" json:",inline" toml:",inline" validate:"required"`
	MaxSleep       time.Duration `yaml:"max_sleep,omitempty" json:"max_sleep,omitempty" toml:"max_sleep,omitempty" validate:"max_sleep"`
}

// GetMaxSleep returns the configured max sleep, falling back to MaxSleepLimit.
func (c *Config) GetMaxSleep() time.Duration {
	if c == nil || c.MaxSleep == 0 {
		return MaxSleepLimit
	}

	return c.MaxSleep
}

func decorateConfig(cfg *Config) *config.Config {
	return cfg.Config
}

func healthConfig(cfg *Config) *health.Config {
	return cfg.Health
}

func decorateValidator(v *config.Validator) *config.Validator {
	runtime.Must(v.RegisterValidation("max_sleep", validateMaxSleep))
	return v
}

func validateMaxSleep(fl config.FieldLevel) bool {
	maxSleep := time.Duration(fl.Field().Int())
	return maxSleep == 0 || (maxSleep > 0 && maxSleep <= MaxSleepLimit)
}
