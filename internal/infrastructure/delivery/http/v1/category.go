package v1

import (
	"net/http"

	"github.com/bells307/adv-service/internal/usecase"
	"github.com/bells307/adv-service/pkg/gin/err_resp"
	"github.com/gin-gonic/gin"
)

type categoryHandler struct {
	catUsecase *usecase.CategoryUsecase
}

func NewCategoryHandler(catUsecase *usecase.CategoryUsecase) *categoryHandler {
	return &categoryHandler{catUsecase}
}

// Зарегистрировать роуты
func (h *categoryHandler) Register(g *gin.RouterGroup) {
	v1 := g.Group("/v1")
	{
		advertisments := v1.Group("/advertisment")
		{
			advertisments.POST("", h.createCategory)
		}
	}
}

func (h *categoryHandler) createCategory(c *gin.Context) {
	var createCategory usecase.CreateCategory
	if err := c.Bind(&createCategory); err != nil {
		return
	}

	cat, err := h.catUsecase.CreateCategory(c.Request.Context(), &createCategory)
	if err != nil {
		err_resp.ErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, cat)
}
