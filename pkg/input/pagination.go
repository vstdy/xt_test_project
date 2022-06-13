package input

import (
	"fmt"
)

const (
	PageLimit = 50
)

// PageParams keeps query pagination params.
type PageParams struct {
	Offset int
	Limit  int
}

// NewPageParams creates a new PageParams.
func NewPageParams(pageNumber, pageSize int) (PageParams, error) {
	if pageNumber < 0 {
		return PageParams{}, fmt.Errorf("page number: must be GTE 1")
	}
	if pageSize < 0 || pageSize > PageLimit {
		return PageParams{}, fmt.Errorf("page size: must be GTE 1 and LTE %d", PageLimit)
	}
	if pageNumber == 0 {
		pageNumber = 1
	}
	if pageSize == 0 {
		pageSize = 10
	}

	pp := PageParams{
		Offset: (pageNumber - 1) * pageSize,
		Limit:  pageSize,
	}

	return pp, nil
}
