package handler

import (
	"context"

	"github.com/bells307/adv-service/internal/dto"
	"github.com/bells307/adv-service/internal/model"
	"github.com/gin-gonic/gin"
)

type advertismentHandler struct {
	advService advertismentService
}

type advertismentService interface {
	Get(ctx context.Context) ([]model.Advertisment, error)
	GetOne(ctx context.Context, id string) (model.Advertisment, error)
	Create(ctx context.Context, createAdv dto.CreateAdvertisment) (adv model.Advertisment, err error)
}

func NewAdvertismentHandler(advService advertismentService) *advertismentHandler {
	return &advertismentHandler{advService}
}

func (h *advertismentHandler) Register(e *gin.Engine) {
	advertisments := e.Group("/advertisment")
	{
		advertisments.GET("/:id", h.getAdvertisment)
	}
}

func (h *advertismentHandler) getAdvertisment(c *gin.Context) {

}
