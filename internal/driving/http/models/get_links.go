package models

import (
	"github.com/andreglatz/hopper-go/internal/application/types"
	usecases "github.com/andreglatz/hopper-go/internal/application/use_cases"
)

type GetLinksFilters struct {
	Offset   int32  `form:"offset"`
	Limit    int32  `form:"limit"`
	Short    string `form:"short"`
	Original string `form:"original"`
}

func (g GetLinksFilters) ToGetLinksParams() usecases.GetLinksParams {
	return usecases.GetLinksParams{
		Offset:   g.Offset,
		Limit:    g.Limit,
		Short:    g.Short,
		Original: g.Original,
	}
}

type GetLinksResponse struct {
	Pagination Pagination `json:"pagination"`
	Links      []Link     `json:"data"`
}

func NewGetLinksResponse(links []types.Link, pagination types.Pagination) GetLinksResponse {
	var linksResponse []Link

	for _, link := range links {
		linksResponse = append(linksResponse, NewLink(link))
	}

	return GetLinksResponse{
		Links:      linksResponse,
		Pagination: NewPagination(pagination),
	}
}
