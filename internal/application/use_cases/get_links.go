package usecases

import (
	"github.com/andreglatz/hopper-go/internal/application/entities"
	"github.com/andreglatz/hopper-go/internal/application/types"
	"github.com/andreglatz/hopper-go/internal/driven/db/repositories"
)

type GetLinksResponse struct {
	Links      []types.Link
	Pagination types.Pagination
}

type GetLinksParams struct {
	Offset   int32
	Limit    int32
	Short    string
	Original string
}

type GetLinks interface {
	GetLinks(GetLinksParams) (GetLinksResponse, error)
}

type GetLinksUseCase struct {
	repository repositories.LinkRepository
}

func NewGetLinksUseCase(r repositories.LinkRepository) GetLinks {
	return &GetLinksUseCase{
		repository: r,
	}
}

func (u *GetLinksUseCase) GetLinks(query GetLinksParams) (GetLinksResponse, error) {
	params := formatDBParams(query)

	links, count, err := u.repository.GetLinks(params)
	if err != nil {
		return GetLinksResponse{}, err
	}

	response := GetLinksResponse{
		Links: toLinks(links),
		Pagination: types.Pagination{
			Total:  count,
			Offset: params.Offset,
			Limit:  params.Limit,
		},
	}

	return response, nil
}

func formatDBParams(query GetLinksParams) repositories.GetLinksParams {
	params := repositories.GetLinksParams{
		Offset:   query.Offset,
		Limit:    query.Limit,
		Short:    query.Short,
		Original: query.Original,
	}

	if query.Limit > 50 {
		params.Limit = 50
	}

	if params.Limit == 0 {
		params.Limit = 10
	}

	return params
}

func toLinks(links []*entities.Link) []types.Link {
	var result []types.Link

	for _, link := range links {
		result = append(result, types.NewLink(*link))
	}

	return result
}
