//go:generate mockgen -source=interface.go -destination=./mock/client.go -package=currencyRateMock
package currencyRate

import (
	canonical "github.com/vstdy/xt_test_project/model"
)

type CurrencyRateProvider interface {
	// BtcUsdtRate gets BTC-USDT rate.
	BtcUsdtRate() (canonical.BtcUsdt, error)
	// CurRubRates gets currencies to RUB rates.
	CurRubRates() (canonical.CurRub, error)
}
