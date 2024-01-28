package types

import "github.com/andreglatz/hopper-go/internal/application/entities"

type Link struct {
	ID       uint
	Original string
	Short    string
}

func NewLink(link entities.Link) Link {
	return Link{
		ID:       link.ID,
		Original: link.Original,
		Short:    link.Short,
	}
}

type Pagination struct {
	Offset int32
	Limit  int32
	Total  uint
}
