package currencyRate

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"time"

	"golang.org/x/text/encoding/charmap"

	canonical "github.com/vstdy/xt_test_project/model"
	"github.com/vstdy/xt_test_project/provider/currency_rate"
	"github.com/vstdy/xt_test_project/provider/currency_rate/http/model"
)

var _ currencyRate.CurrencyRateProvider = (*CurrencyRate)(nil)

// CurrencyRate keeps currency rate provider configuration.
type (
	CurrencyRate struct {
		config Config
		client http.Client
	}

	// CurrencyRateOption defines functional argument for CurrencyRate constructor.
	CurrencyRateOption func(*CurrencyRate) error
)

// WithConfig sets Config.
func WithConfig(config Config) CurrencyRateOption {
	return func(svc *CurrencyRate) error {
		svc.config = config

		return nil
	}
}

// NewCurrencyRateProvider returns a new CurrencyRate instance.
func NewCurrencyRateProvider(timeout time.Duration, opts ...CurrencyRateOption) (*CurrencyRate, error) {
	acr := &CurrencyRate{
		config: NewDefaultConfig(),
	}
	for optIdx, opt := range opts {
		if err := opt(acr); err != nil {
			return nil, fmt.Errorf("applying option [%d]: %w", optIdx, err)
		}
	}

	if err := acr.config.Validate(); err != nil {
		return nil, fmt.Errorf("config validation: %w", err)
	}

	acr.client = http.Client{Timeout: timeout}
	transport := &http.Transport{}
	transport.MaxIdleConns = 1
	transport.IdleConnTimeout = 0
	acr.client.Transport = transport

	return acr, nil
}

// BtcUsdtRate gets BTC-USDT rate.
func (cr *CurrencyRate) BtcUsdtRate() (canonical.BtcUsdt, error) {
	r, err := cr.client.Get(cr.config.BtcUsdtRateURL)
	if err != nil {
		return canonical.BtcUsdt{}, fmt.Errorf("retrieving BtcUsdt object: %w", err)
	}
	defer r.Body.Close()

	if r.StatusCode != http.StatusOK {
		return canonical.BtcUsdt{}, nil
	}

	var btcUsdtRate model.BtcUsdtRate

	if err = json.NewDecoder(r.Body).Decode(&btcUsdtRate); err != nil {
		return canonical.BtcUsdt{}, fmt.Errorf("decoding BtcUsdtRate object: %w", err)
	}

	obj, err := btcUsdtRate.ToCanonical()
	if err != nil {
		return canonical.BtcUsdt{}, fmt.Errorf("converting BtcUsdtRate object to canonical: %w", err)
	}

	return obj, nil
}

// CurRubRates gets various currencies to RUB rates.
func (cr *CurrencyRate) CurRubRates() (canonical.CurRub, error) {
	r, err := cr.client.Get(cr.config.CurRubRateURL)
	if err != nil {
		return canonical.CurRub{}, fmt.Errorf("retrieving CurRub object: %w", err)
	}
	defer r.Body.Close()

	if r.StatusCode != http.StatusOK {
		return canonical.CurRub{}, nil
	}

	var currenciesToRubRates model.CurRubRates

	decoder := xml.NewDecoder(r.Body)
	decoder.CharsetReader = makeCharsetReader
	if err = decoder.Decode(&currenciesToRubRates); err != nil {
		return canonical.CurRub{}, fmt.Errorf("decoding CurRubRates object: %w", err)
	}

	obj, err := currenciesToRubRates.ToCanonical()
	if err != nil {
		return canonical.CurRub{}, fmt.Errorf("converting CurRubRates object to canonical: %w", err)
	}

	return obj, nil
}

func makeCharsetReader(charset string, input io.Reader) (io.Reader, error) {
	if charset == "windows-1251" {
		return charmap.Windows1251.NewDecoder().Reader(input), nil
	}

	return nil, fmt.Errorf("unknown charset: %s", charset)
}
