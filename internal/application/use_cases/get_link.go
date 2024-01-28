package usecases

import "github.com/andreglatz/hopper-go/internal/driven/db/repositories"

type GetLink interface {
	GetLink(path string) (string, error)
}

type getLinkUseCase struct {
	repository repositories.LinkRepository
}

func NewGetLinkUseCase(r repositories.LinkRepository) GetLink {
	return &getLinkUseCase{
		repository: r,
	}
}

func (u *getLinkUseCase) GetLink(path string) (string, error) {
	link, err := u.repository.GetByShort(path)
	if err != nil {
		return "", err
	}

	return link.Original, nil
}
