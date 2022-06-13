//go:generate mockgen -source=interface.go -destination=./mock/storage.go -package=storagemock
package storage

import (
	"context"
	"github.com/vstdy/xt_test_project/pkg/input"
	"io"

	canonical "github.com/vstdy/xt_test_project/model"
)

type Storage interface {
	io.Closer

	// AddBtcUsdtRate gets latest BTC-USDT rate.
	AddBtcUsdtRate(ctx context.Context, obj canonical.BtcUsdt) error
	// BtcUsdtRateLatest gets latest BTC-USDT rate.
	BtcUsdtRateLatest(ctx context.Context) (canonical.BtcUsdt, error)
	// BtcUsdtRateHistory gets BTC-USDT rate history.
	BtcUsdtRateHistory(
		ctx context.Context,
		pageParams input.PageParams,
		dateTimeParams input.DateTimeParams,
	) (int, []canonical.BtcUsdt, error)

	// AddCurRubRates gets latest currencies to RUB rates.
	AddCurRubRates(ctx context.Context, obj canonical.CurRub) error
	// CurRubRatesLatest gets latest currencies to RUB rates.
	CurRubRatesLatest(ctx context.Context) (canonical.CurRub, error)
	// CurRubRatesHistory gets currencies to RUB rates history.
	CurRubRatesHistory(
		ctx context.Context,
		pageParams input.PageParams,
		dateTimeParams input.DateTimeParams,
		cur string,
	) (int, []canonical.CurRub, error)

	// AddCurBtcRates gets latest currencies to BTC rates.
	AddCurBtcRates(ctx context.Context, obj canonical.CurBtc) error
	// CurBtcRatesLatest gets latest currencies to BTC rates.
	CurBtcRatesLatest(ctx context.Context) (canonical.CurBtc, error)
	// CurBtcRatesHistory gets currencies to BTC rates history.
	CurBtcRatesHistory(
		ctx context.Context,
		pageParams input.PageParams,
		dateTimeParams input.DateTimeParams,
		cur string,
	) (int, []canonical.CurBtc, error)
}
