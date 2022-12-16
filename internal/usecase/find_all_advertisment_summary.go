package usecase

import (
	"context"
	"fmt"

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
		Category string `json:"category"`
		// Цена
		Price domain.Price `json:"price"`
		// Ссылка на главное фото
		MainPhotoURL string `json:"mainPhotoURL"`
	}

	findAllAdvertismentSummaryInteractor struct {
		advRepo   domain.AdvertismentRepository
		catRepo   domain.CategoryRepository
		presenter FindAllAdvertismentSummaryPresenter
	}
)

func NewFindAllAdvertismentSummaryInteractor(
	advRepo domain.AdvertismentRepository,
	catRepo domain.CategoryRepository,
	presenter FindAllAdvertismentSummaryPresenter,
) FindAllAdvertismentSummaryUseCase {
	return findAllAdvertismentSummaryInteractor{
		advRepo,
		catRepo,
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

	// Получаем категории
	var categoryIDs []string
	for _, adv := range advs {
		categoryIDs = append(categoryIDs, adv.CategoryID)
	}

	categories, err := i.catRepo.FindAllByID(ctx, categoryIDs)
	if err != nil {
		return i.presenter.Output([]domain.AdvertismentSummary{}), err
	}

	var summary []domain.AdvertismentSummary
	for _, adv := range advs {
		category, found := findCategory(categories, adv.CategoryID)
		if !found {
			return i.presenter.Output(
					[]domain.AdvertismentSummary{}),
				fmt.Errorf("can't find category id %s", adv.CategoryID)
		}

		summary = append(summary, domain.AdvertismentSummary{
			Name:         adv.Name,
			Category:     category,
			Price:        adv.Price,
			MainPhotoURL: adv.MainPhotoURL,
		})
	}

	return i.presenter.Output(summary), nil
}

func findCategory(categories []domain.Category, id string) (domain.Category, bool) {
	for _, cat := range categories {
		if cat.ID == id {
			return cat, true
		}
	}
	return domain.Category{}, false
}
