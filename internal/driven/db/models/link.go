package models

import (
	"github.com/andreglatz/hopper-go/internal/application/entities"
	sql "github.com/andreglatz/hopper-go/tools/sqlc"
)

type Link struct {
	ID       uint
	Short    string
	Original string
	Clicks   uint
}

func NewLink(link sql.Link) *Link {
	return &Link{
		ID:       uint(link.ID),
		Short:    link.Short,
		Original: link.Original,
		Clicks:   uint(link.Clicks),
	}
}

func (l *Link) ToEntity() *entities.Link {
	return &entities.Link{
		ID:       l.ID,
		Short:    l.Short,
		Original: l.Original,
		Clicks:   l.Clicks,
	}
}
