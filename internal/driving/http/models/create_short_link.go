package models

import usecases "github.com/andreglatz/hopper-go/internal/application/use_cases"

type CreateShortLinkPayload struct {
	Short    string `json:"short" binding:"required,min=3,max=10"`
	Original string `json:"original" binding:"required"`
}

func (b *CreateShortLinkPayload) ToUseCase() usecases.CreateShortLinkParams {
	return usecases.CreateShortLinkParams{
		Short:    b.Short,
		Original: b.Original,
	}
}

type CreateShortLinkResponse struct {
	Id       uint   `json:"id"`
	Short    string `json:"short"`
	Original string `json:"original"`
}

func NewCreateShortLinkResponse(link usecases.CreateLinkResponse) CreateShortLinkResponse {
	return CreateShortLinkResponse{
		Id:       link.Id,
		Short:    link.Short,
		Original: link.Original,
	}
}
