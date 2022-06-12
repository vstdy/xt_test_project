package cmd

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/suite"

	"github.com/vstdy/xt_test_project/api"
	"github.com/vstdy/xt_test_project/cmd/exchange/cmd/common"
	"github.com/vstdy/xt_test_project/provider/currency_rate/mock"
	"github.com/vstdy/xt_test_project/service/exchange/v1"
	"github.com/vstdy/xt_test_project/storage/psql"
	"github.com/vstdy/xt_test_project/storage/psql/fixtures"
	"github.com/vstdy/xt_test_project/testutils"
)

type TestSuite struct {
	suite.Suite

	srv     *httptest.Server
	prvMock *currencyRateMock.MockCurrencyRateProvider

	config common.Config
	client *http.Client

	container *testutils.PostgreSQLContainer
	storage   *psql.Storage
	fixtures  fixtures.Fixtures

	ctx context.Context
	err error
}

func (s *TestSuite) SetupSuite() {
	logWriter := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
	}
	log.Logger = log.Output(logWriter).Level(1)

	ctx, ctxCancel := context.WithTimeout(context.Background(), time.Minute)
	defer ctxCancel()

	c, err := testutils.NewPostgreSQLContainer(ctx)
	s.Require().NoError(err)

	config := common.BuildDefaultConfig()
	config.PSQLStorage.URI = c.GetDSN()

	st, err := config.BuildPsqlStorage()
	s.Require().NoError(err)

	s.Require().NoError(st.Migrate(ctx))

	fixts, err := fixtures.LoadFixtures(ctx, st.DB)
	s.Require().NoError(err)

	mockCtrl := gomock.NewController(s.T())
	crp := currencyRateMock.NewMockCurrencyRateProvider(mockCtrl)

	svc, err := exchange.NewService(
		ctx,
		exchange.WithConfig(config.Service),
		exchange.WithCurrencyRateProvider(crp),
		exchange.WithStorage(st),
	)
	timeout := time.Second

	clt := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Timeout: timeout,
	}

	r := api.NewRouter(svc, config.HTTPServer)
	srv := httptest.NewServer(r)

	s.srv = srv
	s.prvMock = crp
	s.config = config
	s.client = clt
	s.container = c
	s.storage = st
	s.fixtures = fixts
	s.ctx = ctx
	s.err = errors.New("mock error")
}

func (s *TestSuite) TearDownSuite() {
	ctx, ctxCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer ctxCancel()

	s.Require().NoError(s.storage.Close())
	s.Require().NoError(s.container.Terminate(ctx))
	s.srv.Close()
}

func (s TestSuite) testRequest(method, path, body, contentType string) (*http.Response, string) {
	req, err := http.NewRequest(method, s.srv.URL+path, strings.NewReader(body))
	s.Require().NoError(err)
	req.Header.Set("Content-Type", contentType)

	resp, err := s.client.Do(req)
	s.Require().NoError(err)

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	s.Require().NoError(err)

	defer resp.Body.Close()

	return resp, string(respBody)
}

func TestSuite_Service(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
