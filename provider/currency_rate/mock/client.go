// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package currencyRateMock is a generated GoMock package.
package currencyRateMock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/vstdy/xt_test_project/model"
)

// MockCurrencyRateProvider is a mock of CurrencyRateProvider interface.
type MockCurrencyRateProvider struct {
	ctrl     *gomock.Controller
	recorder *MockCurrencyRateProviderMockRecorder
}

// MockCurrencyRateProviderMockRecorder is the mock recorder for MockCurrencyRateProvider.
type MockCurrencyRateProviderMockRecorder struct {
	mock *MockCurrencyRateProvider
}

// NewMockCurrencyRateProvider creates a new mock instance.
func NewMockCurrencyRateProvider(ctrl *gomock.Controller) *MockCurrencyRateProvider {
	mock := &MockCurrencyRateProvider{ctrl: ctrl}
	mock.recorder = &MockCurrencyRateProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCurrencyRateProvider) EXPECT() *MockCurrencyRateProviderMockRecorder {
	return m.recorder
}

// BtcUsdtRate mocks base method.
func (m *MockCurrencyRateProvider) BtcUsdtRate() (model.BtcUsdt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BtcUsdtRate")
	ret0, _ := ret[0].(model.BtcUsdt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BtcUsdtRate indicates an expected call of BtcUsdtRate.
func (mr *MockCurrencyRateProviderMockRecorder) BtcUsdtRate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BtcUsdtRate", reflect.TypeOf((*MockCurrencyRateProvider)(nil).BtcUsdtRate))
}

// CurRubRates mocks base method.
func (m *MockCurrencyRateProvider) CurRubRates() (model.CurRub, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CurRubRates")
	ret0, _ := ret[0].(model.CurRub)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CurRubRates indicates an expected call of CurRubRates.
func (mr *MockCurrencyRateProviderMockRecorder) CurRubRates() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CurRubRates", reflect.TypeOf((*MockCurrencyRateProvider)(nil).CurRubRates))
}
