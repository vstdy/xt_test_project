package exchange

import (
	"context"
	"io"

	canonical "github.com/vstdy/xt_test_project/model"
)

type Service interface {
	io.Closer

	// BtcUsdtRateLatest returns latest BTC-USDT rate.
	BtcUsdtRateLatest(ctx context.Context) (canonical.BtcUsdt, error)
	// BtcUsdtRateHistory returns BTC-USDT rate history.
	BtcUsdtRateHistory(
		ctx context.Context,
		pNum, pSize int,
		since, till string,
	) (int, []canonical.BtcUsdt, error)

	// CurRubRatesLatest returns latest currencies to RUB rate.
	CurRubRatesLatest(ctx context.Context) (canonical.CurRub, error)
	// CurRubRatesHistory returns currencies to RUB rates history.
	CurRubRatesHistory(
		ctx context.Context,
		pNum, pSize int,
		since, till, cur string,
	) (int, []canonical.CurRub, error)

	// CurBtcRatesLatest returns latest currencies to BTC rate.
	CurBtcRatesLatest(ctx context.Context) (canonical.CurBtc, error)
	// CurBtcRatesHistory returns currencies to BTC rates history.
	CurBtcRatesHistory(
		ctx context.Context,
		pNum, pSize int,
		since, till, cur string,
	) (int, []canonical.CurBtc, error)
}
