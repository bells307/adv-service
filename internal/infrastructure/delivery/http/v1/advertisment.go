package v1

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	_ "github.com/bells307/adv-service/docs"
	"github.com/bells307/adv-service/internal/adapter/presenter"
	"github.com/bells307/adv-service/internal/domain"
	"github.com/bells307/adv-service/internal/usecase"
	"github.com/bells307/adv-service/pkg/gin/err_resp"
	"github.com/gin-gonic/gin"
)

type advertismentHandler struct {
	advRepo domain.AdvertismentRepository
	catRepo domain.CategoryRepository
}

func NewAdvertismentHandler(advRepo domain.AdvertismentRepository, catRepo domain.CategoryRepository) *advertismentHandler {
	return &advertismentHandler{advRepo, catRepo}
}

// Зарегистрировать роуты
func (h *advertismentHandler) Register(g *gin.RouterGroup) {
	v1 := g.Group("/v1")
	{
		advertisments := v1.Group("/advertisment")
		{
			advertisments.GET("/:id", h.getAdvertisment)
			advertisments.GET("/summary", h.getAdvertismentSummary)
			advertisments.POST("", h.createAdvertisment)
			advertisments.DELETE("/:id", h.deleteAdvertisment)
		}
	}
}

//	@Summary	Получить информация об объявлении
//	@Tags		advertisment
//	@ID			get-advertisment
//	@Produce	json
//	@Param		id	path		string							true	"ID объявления"
//	@Success	200	{object}	usecase.FindAdvertismentOutput	"Информация об объявлении"
//	@Failure	404	{string}	string							"Объявление не найдено"
//	@Failure	500	{object}	err_resp.ErrorResponse			"Внутренняя ошибка сервиса"
//	@Router		/api/v1/advertisment/{id} [get]
func (h *advertismentHandler) getAdvertisment(c *gin.Context) {
	id := c.Param("id")

	uc := usecase.NewFindAdvertismentInteractor(h.advRepo, presenter.NewFindAdvertismentPresenter())
	out, err := uc.Execute(c.Request.Context(), usecase.FindAdvertismentInput{
		ID: id,
	})
	if err != nil {
		if errors.Is(err, domain.ErrAdvNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		} else {
			err_resp.NewErrorResponse(c, err)
			return
		}

	}

	c.JSON(http.StatusOK, out)
}

//	@Summary	Получить краткую информация об объявлениях с возможностью запроса по страницам
//	@Tags		advertisment
//	@ID			get-advertisment-summary
//	@Produce	json
//	@Param		page	query		integer										false	"Номер страницы"
//	@Success	200		{object}	usecase.FindAllAdvertismentSummaryOutput	"Краткая информация об объявлениях"
//	@Failure	500		{object}	err_resp.ErrorResponse						"Внутренняя ошибка сервиса"
//	@Router		/api/v1/advertisment/summary [get]
func (h *advertismentHandler) getAdvertismentSummary(c *gin.Context) {
	var page *uint
	pageQuery, ok := c.Request.URL.Query()["page"]
	if ok {
		if len(pageQuery) > 1 {
			err_resp.NewErrorResponse(c, errors.New("page query field can't contain more than one value"))
			return
		}
		pageNum, err := strconv.Atoi(pageQuery[0])
		if err != nil {
			err_resp.NewErrorResponse(c, fmt.Errorf("can't parse page query field: %v", err))
			return
		}
		page_ := uint(pageNum)
		page = &page_
	} else {
		page = nil
	}

	uc := usecase.NewFindAllAdvertismentSummaryInteractor(
		h.advRepo,
		presenter.NewFindAllAdvertismentSummaryPresenter(),
	)

	out, err := uc.Execute(c.Request.Context(), usecase.FindAllAdvertismentSummaryInput{
		Page: page,
	})

	if err != nil {
		err_resp.NewErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, out)
}

//	@Summary	Создать объявление
//	@Tags		advertisment
//	@ID			create-advertisment
//	@Accept		json
//	@Produce	json
//	@Param		input	body		usecase.CreateAdvertismentInput		true	"Создание объявления"
//	@Success	200		{object}	usecase.CreateAdvertismentOutput	"Созданное объявление"
//	@Failure	400		{string}	string								"Ошибка формирования запроса"
//	@Failure	500		{object}	err_resp.ErrorResponse				"Внутренняя ошибка сервиса"
//	@Router		/api/v1/advertisment [post]
func (h *advertismentHandler) createAdvertisment(c *gin.Context) {
	var input usecase.CreateAdvertismentInput
	if err := c.Bind(&input); err != nil {
		return
	}

	uc := usecase.NewCreateAdvertismentInteractor(
		h.advRepo,
		h.catRepo,
		presenter.NewCreateAdvertismentPresenter(),
	)
	out, err := uc.Execute(c.Request.Context(), input)
	if err != nil {
		err_resp.NewErrorResponse(c, err)
		return
	}

	// Возвращаем клиенту ID созданного объявления
	c.JSON(http.StatusCreated, gin.H{"id": out.ID})
}

//	@Summary	Удалить объявление
//	@Tags		advertisment
//	@ID			delete-advertisment
//	@Param		id	query		string					false	"Идентификатор объявления"
//	@Success	204	{object}	string					"Удаление произведено успешно"
//	@Failure	500	{object}	err_resp.ErrorResponse	"Внутренняя ошибка сервиса"
//	@Router		/api/v1/advertisment/{id} [delete]
func (h *advertismentHandler) deleteAdvertisment(c *gin.Context) {
	id := c.Param("id")
	if len(id) == 0 {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	input := usecase.DeleteAdvertismentInput{
		ID: id,
	}

	uc := usecase.NewDeleteAdvertismentInteractor(h.advRepo)
	if err := uc.Execute(c.Request.Context(), input); err != nil {
		err_resp.NewErrorResponse(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}
