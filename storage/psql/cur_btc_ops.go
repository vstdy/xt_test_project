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

const curBtcTableName = "cur_btc"

// AddCurBtcRates adds currencies to BTC to storage.
func (st *Storage) AddCurBtcRates(ctx context.Context, obj canonical.CurBtc) error {
	dbObj := schema.NewCurBtcFromCanonical(obj)
	dbObj.ID = uuid.Nil

	_, err := st.DB.NewInsert().
		Model(&dbObj).
		Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

// CurBtcRatesLatest gets latest currencies to BTC rates.
func (st *Storage) CurBtcRatesLatest(ctx context.Context) (canonical.CurBtc, error) {
	var dbObj schema.CurBtc

	err := st.DB.NewSelect().
		Model(&dbObj).
		Order("timestamp DESC").
		Limit(1).
		Scan(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return canonical.CurBtc{}, pkg.ErrNoValue
		}

		return canonical.CurBtc{}, err
	}

	obj := dbObj.ToCanonical()

	return obj, nil
}

// CurBtcRatesHistory returns currencies to BTC rates history.
func (st *Storage) CurBtcRatesHistory(ctx context.Context) (int, []canonical.CurBtc, error) {
	var dbObjs schema.CurBtcs

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
