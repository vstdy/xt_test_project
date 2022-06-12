package common

import (
	"context"
	"fmt"
	"time"

	"github.com/vstdy/xt_test_project/api"
	"github.com/vstdy/xt_test_project/pkg"
	"github.com/vstdy/xt_test_project/provider/currency_rate/http"
	"github.com/vstdy/xt_test_project/service/exchange/v1"
	"github.com/vstdy/xt_test_project/storage"
	"github.com/vstdy/xt_test_project/storage/psql"
)

// Config combines sub-configs for all services, storages and providers.
type Config struct {
	Timeout      time.Duration       `mapstructure:"timeout"`
	StorageType  string              `mapstructure:"storage_type"`
	HTTPServer   api.Config          `mapstructure:"server,squash"`
	CurrencyRate currencyRate.Config `mapstructure:"server,squash"`
	Service      exchange.Config     `mapstructure:"service,squash"`
	PSQLStorage  psql.Config         `mapstructure:"psql_storage,squash"`
}

const (
	psqlStorage = "psql"
)

// BuildDefaultConfig builds a Config with default values.
func BuildDefaultConfig() Config {
	return Config{
		Timeout:      5 * time.Second,
		StorageType:  psqlStorage,
		HTTPServer:   api.NewDefaultConfig(),
		CurrencyRate: currencyRate.NewDefaultConfig(),
		Service:      exchange.NewDefaultConfig(),
		PSQLStorage:  psql.NewDefaultConfig(),
	}
}

// BuildPsqlStorage builds psql.Storage dependency.
func (config Config) BuildPsqlStorage() (*psql.Storage, error) {
	st, err := psql.NewStorage(
		psql.WithConfig(config.PSQLStorage),
	)
	if err != nil {
		return nil, fmt.Errorf("building psql storage: %w", err)
	}

	return st, nil
}

// BuildService builds exchange.Service dependency.
func (config Config) BuildService(ctx context.Context) (*exchange.Service, error) {
	var st storage.Storage
	var err error

	crp, err := currencyRate.NewCurrencyRateProvider(
		config.Timeout,
		currencyRate.WithConfig(config.CurrencyRate),
	)
	if err != nil {
		return nil, fmt.Errorf("building currency rate provider: %w", err)
	}

	switch config.StorageType {
	case psqlStorage:
		st, err = config.BuildPsqlStorage()
	default:
		err = pkg.ErrUnsupportedStorageType
	}
	if err != nil {
		return nil, fmt.Errorf("building storage: %w", err)
	}

	svc, err := exchange.NewService(
		ctx,
		exchange.WithConfig(config.Service),
		exchange.WithCurrencyRateProvider(crp),
		exchange.WithStorage(st),
	)
	if err != nil {
		return nil, fmt.Errorf("building service: %w", err)
	}

	return svc, nil
}
