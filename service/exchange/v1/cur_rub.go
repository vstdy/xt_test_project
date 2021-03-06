package exchange

import (
	"context"
	"fmt"
	"time"

	"github.com/rs/zerolog/log"

	canonical "github.com/vstdy/xt_test_project/model"
	"github.com/vstdy/xt_test_project/pkg/input"
)

// CurRubRatesLatest gets latest currencies to RUB rate.
func (svc *Service) CurRubRatesLatest(ctx context.Context) (canonical.CurRub, error) {
	result, err := svc.storage.CurRubRatesLatest(ctx)
	if err != nil {
		return canonical.CurRub{}, err
	}

	return result, nil
}

// CurRubRatesHistory gets currencies to RUB rates history.
func (svc *Service) CurRubRatesHistory(
	ctx context.Context,
	pNum, pSize int,
	since, till, cur string,
) (int, []canonical.CurRub, error) {
	pp, err := input.NewPageParams(pNum, pSize)
	if err != nil {
		return 0, nil, err
	}
	dt, err := input.NewDateTimeParams(since, till, input.Date)
	if err != nil {
		return 0, nil, err
	}
	if cur != "" {
		if err = canonical.ValidateCurRub(cur); err != nil {
			return 0, nil, err
		}
	}

	total, history, err := svc.storage.CurRubRatesHistory(ctx, pp, dt, cur)
	if err != nil {
		return 0, nil, err
	}

	return total, history, nil
}

// curRubRateUpdater updates currencies to RUB rate.
func (svc *Service) curRubRateUpdater(ctx context.Context) {
	update := func() error {
		currencyRate, err := svc.currencyRateProvider.CurRubRates()
		if err != nil {
			return fmt.Errorf("currencyRateProvider: %w", err)
		}

		updCtx, cancel := context.WithTimeout(context.Background(), svc.config.UpdaterTimeout)
		defer cancel()

		if err = svc.storage.AddCurRubRates(updCtx, currencyRate); err != nil {
			return fmt.Errorf("add accruals: %w", err)
		}

		return nil
	}

	if err := update(); err != nil {
		log.Warn().Err(err).Msg("curRubRateUpdater:")
	}

	ticker := time.NewTicker(svc.config.CurRubRateCheckInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			log.Info().Msg("curRubRateUpdater closed")
			return
		case <-ticker.C:
			if err := update(); err != nil {
				log.Warn().Err(err).Msg("curRubRateUpdater:")
			}
		}
	}
}
