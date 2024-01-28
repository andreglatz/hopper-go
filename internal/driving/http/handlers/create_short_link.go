package handlers

import (
	"errors"
	"io"
	"net/http"

	usecases "github.com/andreglatz/hopper-go/internal/application/use_cases"
	"github.com/andreglatz/hopper-go/internal/driving/http/models"
	"github.com/andreglatz/hopper-go/internal/utils"

	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
)

type CreateShortLinkHandler struct {
	createShortLinkUseCase usecases.CreteShortLink
	translator             ut.Translator
}

func NewCreateShortLinkHandler(t ut.Translator, u usecases.CreteShortLink) *CreateShortLinkHandler {
	return &CreateShortLinkHandler{
		createShortLinkUseCase: u,
		translator:             t,
	}
}

func (h *CreateShortLinkHandler) Handle(ctx *gin.Context) {
	payload := models.CreateShortLinkPayload{}

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		if errors.Is(err, io.EOF) {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
			return
		}

		ctx.JSON(http.StatusBadRequest, gin.H{"error": utils.TranslateError(err, h.translator)})
		return
	}

	link, err := h.createShortLinkUseCase.Create(payload.ToUseCase())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, models.NewCreateShortLinkResponse(link))
}
