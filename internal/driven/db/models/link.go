package models

import "github.com/andreglatz/hopper-go/internal/application/entities"

type Link struct {
	ID       uint   `db:"ID"`
	Short    string `db:"short"`
	Original string `db:"original"`
}

func NewLink(link entities.Link) *Link {
	return &Link{
		ID:       link.ID,
		Short:    link.Short,
		Original: link.Original,
	}
}
