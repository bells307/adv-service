package usecase

import (
	"context"

	"github.com/bells307/adv-service/internal/domain"
)

type (
	// Получение списка объявлений с краткой информацией
	FindAllAdvertismentSummaryUseCase interface {
		Execute(context.Context, FindAllAdvertismentSummaryInput) (FindAllAdvertismentSummaryOutput, error)
	}

	// Найти краткую информацию по объявлениям
	FindAllAdvertismentSummaryInput struct {
		// Номер страницы
		// Если nil - передаются все страницы
		Page *uint `json:"page"`
	}

	// Порт выхода презентера
	FindAllAdvertismentSummaryPresenter interface {
		Output([]domain.AdvertismentSummary) FindAllAdvertismentSummaryOutput
	}

	// Объявления на странице
	FindAllAdvertismentSummaryOutput []struct {
		// Название объявления
		Name string `json:"name"`
		// Категория
		Categories []string `json:"categories"`
		// Цена
		Price domain.Price `json:"price"`
		// Ссылка на главное фото
		MainPhotoURL string `json:"mainPhotoURL"`
	}

	findAllAdvertismentSummaryInteractor struct {
		advRepo   domain.AdvertismentRepository
		presenter FindAllAdvertismentSummaryPresenter
	}
)

func NewFindAllAdvertismentSummaryInteractor(
	advRepo domain.AdvertismentRepository,
	presenter FindAllAdvertismentSummaryPresenter,
) FindAllAdvertismentSummaryUseCase {
	return findAllAdvertismentSummaryInteractor{
		advRepo,
		presenter,
	}
}

func (i findAllAdvertismentSummaryInteractor) Execute(ctx context.Context, input FindAllAdvertismentSummaryInput) (FindAllAdvertismentSummaryOutput, error) {
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

	advs, err := i.advRepo.Find(ctx, limit, offset)
	if err != nil {
		return i.presenter.Output([]domain.AdvertismentSummary{}), err
	}

	var summaries []domain.AdvertismentSummary
	for _, adv := range advs {
		summaries = append(summaries, adv.Summarize())
	}

	return i.presenter.Output(summaries), nil
}
