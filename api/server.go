package api

import (
	"net/http"

	"github.com/vstdy/xt_test_project/service/exchange/v1"
)

// NewServer returns server.
func NewServer(svc *exchange.Service, config Config) *http.Server {
	router := NewRouter(svc, config)

	return &http.Server{Addr: config.RunAddress, Handler: router}
}
