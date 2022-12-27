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
	advRepo domain.AdvertismentRepository
	catRepo domain.CategoryRepository
}

func NewCategoryHandler(advRepo domain.AdvertismentRepository, catRepo domain.CategoryRepository) *categoryHandler {
	return &categoryHandler{advRepo, catRepo}
}

// Зарегистрировать роуты
func (h *categoryHandler) Register(g *gin.RouterGroup) {
	v1 := g.Group("/v1")
	{
		category := v1.Group("/category")
		{
			category.POST("", h.createCategory)
			category.DELETE("/:id", h.deleteCategory)
		}
	}
}

func (h *categoryHandler) createCategory(c *gin.Context) {
	var input usecase.CreateCategoryInput
	if err := c.Bind(&input); err != nil {
		return
	}

	uc := usecase.NewCreateCategoryInteractor(h.catRepo, presenter.NewCreateCategoryPresenter())
	out, err := uc.Execute(c.Request.Context(), input)
	if err != nil {
		err_resp.NewErrorResponse(c, err)
		return
	}

	// Возвращаем клиенту ID созданного объявления
	c.JSON(http.StatusCreated, gin.H{"id": out.ID})
}

func (h *categoryHandler) deleteCategory(c *gin.Context) {
	id := c.Param("id")
	if len(id) == 0 {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	input := usecase.DeleteCategoryInput{
		ID: id,
	}

	uc := usecase.NewDeleteCategoryInteractor(h.advRepo, h.catRepo)
	err := uc.Execute(c.Request.Context(), input)
	if err != nil {
		err_resp.NewErrorResponse(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}
