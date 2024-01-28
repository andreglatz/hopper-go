package repositories

import (
	"github.com/andreglatz/hopper-go/internal/application/entities"
)

type LinkRepository interface {
	Save(*entities.Link) error
}

type InMemoryLinkRepository struct {
	links map[uint]entities.Link
}

func NewInMemoryLinkRepository() LinkRepository {
	return &InMemoryLinkRepository{
		links: make(map[uint]entities.Link),
	}
}

func (r *InMemoryLinkRepository) Save(link *entities.Link) error {
	lastId := uint(len(r.links))

	link.Id = lastId + 1

	r.links[link.Id] = *link

	return nil
}
