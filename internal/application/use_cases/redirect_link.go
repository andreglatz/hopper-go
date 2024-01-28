package usecases

import "github.com/andreglatz/hopper-go/internal/driven/db/repositories"

type RedirectLink interface {
	GetLink(path string) (string, error)
}

type RedirectLinkUseCase struct {
	repository repositories.LinkRepository
}

func NewRedirectLinkUseCase(r repositories.LinkRepository) RedirectLink {
	return &RedirectLinkUseCase{
		repository: r,
	}
}

func (u *RedirectLinkUseCase) GetLink(path string) (string, error) {
	link, err := u.repository.GetByShort(path)
	if err != nil {
		return "", err
	}

	link.AddClick()

	err = u.repository.Update(link)
	if err != nil {
		return "", err
	}

	return link.Original, nil
}
