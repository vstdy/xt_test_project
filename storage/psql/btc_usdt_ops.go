package psql

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"

	canonical "github.com/vstdy/xt_test_project/model"
	"github.com/vstdy/xt_test_project/pkg"
	"github.com/vstdy/xt_test_project/storage/psql/schema"
)

const btcUsdtTableName = "btc_usdt"

// AddBtcUsdtRate adds BTC-USDT rate to storage.
func (st *Storage) AddBtcUsdtRate(ctx context.Context, obj canonical.BtcUsdt) error {
	dbObj := schema.NewBtcUsdtFromCanonical(obj)
	dbObj.ID = uuid.Nil

	_, err := st.DB.NewInsert().
		Model(&dbObj).
		Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

// BtcUsdtRateLatest returns latest BTC-USDT rate.
func (st *Storage) BtcUsdtRateLatest(ctx context.Context) (canonical.BtcUsdt, error) {
	var dbObj schema.BtcUsdt

	err := st.DB.NewSelect().
		Model(&dbObj).
		Order("timestamp DESC").
		Limit(1).
		Scan(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return canonical.BtcUsdt{}, pkg.ErrNoValue
		}

		return canonical.BtcUsdt{}, err
	}

	obj := dbObj.ToCanonical()

	return obj, nil
}

// BtcUsdtRateHistory returns BTC-USDT rate history.
func (st *Storage) BtcUsdtRateHistory(ctx context.Context) (int, []canonical.BtcUsdt, error) {
	var dbObjs schema.BtcUsdts

	count, err := st.DB.NewSelect().
		Model(&dbObjs).
		Order("timestamp DESC").
		Limit(10).
		Offset(0).
		ScanAndCount(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, nil, pkg.ErrNoValue
		}

		return 0, nil, err
	}

	obj := dbObjs.ToCanonical()

	return count, obj, nil
}
