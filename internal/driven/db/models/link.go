package models

import (
	"github.com/andreglatz/hopper-go/internal/application/entities"
	sql "github.com/andreglatz/hopper-go/tools/sqlc"
)

type Link struct {
	ID       uint   `db:"ID"`
	Short    string `db:"short"`
	Original string `db:"original"`
}

func NewLink(link sql.Link) *Link {
	return &Link{
		ID:       uint(link.ID),
		Short:    link.Short,
		Original: link.Original,
	}
}

func (l *Link) ToEntity() *entities.Link {
	return &entities.Link{
		ID:       l.ID,
		Short:    l.Short,
		Original: l.Original,
	}
}
