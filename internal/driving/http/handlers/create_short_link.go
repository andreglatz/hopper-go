package handlers

import (
	"errors"
	"io"
	"net/http"

	usecases "github.com/andreglatz/hopper-go/internal/application/use_cases"
	"github.com/andreglatz/hopper-go/internal/driving/http/models"
	"github.com/andreglatz/hopper-go/internal/utils"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
)

type CreateShortLinkHandler struct {
	createShortLinkUseCase usecases.CreteShortLink
	translator             ut.Translator
	logger                 *zap.SugaredLogger
}

func NewCreateShortLinkHandler(t ut.Translator, l *zap.SugaredLogger, u usecases.CreteShortLink) *CreateShortLinkHandler {
	return &CreateShortLinkHandler{
		createShortLinkUseCase: u,
		translator:             t,
		logger:                 l,
	}
}

func (h *CreateShortLinkHandler) Handle(ctx *gin.Context) {
	payload := models.CreateShortLinkPayload{}

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		h.logger.Warnw("Error while binding JSON", "error", err)

		if errors.Is(err, io.EOF) {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
			return
		}

		ctx.JSON(http.StatusBadRequest, gin.H{"error": utils.TranslateError(err, h.translator)})
		return
	}

	link, err := h.createShortLinkUseCase.Create(payload.ToUseCase())
	if err != nil {
		h.logger.Errorw("Error while creating short link", "error", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	ctx.JSON(http.StatusOK, models.NewCreateShortLinkResponse(link))
}
