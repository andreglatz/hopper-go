package handlers

import (
	"net/http"

	usecases "github.com/andreglatz/hopper-go/internal/application/use_cases"
	"github.com/andreglatz/hopper-go/internal/driving/http/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type GetLinksHandler struct {
	getLinkUseCase usecases.GetLinks
	logger         *zap.SugaredLogger
}

func NewGetLinksHandler(l *zap.SugaredLogger, u usecases.GetLinks) *GetLinksHandler {
	return &GetLinksHandler{
		getLinkUseCase: u,
		logger:         l,
	}
}

func (h *GetLinksHandler) Handle(ctx *gin.Context) {
	query := models.GetLinksFilters{}

	if err := ctx.ShouldBindQuery(&query); err != nil {
		h.logger.Errorw("Error while binding query", "error", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query"})
		return
	}

	response, err := h.getLinkUseCase.GetLinks(query.ToGetLinksParams())
	if err != nil {
		h.logger.Errorw("Error while getting links", "error", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	ctx.JSON(http.StatusOK, models.NewGetLinksResponse(response.Links, response.Pagination))
}
