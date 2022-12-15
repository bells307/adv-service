package usecase

import (
	"context"

	"github.com/bells307/adv-service/internal/domain"
)

type (
	// Получение списка объявлений с краткой информацией по номеру страницы
	FindAllAdvertismentSummaryByPageUseCase interface {
		Execute(context.Context, FindAllAdvertismentSummaryByPageInput) (FindAllAdvertismentSummaryByPageOutput, error)
	}

	// Найти все объявления по номеру страницы
	FindAllAdvertismentSummaryByPageInput struct {
		// Номер страницы
		// Если nil - передаются все страницы
		Page *uint `json:"page"`
	}

	// Порт выхода презентера
	FindAllAdvertismentSummaryByPagePresenter interface {
		Output([]domain.AdvertismentSummary) FindAllAdvertismentSummaryByPageOutput
	}

	// Объявления на странице
	FindAllAdvertismentSummaryByPageOutput []struct {
		// Название объявления
		Name string `json:"name"`
		// Цена
		Price domain.Price `json:"price"`
		// Ссылка на главное фото
		MainPhotoURL string `json:"mainPhotoURL"`
	}

	findAllAdvertismentSummaryByPageInteractor struct {
		repo      domain.AdvertismentRepository
		presenter FindAllAdvertismentSummaryByPagePresenter
	}
)

func NewFindAllAdvertismentSummaryByPageInteractor(
	repo domain.AdvertismentRepository,
	presenter FindAllAdvertismentSummaryByPagePresenter,
) FindAllAdvertismentSummaryByPageUseCase {
	return findAllAdvertismentSummaryByPageInteractor{
		repo,
		presenter,
	}
}

func (i findAllAdvertismentSummaryByPageInteractor) Execute(ctx context.Context, input FindAllAdvertismentSummaryByPageInput) (FindAllAdvertismentSummaryByPageOutput, error) {
	// Вычисляем limit и offset по номеру страницы
	var limit uint
	var offset uint

	if input.Page == nil {
		// Если nil - передаются все страницы
		limit = 0
		offset = 0
	} else {
		limit = domain.ADV_ON_PAGE_COUNT
		offset = limit * (uint(*input.Page) - 1)
	}

	advs, err := i.repo.Find(ctx, limit, offset)
	if err != nil {
		return i.presenter.Output([]domain.AdvertismentSummary{}), err
	}

	var summary []domain.AdvertismentSummary
	for _, advs := range advs {
		summary = append(summary, advs.Summary())
	}

	return i.presenter.Output(summary), nil
}
