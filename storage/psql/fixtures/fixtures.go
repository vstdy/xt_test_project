package fixtures

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"text/template"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dbfixture"

	"github.com/vstdy/xt_test_project/storage/psql/schema"
)

// Fixtures keeps all fixture objects.
type Fixtures struct {
	BtcUsdts schema.BtcUsdts
	CurRubs  schema.CurRubs
	CurBtcs  schema.CurBtcs
}

func (f *Fixtures) appendBtcUsdt(obj interface{}) error {
	dbObj, ok := obj.(*schema.BtcUsdt)
	if !ok {
		return fmt.Errorf("BtcUsdt: type assert failed: %T", obj)
	}
	if dbObj == nil {
		return fmt.Errorf("BtcUsdt: type assert failed: nil")
	}
	f.BtcUsdts = append(f.BtcUsdts, *dbObj)

	return nil
}

func (f *Fixtures) appendCurRub(obj interface{}) error {
	dbObj, ok := obj.(*schema.CurRub)
	if !ok {
		return fmt.Errorf("CurRub: type assert failed: %T", obj)
	}
	if dbObj == nil {
		return fmt.Errorf("CurRub: type assert failed: nil")
	}
	f.CurRubs = append(f.CurRubs, *dbObj)

	return nil
}

func (f *Fixtures) appendCurBtc(obj interface{}) error {
	dbObj, ok := obj.(*schema.CurBtc)
	if !ok {
		return fmt.Errorf("CurBtc: type assert failed: %T", obj)
	}
	if dbObj == nil {
		return fmt.Errorf("CurBtc: type assert failed: nil")
	}
	f.CurBtcs = append(f.CurBtcs, *dbObj)

	return nil
}

// LoadFixtures load fixtures to DB and returns DB objects aggregate.
func LoadFixtures(ctx context.Context, db *bun.DB) (Fixtures, error) {
	type fixturesAppender struct {
		id     string
		append func(obj interface{}) error
	}

	fixtureManager := dbfixture.New(
		db,
		dbfixture.WithTemplateFuncs(template.FuncMap{
			"yesterday": func() string {
				return time.Now().Add(-24 * time.Hour).UTC().Format(time.RFC3339Nano)
			},
			"now": func() string {
				return time.Now().UTC().Format(time.RFC3339Nano)
			},
			"uuid": func() uuid.UUID {
				return uuid.New()
			},
		}),
	)

	err := fixtureManager.Load(
		ctx,
		os.DirFS(getFixturesDir()),
		"btc_usdt.yaml", "cur_rub.yaml", "cur_btc.yaml",
	)
	if err != nil {
		return Fixtures{}, fmt.Errorf("loading fixtures: %w", err)
	}

	fixtures := Fixtures{}
	appenders := []fixturesAppender{
		{id: "BtcUsdt.btcUsdt_1", append: fixtures.appendBtcUsdt},
		{id: "BtcUsdt.btcUsdt_2", append: fixtures.appendBtcUsdt},
		{id: "CurRub.curRub_1", append: fixtures.appendCurRub},
		{id: "CurRub.curRub_2", append: fixtures.appendCurRub},
		{id: "CurBtc.curBtc_1", append: fixtures.appendCurBtc},
		{id: "CurBtc.curBtc_2", append: fixtures.appendCurBtc},
	}
	for _, appender := range appenders {
		obj, err := fixtureManager.Row(appender.id)
		if err != nil {
			return Fixtures{}, fmt.Errorf("reading fixtures row (%s): %w", appender.id, err)
		}
		if obj == nil {
			return Fixtures{}, fmt.Errorf("reading fixtures row (%s): nil", appender.id)
		}
		if err := appender.append(obj); err != nil {
			return Fixtures{}, fmt.Errorf("appending fixtures row (%s): %w", appender.id, err)
		}
	}

	return fixtures, nil
}

// getFixturesDir returns current file directory.
func getFixturesDir() string {
	_, filePath, _, ok := runtime.Caller(1)
	if !ok {
		return ""
	}

	return filepath.Dir(filePath)
}
