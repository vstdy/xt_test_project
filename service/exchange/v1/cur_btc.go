package exchange

import (
	"context"

	canonical "github.com/vstdy/xt_test_project/model"
	"github.com/vstdy/xt_test_project/pkg/input"
)

// CurBtcRatesLatest gets the latest currencies to BTC rate.
func (svc *Service) CurBtcRatesLatest(ctx context.Context) (canonical.CurBtc, error) {
	result, err := svc.storage.CurBtcRatesLatest(ctx)
	if err != nil {
		return canonical.CurBtc{}, err
	}

	return result, nil
}

// CurBtcRatesHistory gets currencies to BTC rates history.
func (svc *Service) CurBtcRatesHistory(
	ctx context.Context,
	pNum, pSize int,
	since, till, cur string,
) (int, []canonical.CurBtc, error) {
	pp, err := input.NewPageParams(pNum, pSize)
	if err != nil {
		return 0, nil, err
	}
	dt, err := input.NewDateTimeParams(since, till, input.Datetime)
	if err != nil {
		return 0, nil, err
	}
	if cur != "" {
		if err = canonical.ValidateCurBtc(cur); err != nil {
			return 0, nil, err
		}
	}

	total, history, err := svc.storage.CurBtcRatesHistory(ctx, pp, dt, cur)
	if err != nil {
		return 0, nil, err
	}

	return total, history, nil
}
