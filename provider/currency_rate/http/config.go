package currencyRate

import (
	"fmt"
)

// Config keeps Accrual params.
type Config struct {
	BtcUsdtRateURL string `mapstructure:"btc_usdt_rate_address"`
	CurRubRateURL  string `mapstructure:"cur_rub_rate_address"`
}

// Validate performs a basic validation.
func (config Config) Validate() error {
	if config.BtcUsdtRateURL == "" {
		return fmt.Errorf("btc_usdt_rate_address field: empty")
	}

	if config.CurRubRateURL == "" {
		return fmt.Errorf("cur_rub_rate_address field: empty")
	}

	return nil
}

// NewDefaultConfig builds a Config with default values.
func NewDefaultConfig() Config {
	return Config{
		BtcUsdtRateURL: "https://api.kucoin.com/api/v1/market/stats?symbol=BTC-USDT",
		CurRubRateURL:  "http://www.cbr.ru/scripts/XML_daily.asp",
	}
}
