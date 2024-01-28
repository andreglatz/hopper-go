package handlers

import (
	"net/http"

	usecases "github.com/andreglatz/hopper-go/internal/application/use_cases"
	"github.com/andreglatz/hopper-go/internal/driving/http/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type RedirectLinkHandler struct {
	getLinkUseCase usecases.GetLink
	logger         *zap.SugaredLogger
}

func NewRedirectLinkHandler(l *zap.SugaredLogger, u usecases.GetLink) *RedirectLinkHandler {
	return &RedirectLinkHandler{
		getLinkUseCase: u,
		logger:         l,
	}
}

func (h *RedirectLinkHandler) Handle(ctx *gin.Context) {
	path := models.RedirectLinkPath{
		Short: ctx.Param("short"),
	}

	link, err := h.getLinkUseCase.GetLink(path.Short)
	if err != nil {
		h.logger.Errorw("Error while getting link", "error", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	ctx.Redirect(http.StatusMovedPermanently, link)
}
