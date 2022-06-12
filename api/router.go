package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/vstdy/xt_test_project/service/exchange"
)

// NewRouter returns router.
func NewRouter(svc exchange.Service, config Config) chi.Router {
	h := NewHandler(svc)
	r := chi.NewRouter()

	r.Use(
		middleware.RequestID,
		middleware.RealIP,
		middleware.Logger,
		middleware.Recoverer,
		middleware.StripSlashes,
		middleware.Timeout(config.Timeout),
	)

	r.Route("/api", func(r chi.Router) {
		r.Route("/btcusdt", func(r chi.Router) {
			r.Get("/", h.btcUsdtLatest)
			r.Post("/", h.btcUsdtHistory)
		})

		r.Route("/currencies", func(r chi.Router) {
			r.Get("/", h.curRubLatest)
			r.Post("/", h.curRubHistory)
		})

		r.Route("/latest", func(r chi.Router) {
			r.Get("/", h.curBtcLatest)
			r.Post("/", h.curBtcHistory)
		})
	})

	return r
}
