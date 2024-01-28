package models

import (
	"github.com/andreglatz/hopper-go/internal/application/types"
)

type Pagination struct {
	Total  uint  `json:"total"`
	Offset int32 `json:"offset"`
	Limit  int32 `json:"limit"`
}

func NewPagination(p types.Pagination) Pagination {
	return Pagination{
		Total:  p.Total,
		Offset: p.Offset,
		Limit:  p.Limit,
	}
}
