package v1

import (
	"net/http"

	"github.com/bells307/adv-service/internal/adapter/presenter"
	"github.com/bells307/adv-service/internal/domain"
	"github.com/bells307/adv-service/internal/usecase"
	"github.com/bells307/adv-service/pkg/gin/err_resp"
	"github.com/gin-gonic/gin"
)

type categoryHandler struct {
	repo domain.CategoryRepository
}

func NewCategoryHandler(repo domain.CategoryRepository) *categoryHandler {
	return &categoryHandler{repo}
}

// Зарегистрировать роуты
func (h *categoryHandler) Register(g *gin.RouterGroup) {
	v1 := g.Group("/v1")
	{
		advertisments := v1.Group("/category")
		{
			advertisments.POST("", h.createCategory)
		}
	}
}

func (h *categoryHandler) createCategory(c *gin.Context) {
	var input usecase.CreateCategoryInput
	if err := c.Bind(&input); err != nil {
		return
	}

	uc := usecase.NewCreateCategoryInteractor(h.repo, presenter.NewCreateCategoryPresenter())
	out, err := uc.Execute(c.Request.Context(), input)
	if err != nil {
		err_resp.ErrorResponse(c, err)
		return
	}

	// Возвращаем клиенту ID созданного объявления
	c.JSON(http.StatusCreated, gin.H{"id": out.ID})
}
