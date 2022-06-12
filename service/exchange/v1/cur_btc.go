package exchange

import (
	"context"

	canonical "github.com/vstdy/xt_test_project/model"
)

// CurBtcRatesLatest gets latest currencies to RUB rate.
func (svc *Service) CurBtcRatesLatest(ctx context.Context) (canonical.CurBtc, error) {
	result, err := svc.storage.CurBtcRatesLatest(ctx)
	if err != nil {
		return canonical.CurBtc{}, err
	}

	return result, nil
}

// CurBtcRatesHistory gets currencies to RUB rates history.
func (svc *Service) CurBtcRatesHistory(ctx context.Context) (int, []canonical.CurBtc, error) {
	total, history, err := svc.storage.CurBtcRatesHistory(ctx)
	if err != nil {
		return 0, nil, err
	}

	return total, history, nil
}
