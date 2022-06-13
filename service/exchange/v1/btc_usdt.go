package exchange

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/rs/zerolog/log"

	canonical "github.com/vstdy/xt_test_project/model"
	"github.com/vstdy/xt_test_project/pkg"
	"github.com/vstdy/xt_test_project/pkg/input"
)

// BtcUsdtRateLatest gets latest BTC-USDT rate.
func (svc *Service) BtcUsdtRateLatest(ctx context.Context) (canonical.BtcUsdt, error) {
	result, err := svc.storage.BtcUsdtRateLatest(ctx)
	if err != nil {
		return canonical.BtcUsdt{}, err
	}

	return result, nil
}

// BtcUsdtRateHistory gets BTC-USDT rate history.
func (svc *Service) BtcUsdtRateHistory(
	ctx context.Context,
	pNum, pSize int,
	since, till string,
) (int, []canonical.BtcUsdt, error) {
	pp, err := input.NewPageParams(pNum, pSize)
	if err != nil {
		return 0, nil, err
	}
	dt, err := input.NewDateTimeParams(since, till, input.Datetime)
	if err != nil {
		return 0, nil, err
	}

	total, history, err := svc.storage.BtcUsdtRateHistory(ctx, pp, dt)
	if err != nil {
		return 0, nil, err
	}

	return total, history, nil
}

// btcUsdtRateUpdater updates BTC-USDT rate.
func (svc *Service) btcUsdtRateUpdater(ctx context.Context) {
	btcUsdtRateLatest := func() (canonical.BtcUsdt, error) {
		updCtx, cancel := context.WithTimeout(context.Background(), svc.config.UpdaterTimeout)
		defer cancel()

		return svc.storage.BtcUsdtRateLatest(updCtx)
	}

	addBtcUsdtRate := func(btcUstdRate canonical.BtcUsdt) error {
		updCtx, cancel := context.WithTimeout(context.Background(), svc.config.UpdaterTimeout)
		defer cancel()

		return svc.storage.AddBtcUsdtRate(updCtx, btcUstdRate)
	}

	curRubRatesLatest := func() (canonical.CurRub, error) {
		updCtx, cancel := context.WithTimeout(context.Background(), svc.config.UpdaterTimeout)
		defer cancel()

		return svc.storage.CurRubRatesLatest(updCtx)
	}

	addCurBtcRates := func(curBtcRates canonical.CurBtc) error {
		updCtx, cancel := context.WithTimeout(context.Background(), svc.config.UpdaterTimeout)
		defer cancel()

		return svc.storage.AddCurBtcRates(updCtx, curBtcRates)
	}

	update := func() error {
		btcUsdtRate, err := svc.currencyRateProvider.BtcUsdtRate()
		if err != nil {
			return fmt.Errorf("currencyRateProvider: %w", err)
		}

		btcUsdtRateSaved, err := btcUsdtRateLatest()
		if err != nil && !errors.Is(err, pkg.ErrNoValue) {
			return fmt.Errorf("get latest BTC-USDT rate: %w", err)
		}

		if btcUsdtRate.Buy != btcUsdtRateSaved.Buy ||
			btcUsdtRate.Sell != btcUsdtRateSaved.Sell ||
			errors.Is(err, pkg.ErrNoValue) {
			if err = addBtcUsdtRate(btcUsdtRate); err != nil {
				return fmt.Errorf("add BTC-USDT rate: %w", err)
			}

			curRubRates, err := curRubRatesLatest()
			if err != nil {
				return fmt.Errorf("get latest currencies to RUB rates: %w", err)
			}

			curBtc := canonical.NewCurBtc(curRubRates, btcUsdtRate)

			if err = addCurBtcRates(curBtc); err != nil {
				return fmt.Errorf("add currencies to BTC rates: %w", err)
			}
		}

		return nil
	}

	if err := update(); err != nil {
		log.Warn().Err(err).Msg("btcUsdtRateUpdater:")
	}

	ticker := time.NewTicker(svc.config.BtcUsdtRateCheckInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			log.Info().Msg("btcUsdtRateUpdater closed")
			return
		case <-ticker.C:
			if err := update(); err != nil {
				log.Warn().Err(err).Msg("btcUsdtRateUpdater:")
			}
		}
	}
}
