package psql

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"

	canonical "github.com/vstdy/xt_test_project/model"
	"github.com/vstdy/xt_test_project/pkg"
	"github.com/vstdy/xt_test_project/pkg/input"
	"github.com/vstdy/xt_test_project/storage/psql/schema"
)

const curRubTableName = "cur_rub"

// AddCurRubRates adds currencies to RUB to storage.
func (st *Storage) AddCurRubRates(ctx context.Context, obj canonical.CurRub) error {
	dbObj := schema.NewCurRubFromCanonical(obj)
	dbObj.ID = uuid.Nil

	_, err := st.DB.NewInsert().
		Model(&dbObj).
		On("CONFLICT (\"date\") DO NOTHING").
		Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

// CurRubRatesLatest gets latest currencies to RUB rates.
func (st *Storage) CurRubRatesLatest(ctx context.Context) (canonical.CurRub, error) {
	var dbObj schema.CurRub

	err := st.DB.NewSelect().
		Model(&dbObj).
		Order("date DESC").
		Limit(1).
		Scan(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return canonical.CurRub{}, pkg.ErrNoValue
		}

		return canonical.CurRub{}, err
	}

	obj := dbObj.ToCanonical()

	return obj, nil
}

// CurRubRatesHistory returns currencies to RUB rates history.
func (st *Storage) CurRubRatesHistory(
	ctx context.Context,
	pageParams input.PageParams,
	dateTimeParams input.DateTimeParams,
	cur string,
) (int, []canonical.CurRub, error) {
	var dbObjs schema.CurRubs

	q := st.DB.NewSelect().
		Model(&dbObjs).
		Order("date DESC")

	columnFilterQuery(q, cur, "date")
	dateTimeFilterQuery(q, "date", dateTimeParams)
	paginateQuery(q, pageParams)

	count, err := q.ScanAndCount(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, nil, pkg.ErrNoValue
		}

		return 0, nil, err
	}

	obj := dbObjs.ToCanonical()

	return count, obj, nil
}
