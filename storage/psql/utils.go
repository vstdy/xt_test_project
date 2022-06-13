package psql

import (
	"github.com/uptrace/bun"
	"strings"

	"github.com/vstdy/xt_test_project/pkg/input"
)

// paginateQuery adds pagination parameters to select query.
func paginateQuery(q *bun.SelectQuery, params input.PageParams) {
	q.Limit(params.Limit)
	q.Offset(params.Offset)
}

// dateTimeFilterQuery adds datetime parameters to select query.
func dateTimeFilterQuery(q *bun.SelectQuery, filterName string, dtparams input.DateTimeParams) {
	if !dtparams.Since.IsZero() {
		q.Where("? >= ?", bun.Ident(filterName), dtparams.Since)
	}
	if !dtparams.Till.IsZero() {
		q.Where("? <= ?", bun.Ident(filterName), dtparams.Till)
	}
}

// columnsFilterQuery retrieves passed column from the query.
func columnFilterQuery(q *bun.SelectQuery, column string, extras ...string) {
	if column != "" {
		q.Column(extras...)
		q.Column(strings.ToLower(column))
	}
}
