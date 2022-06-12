package cmd

import (
	"encoding/json"
	"net/http"

	"github.com/vstdy/xt_test_project/api/model"
	canonical "github.com/vstdy/xt_test_project/model"
	"github.com/vstdy/xt_test_project/provider/currency_rate/mock"
)

func (s *TestSuite) Test_btcUsdtLatest() {
	type request struct {
		method      string
		path        string
		body        string
		contentType string
	}

	type expected struct {
		code        int
		prepareBody func(obj canonical.BtcUsdt) string
		contentType string
	}

	type testCase struct {
		name         string
		prepareMocks func(curRatePrvMock *currencyRateMock.MockCurrencyRateProvider)
		request      request
		expected     expected
	}
	testCases := []testCase{
		{
			name: "OK",
			prepareMocks: func(curRatePrvMock *currencyRateMock.MockCurrencyRateProvider) {
				curRatePrvMock.EXPECT().
					BtcUsdtRate().
					Return(canonical.BtcUsdt{}, s.err)

				curRatePrvMock.EXPECT().
					CurRubRates().
					Return(canonical.CurRub{}, s.err)
			},
			request: request{
				method:      http.MethodGet,
				path:        "/api/btcusdt",
				body:        "",
				contentType: "",
			},
			expected: expected{
				code: http.StatusOK,
				prepareBody: func(obj canonical.BtcUsdt) string {
					respObj := model.NewBtcUsdtLatestResponseFromCanonical(obj)

					res, err := json.Marshal(respObj)
					s.Assert().NoError(err)

					return string(res)
				},
				contentType: "application/json",
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			tc.prepareMocks(s.prvMock)

			resp, body := s.testRequest(
				tc.request.method, tc.request.path, tc.request.body, tc.request.contentType)
			defer resp.Body.Close()

			s.Assert().Equal(tc.expected.code, resp.StatusCode)
			s.Assert().Equal(tc.expected.contentType, resp.Header.Get("Content-Type"))
			s.Assert().Equal(tc.expected.prepareBody(s.fixtures.BtcUsdts[1].ToCanonical()), body)
		})
	}
}

func (s *TestSuite) Test_btcUsdtHistory() {
	type request struct {
		method      string
		path        string
		body        string
		contentType string
	}

	type expected struct {
		code        int
		prepareBody func(count int, objs []canonical.BtcUsdt) string
		contentType string
	}

	type testCase struct {
		name         string
		prepareMocks func(curRatePrvMock *currencyRateMock.MockCurrencyRateProvider)
		request      request
		expected     expected
	}
	testCases := []testCase{
		{
			name: "OK",
			prepareMocks: func(curRatePrvMock *currencyRateMock.MockCurrencyRateProvider) {
				curRatePrvMock.EXPECT().
					BtcUsdtRate().
					Return(canonical.BtcUsdt{}, s.err)

				curRatePrvMock.EXPECT().
					CurRubRates().
					Return(canonical.CurRub{}, s.err)
			},
			request: request{
				method:      http.MethodPost,
				path:        "/api/btcusdt",
				body:        "",
				contentType: "application/json",
			},
			expected: expected{
				code: http.StatusOK,
				prepareBody: func(count int, objs []canonical.BtcUsdt) string {
					for i, j := 0, len(objs)-1; i < j; i, j = i+1, j-1 {
						objs[i], objs[j] = objs[j], objs[i]
					}
					batchRes := model.NewBtcUsdtHistoryResponseFromCanonical(count, objs)

					res, err := json.Marshal(batchRes)
					s.Assert().NoError(err)

					return string(res)
				},
				contentType: "application/json",
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			tc.prepareMocks(s.prvMock)

			resp, body := s.testRequest(
				tc.request.method, tc.request.path, tc.request.body, tc.request.contentType)
			defer resp.Body.Close()

			s.Assert().Equal(tc.expected.code, resp.StatusCode)
			s.Assert().Equal(tc.expected.contentType, resp.Header.Get("Content-Type"))
			s.Assert().Equal(tc.expected.prepareBody(
				len(s.fixtures.BtcUsdts), s.fixtures.BtcUsdts.ToCanonical()), body)
		})
	}
}

func (s *TestSuite) Test_curRubLatest() {
	type request struct {
		method      string
		path        string
		body        string
		contentType string
	}

	type expected struct {
		code        int
		prepareBody func(obj canonical.CurRub) string
		contentType string
	}

	type testCase struct {
		name         string
		prepareMocks func(curRatePrvMock *currencyRateMock.MockCurrencyRateProvider)
		request      request
		expected     expected
	}
	testCases := []testCase{
		{
			name: "OK",
			prepareMocks: func(curRatePrvMock *currencyRateMock.MockCurrencyRateProvider) {
				curRatePrvMock.EXPECT().
					BtcUsdtRate().
					Return(canonical.BtcUsdt{}, s.err)

				curRatePrvMock.EXPECT().
					CurRubRates().
					Return(canonical.CurRub{}, s.err)
			},
			request: request{
				method:      http.MethodGet,
				path:        "/api/currencies",
				body:        "",
				contentType: "",
			},
			expected: expected{
				code: http.StatusOK,
				prepareBody: func(obj canonical.CurRub) string {
					respObj := model.NewCurRubLatestResponseFromCanonical(obj)

					res, err := json.Marshal(respObj)
					s.Assert().NoError(err)

					return string(res)
				},
				contentType: "application/json",
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			tc.prepareMocks(s.prvMock)

			resp, body := s.testRequest(
				tc.request.method, tc.request.path, tc.request.body, tc.request.contentType)
			defer resp.Body.Close()

			s.Assert().Equal(tc.expected.code, resp.StatusCode)
			s.Assert().Equal(tc.expected.contentType, resp.Header.Get("Content-Type"))
			s.Assert().Equal(tc.expected.prepareBody(s.fixtures.CurRubs[1].ToCanonical()), body)
		})
	}
}

func (s *TestSuite) Test_curRubHistory() {
	type request struct {
		method      string
		path        string
		body        string
		contentType string
	}

	type expected struct {
		code        int
		prepareBody func(count int, objs []canonical.CurRub) string
		contentType string
	}

	type testCase struct {
		name         string
		prepareMocks func(curRatePrvMock *currencyRateMock.MockCurrencyRateProvider)
		request      request
		expected     expected
	}
	testCases := []testCase{
		{
			name: "OK",
			prepareMocks: func(curRatePrvMock *currencyRateMock.MockCurrencyRateProvider) {
				curRatePrvMock.EXPECT().
					BtcUsdtRate().
					Return(canonical.BtcUsdt{}, s.err)

				curRatePrvMock.EXPECT().
					CurRubRates().
					Return(canonical.CurRub{}, s.err)
			},
			request: request{
				method:      http.MethodPost,
				path:        "/api/currencies",
				body:        "",
				contentType: "application/json",
			},
			expected: expected{
				code: http.StatusOK,
				prepareBody: func(count int, objs []canonical.CurRub) string {
					for i, j := 0, len(objs)-1; i < j; i, j = i+1, j-1 {
						objs[i], objs[j] = objs[j], objs[i]
					}
					batchRes := model.NewCurRubHistoryResponseFromCanonical(count, objs)

					res, err := json.Marshal(batchRes)
					s.Assert().NoError(err)

					return string(res)
				},
				contentType: "application/json",
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			tc.prepareMocks(s.prvMock)

			resp, body := s.testRequest(
				tc.request.method, tc.request.path, tc.request.body, tc.request.contentType)
			defer resp.Body.Close()

			s.Assert().Equal(tc.expected.code, resp.StatusCode)
			s.Assert().Equal(tc.expected.contentType, resp.Header.Get("Content-Type"))
			s.Assert().Equal(tc.expected.prepareBody(
				len(s.fixtures.BtcUsdts), s.fixtures.CurRubs.ToCanonical()), body)
		})
	}
}

func (s *TestSuite) Test_curBtcLatest() {
	type request struct {
		method      string
		path        string
		body        string
		contentType string
	}

	type expected struct {
		code        int
		prepareBody func(obj canonical.CurBtc) string
		contentType string
	}

	type testCase struct {
		name         string
		prepareMocks func(curRatePrvMock *currencyRateMock.MockCurrencyRateProvider)
		request      request
		expected     expected
	}
	testCases := []testCase{
		{
			name: "OK",
			prepareMocks: func(curRatePrvMock *currencyRateMock.MockCurrencyRateProvider) {
				curRatePrvMock.EXPECT().
					BtcUsdtRate().
					Return(canonical.BtcUsdt{}, s.err)

				curRatePrvMock.EXPECT().
					CurRubRates().
					Return(canonical.CurRub{}, s.err)
			},
			request: request{
				method:      http.MethodGet,
				path:        "/api/latest",
				body:        "",
				contentType: "",
			},
			expected: expected{
				code: http.StatusOK,
				prepareBody: func(obj canonical.CurBtc) string {
					respObj := model.NewCurBtcLatestResponseFromCanonical(obj)

					res, err := json.Marshal(respObj)
					s.Assert().NoError(err)

					return string(res)
				},
				contentType: "application/json",
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			tc.prepareMocks(s.prvMock)

			resp, body := s.testRequest(
				tc.request.method, tc.request.path, tc.request.body, tc.request.contentType)
			defer resp.Body.Close()

			s.Assert().Equal(tc.expected.code, resp.StatusCode)
			s.Assert().Equal(tc.expected.contentType, resp.Header.Get("Content-Type"))
			s.Assert().Equal(tc.expected.prepareBody(s.fixtures.CurBtcs[1].ToCanonical()), body)
		})
	}
}

func (s *TestSuite) Test_curBtcHistory() {
	type request struct {
		method      string
		path        string
		body        string
		contentType string
	}

	type expected struct {
		code        int
		prepareBody func(count int, objs []canonical.CurBtc) string
		contentType string
	}

	type testCase struct {
		name         string
		prepareMocks func(curRatePrvMock *currencyRateMock.MockCurrencyRateProvider)
		request      request
		expected     expected
	}
	testCases := []testCase{
		{
			name: "OK",
			prepareMocks: func(curRatePrvMock *currencyRateMock.MockCurrencyRateProvider) {
				curRatePrvMock.EXPECT().
					BtcUsdtRate().
					Return(canonical.BtcUsdt{}, s.err)

				curRatePrvMock.EXPECT().
					CurRubRates().
					Return(canonical.CurRub{}, s.err)
			},
			request: request{
				method:      http.MethodPost,
				path:        "/api/latest",
				body:        "",
				contentType: "application/json",
			},
			expected: expected{
				code: http.StatusOK,
				prepareBody: func(count int, objs []canonical.CurBtc) string {
					for i, j := 0, len(objs)-1; i < j; i, j = i+1, j-1 {
						objs[i], objs[j] = objs[j], objs[i]
					}
					batchRes := model.NewCurBtcHistoryResponseFromCanonical(count, objs)

					res, err := json.Marshal(batchRes)
					s.Assert().NoError(err)

					return string(res)
				},
				contentType: "application/json",
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			tc.prepareMocks(s.prvMock)

			resp, body := s.testRequest(
				tc.request.method, tc.request.path, tc.request.body, tc.request.contentType)
			defer resp.Body.Close()

			s.Assert().Equal(tc.expected.code, resp.StatusCode)
			s.Assert().Equal(tc.expected.contentType, resp.Header.Get("Content-Type"))
			s.Assert().Equal(tc.expected.prepareBody(
				len(s.fixtures.BtcUsdts), s.fixtures.CurBtcs.ToCanonical()), body)
		})
	}
}
