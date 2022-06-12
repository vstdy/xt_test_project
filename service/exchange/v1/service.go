package exchange

import (
	"context"
	"fmt"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/vstdy/xt_test_project/pkg/logging"
	"github.com/vstdy/xt_test_project/provider/currency_rate"
	"github.com/vstdy/xt_test_project/service/exchange"
	"github.com/vstdy/xt_test_project/storage"
)

const (
	serviceName = "exchange"
)

var _ exchange.Service = (*Service)(nil)

type (
	// Service keeps service dependencies.
	Service struct {
		config               Config
		currencyRateProvider currencyRate.CurrencyRateProvider
		storage              storage.Storage
	}

	// ServiceOption defines functional argument for Service constructor.
	ServiceOption func(*Service) error
)

// WithConfig sets Config.
func WithConfig(config Config) ServiceOption {
	return func(svc *Service) error {
		svc.config = config

		return nil
	}
}

// WithCurrencyRateProvider sets currency rate provider.
func WithCurrencyRateProvider(crp currencyRate.CurrencyRateProvider) ServiceOption {
	return func(svc *Service) error {
		svc.currencyRateProvider = crp

		return nil
	}
}

// WithStorage sets Storage.
func WithStorage(st storage.Storage) ServiceOption {
	return func(svc *Service) error {
		svc.storage = st

		return nil
	}
}

// NewService creates a new exchange service.
func NewService(ctx context.Context, opts ...ServiceOption) (*Service, error) {
	svc := &Service{
		config: NewDefaultConfig(),
	}
	for optIdx, opt := range opts {
		if err := opt(svc); err != nil {
			return nil, fmt.Errorf("applying option [%d]: %w", optIdx, err)
		}
	}

	if err := svc.config.Validate(); err != nil {
		return nil, fmt.Errorf("config validation: %w", err)
	}

	if svc.storage == nil {
		return nil, fmt.Errorf("storage: nil")
	}

	go time.AfterFunc(time.Millisecond, func() {
		go svc.btcUsdtRateUpdater(ctx)
		go svc.curRubRateUpdater(ctx)
	})

	return svc, nil
}

// Close closes all service dependencies.
func (svc *Service) Close() error {
	if svc.storage == nil {
		return nil
	}
	if err := svc.storage.Close(); err != nil {
		return fmt.Errorf("closing storage: %w", err)
	}

	return nil
}

// Logger returns logger with service context.
func (svc *Service) Logger() zerolog.Logger {
	logCtx := log.With().Str(logging.ServiceKey, serviceName)

	return logCtx.Logger()
}
