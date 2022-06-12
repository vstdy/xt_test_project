package psql

import (
	"context"
	"time"

	"github.com/uptrace/bun"

	"github.com/vstdy/xt_test_project/pkg/logging"
)

type queryHook struct {
	st *Storage
}

func newQueryHook(storage *Storage) queryHook {
	return queryHook{
		st: storage,
	}
}

func (h queryHook) BeforeQuery(ctx context.Context, event *bun.QueryEvent) context.Context {
	return ctx
}

func (h queryHook) AfterQuery(ctx context.Context, event *bun.QueryEvent) {
	logger := h.st.Logger()
	logger.Debug().
		Dur(logging.RequestDurKey, time.Since(event.StartTime)).
		Msg(event.Query)
}
