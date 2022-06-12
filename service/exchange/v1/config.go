package exchange

import (
	"fmt"
	"time"
)

// Config keeps Service params.
type Config struct {
	UpdaterTimeout           time.Duration `mapstructure:"updater_timeout"`
	BtcUsdtRateCheckInterval time.Duration `mapstructure:"btc_usdt_rate_check_interval"`
	CurRubRateCheckInterval  time.Duration `mapstructure:"cur_rub_rate_check_interval"`
}

// Validate performs a basic validation.
func (config Config) Validate() error {
	if config.UpdaterTimeout < time.Second {
		return fmt.Errorf("updater_timeout field: too short period")
	}

	if config.BtcUsdtRateCheckInterval < time.Second {
		return fmt.Errorf("btc_usdt_rate_check_interval field: too short period")
	}

	if config.CurRubRateCheckInterval < time.Second {
		return fmt.Errorf("cur_rub_rate_check_interval field: too short period")
	}

	return nil
}

// NewDefaultConfig builds a Config with default values.
func NewDefaultConfig() Config {
	return Config{
		UpdaterTimeout:           5 * time.Second,
		BtcUsdtRateCheckInterval: 10 * time.Second,
		CurRubRateCheckInterval:  24 * time.Hour,
	}
}
