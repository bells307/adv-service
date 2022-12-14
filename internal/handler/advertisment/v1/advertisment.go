package advertisment_handler

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	adv_service "github.com/bells307/adv-service/internal/domain/advertisment/service"
	adv_dto "github.com/bells307/adv-service/internal/domain/advertisment/service/dto"
	adv_err "github.com/bells307/adv-service/internal/domain/advertisment/service/error"
	"github.com/gin-gonic/gin"
)

type advertismentHandler struct {
	advService *adv_service.AdvertismentService
}

// Количество элементов на одной странице объявлений
// TODO: в конфигурацию
const PAGE_ELEMENT_COUNT = 10

func NewAdvertismentHandler(advService *adv_service.AdvertismentService) *advertismentHandler {
	return &advertismentHandler{advService}
}

// Зарегистрировать роуты
func (h *advertismentHandler) Register(g *gin.RouterGroup) {
	v1 := g.Group("/v1")
	{
		advertisments := v1.Group("/advertisment")
		{
			advertisments.GET("/:id", h.getAdvertisment)
			advertisments.GET("", h.getAdvertismentMany)
			advertisments.POST("", h.createAdvertisment)
		}
	}
}

// Получить информацию об объявлении
func (h *advertismentHandler) getAdvertisment(c *gin.Context) {
	// Получаем объявление из сервиса
	id := c.Param("id")
	adv, err := h.advService.GetOne(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, adv_err.ErrNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
		} else {
			ErrorResponse(c, err)
		}
		return
	}

	// Ответ на запрос
	resp := GetAdvertismentResponse{
		Name:                adv.Name,
		Price:               adv.Price,
		MainPhotoURL:        adv.MainPhotoURL,
		Description:         nil,
		AdditionalPhotoURLs: nil,
	}

	// Найти значение `value` в query запроса по ключу `fields`
	findField := func(query url.Values, value string) bool {
		fields, found := query["fields"]
		if !found {
			return false
		}

		for i := 0; i < len(fields); i++ {
			if fields[i] == value {
				return true
			}
		}

		return false
	}

	// Добавляем в ответ описание только если клиент запросил это поле
	if findField(c.Request.URL.Query(), "description") {
		resp.Description = &adv.Description
	}

	// Добавляем в ответ ссылки на дополнительные фото только если клиент запросил это поле
	if findField(c.Request.URL.Query(), "additionalPhotoURLs") {
		resp.AdditionalPhotoURLs = &adv.AdditionalPhotoURLs
	}

	c.JSON(http.StatusOK, resp)
}

func (h *advertismentHandler) getAdvertismentMany(c *gin.Context) {
	limit := 0
	offset := 0

	page, ok := c.Request.URL.Query()["page"]
	if ok {
		if len(page) > 1 {
			ErrorResponse(c, errors.New("page query field can't contain more than one value"))
			return
		}

		pageNum, err := strconv.Atoi(page[0])
		if err != nil {
			ErrorResponse(c, fmt.Errorf("can't parse page query field: %v", err))
			return
		}

		limit = PAGE_ELEMENT_COUNT
		offset = limit * (pageNum - 1)
	}

	dto := adv_dto.GetAdvertisments{
		Limit:  int64(limit),
		Offset: int64(offset),
	}

	advs, err := h.advService.Get(c.Request.Context(), &dto)
	if err != nil {
		ErrorResponse(c, err)
		return
	}

	var res []GetAdvertismentManyResponse
	for _, adv := range advs {
		res = append(res, GetAdvertismentManyResponse{
			Name:         adv.Name,
			Price:        adv.Price,
			MainPhotoURL: adv.MainPhotoURL,
		})
	}

	c.JSON(http.StatusOK, res)
}

// Создать объявление
func (h *advertismentHandler) createAdvertisment(c *gin.Context) {
	var createAdvertisment adv_dto.CreateAdvertisment
	if err := c.Bind(&createAdvertisment); err != nil {
		return
	}

	// Создаем объявление
	adv, err := h.advService.Create(c.Request.Context(), &createAdvertisment)
	if err != nil {
		ErrorResponse(c, err)
		return
	}

	// Возвращаем клиенту ID созданного объявления
	c.JSON(http.StatusCreated, gin.H{"id": adv.ID})
}
