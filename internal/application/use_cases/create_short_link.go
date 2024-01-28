package usecases

import (
	"github.com/andreglatz/hopper-go/internal/application/entities"
	"github.com/andreglatz/hopper-go/internal/driven/db/repositories"
)

type CreateShortLinkParams struct {
	Original string
	Short    string
}

type CreateLinkResponse struct {
	Id       uint
	Original string
	Short    string
}

func newCreateResponse(link *entities.Link) CreateLinkResponse {
	return CreateLinkResponse{
		Id:       link.ID,
		Original: link.Original,
		Short:    link.Short,
	}
}

type CreteShortLink interface {
	Create(CreateShortLinkParams) (CreateLinkResponse, error)
}

type createShortLinkUseCase struct {
	linkRepository repositories.LinkRepository
}

func NewCreateShortLinkUseCase(r repositories.LinkRepository) CreteShortLink {
	return &createShortLinkUseCase{
		linkRepository: r,
	}
}

func (c *createShortLinkUseCase) Create(params CreateShortLinkParams) (CreateLinkResponse, error) {
	link := entities.NewLink(params.Original, params.Short)

	if err := c.linkRepository.Save(&link); err != nil {
		return CreateLinkResponse{}, err
	}

	response := newCreateResponse(&link)

	return response, nil
}
