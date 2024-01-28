package models

import (
	"github.com/andreglatz/hopper-go/internal/application/types"
)

type Link struct {
	ID       uint   `json:"id"`
	Short    string `json:"short"`
	Original string `json:"original"`
}

func NewLink(link types.Link) Link {
	return Link{
		ID:       link.ID,
		Short:    link.Short,
		Original: link.Original,
	}
}
