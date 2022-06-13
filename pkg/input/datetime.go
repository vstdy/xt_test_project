package input

import (
	"fmt"
	"github.com/vstdy/xt_test_project/pkg"
	"time"
)

const (
	Datetime = "2006-01-02 15:04:05"
	Date     = "2006-01-02"
)

// DateTimeParams keeps query datetime params.
type DateTimeParams struct {
	Since time.Time
	Till  time.Time
}

// NewDateTimeParams creates a new DateTimeParams.
func NewDateTimeParams(since, till, layout string) (DateTimeParams, error) {
	var s time.Time
	var t time.Time
	var err error

	if since != "" {
		s, err = time.Parse(layout, since)
		if err != nil {
			return DateTimeParams{}, fmt.Errorf("since: wrong formant, must be %s", layout)
		}
	}
	if till != "" {
		t, err = time.Parse(layout, till)
		if err != nil {
			return DateTimeParams{}, fmt.Errorf("till: wrong formant, must be %s", layout)
		}
	}

	if !s.IsZero() && !t.IsZero() {
		if s.After(t) {
			return DateTimeParams{}, fmt.Errorf("%w: range start must be LT end", pkg.ErrInvalidInput)
		}
	}

	dt := DateTimeParams{
		Since: s,
		Till:  t,
	}

	return dt, nil
}
