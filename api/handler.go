package api

import (
	"encoding/json"
	"errors"
	"net/http"

	_ "github.com/jackc/pgx/v4/stdlib"

	canonical "github.com/vstdy/xt_test_project/api/model"
	"github.com/vstdy/xt_test_project/pkg"
	"github.com/vstdy/xt_test_project/service/exchange"
)

// Handler keeps handler dependencies.
type Handler struct {
	service exchange.Service
}

// NewHandler returns a new Handler instance.
func NewHandler(service exchange.Service) Handler {
	return Handler{service: service}
}

// btcUsdtLatest returns latest BTC-USDT rate.
func (h Handler) btcUsdtLatest(w http.ResponseWriter, r *http.Request) {
	obj, err := h.service.BtcUsdtRateLatest(r.Context())
	if err != nil {
		if errors.Is(err, pkg.ErrNoValue) {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respObj := canonical.NewBtcUsdtLatestResponseFromCanonical(obj)

	resp, err := json.Marshal(respObj)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err = w.Write(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// btcUsdtHistory returns BTC-USDT rate history.
func (h Handler) btcUsdtHistory(w http.ResponseWriter, r *http.Request) {
	count, objs, err := h.service.BtcUsdtRateHistory(r.Context())
	if err != nil {
		if errors.Is(err, pkg.ErrNoValue) {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respObjs := canonical.NewBtcUsdtHistoryResponseFromCanonical(count, objs)

	resp, err := json.Marshal(respObjs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err = w.Write(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// curRubLatest returns latest currencies to RUB rates.
func (h Handler) curRubLatest(w http.ResponseWriter, r *http.Request) {
	obj, err := h.service.CurRubRatesLatest(r.Context())
	if err != nil {
		if errors.Is(err, pkg.ErrNoValue) {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respObj := canonical.NewCurRubLatestResponseFromCanonical(obj)

	resp, err := json.Marshal(respObj)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err = w.Write(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// curRubHistory returns currencies to RUB rates history.
func (h Handler) curRubHistory(w http.ResponseWriter, r *http.Request) {
	count, objs, err := h.service.CurRubRatesHistory(r.Context())
	if err != nil {
		if errors.Is(err, pkg.ErrNoValue) {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respObjs := canonical.NewCurRubHistoryResponseFromCanonical(count, objs)

	resp, err := json.Marshal(respObjs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err = w.Write(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// curBtcLatest returns latest currencies to BTC rates.
func (h Handler) curBtcLatest(w http.ResponseWriter, r *http.Request) {
	obj, err := h.service.CurBtcRatesLatest(r.Context())
	if err != nil {
		if errors.Is(err, pkg.ErrNoValue) {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respObj := canonical.NewCurBtcLatestResponseFromCanonical(obj)

	resp, err := json.Marshal(respObj)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err = w.Write(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// curBtcHistory returns currencies to RUB rates history.
func (h Handler) curBtcHistory(w http.ResponseWriter, r *http.Request) {
	count, objs, err := h.service.CurBtcRatesHistory(r.Context())
	if err != nil {
		if errors.Is(err, pkg.ErrNoValue) {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respObjs := canonical.NewCurBtcHistoryResponseFromCanonical(count, objs)

	resp, err := json.Marshal(respObjs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err = w.Write(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
