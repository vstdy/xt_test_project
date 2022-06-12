package psql

import (
	"context"
	"database/sql"
	"fmt"
	"runtime"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/migrate"

	inter "github.com/vstdy/xt_test_project/storage"
	"github.com/vstdy/xt_test_project/storage/psql/migrations"
	"github.com/vstdy/xt_test_project/storage/psql/schema"
)

const (
	serviceName = "psql"

	dbTableLoggingKey     = "db-table"
	dbOperationLoggingKey = "db-operation"
)

var _ inter.Storage = (*Storage)(nil)

type (
	// Storage keeps psql storage dependencies.
	Storage struct {
		config Config
		DB     *bun.DB
	}

	// StorageOption defines functional argument for Storage constructor.
	StorageOption func(st *Storage) error
)

// WithConfig overrides default Storage config.
func WithConfig(config Config) StorageOption {
	return func(st *Storage) error {
		st.config = config

		return nil
	}
}

// NewStorage creates a new psql Storage with custom options.
func NewStorage(opts ...StorageOption) (*Storage, error) {
	st := &Storage{
		config: NewDefaultConfig(),
	}
	for optIdx, opt := range opts {
		if err := opt(st); err != nil {
			return nil, fmt.Errorf("applying option [%d]: %w", optIdx, err)
		}
	}

	if err := st.config.Validate(); err != nil {
		return nil, fmt.Errorf("config validation: %w", err)
	}

	sqlDB := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(st.config.URI)))

	maxOpenConnections := 4 * runtime.GOMAXPROCS(0)

	st.DB = bun.NewDB(sqlDB, pgdialect.New())
	st.DB.AddQueryHook(newQueryHook(st))
	st.DB.SetMaxOpenConns(maxOpenConnections)
	st.DB.SetMaxIdleConns(maxOpenConnections)
	st.DB.RegisterModel(
		(*schema.BtcUsdt)(nil),
		(*schema.CurRub)(nil),
		(*schema.CurBtc)(nil),
	)

	if err := st.DB.Ping(); err != nil {
		return nil, fmt.Errorf("ping for URI (%s) failed: %w", st.config.URI, err)
	}

	return st, nil
}

// Close closes DB connection.
func (st *Storage) Close() error {
	if st.DB == nil {
		return nil
	}

	return st.DB.Close()
}

// Migrate performs DB migrations.
func (st *Storage) Migrate(ctx context.Context) error {
	logger := st.Logger(withOperation("migration"))

	ms, err := migrations.GetMigrations()
	if err != nil {
		return err
	}

	migration := migrate.NewMigrator(st.DB, ms)

	if err = migration.Init(ctx); err != nil {
		return fmt.Errorf("initialising migration: %w", err)
	}

	res, err := migration.Migrate(ctx)
	if err != nil {
		return fmt.Errorf("performing migration: %w", err)
	}

	logger.Info().Msgf("Migration applied: %s", res.Migrations.LastGroup().String())

	return nil
}
