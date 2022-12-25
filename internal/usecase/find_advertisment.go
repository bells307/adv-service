package usecase

import (
	"context"

	"github.com/bells307/adv-service/internal/domain"
)

type (
	// Получение объявления
	FindAdvertismentUseCase interface {
		Execute(context.Context, FindAdvertismentInput) (FindAdvertismentOutput, error)
	}

	// Найти объявление
	FindAdvertismentInput struct {
		// Идентификатор объявления
		ID string `json:"id"`
	}

	// Порт выхода презентера
	FindAdvertismentPresenter interface {
		Output(domain.Advertisment) FindAdvertismentOutput
	}

	// Найденное объявление
	FindAdvertismentOutput struct {
		// Идентификатор
		ID string `json:"id"`
		// Имя
		Name string `json:"name"`
		// Имена категорий
		Categories []string `json:"categories"`
		// Описание
		Description string `json:"description"`
		// Цена
		Price domain.Price `json:"price"`
		// Ссылка на главное фото
		MainPhotoURL string `json:"mainPhotoURL"`
		// Ссылки на дополнительные фото
		AdditionalPhotoURLs []string `json:"additionalPhotoURLs"`
	}

	findAdvertismentInteractor struct {
		repo      domain.AdvertismentRepository
		presenter FindAdvertismentPresenter
	}
)

func NewFindAdvertismentInteractor(
	repo domain.AdvertismentRepository,
	presenter FindAdvertismentPresenter,
) findAdvertismentInteractor {
	return findAdvertismentInteractor{
		repo,
		presenter,
	}
}

func (i findAdvertismentInteractor) Execute(ctx context.Context, input FindAdvertismentInput) (FindAdvertismentOutput, error) {
	adv, err := i.repo.FindByID(ctx, input.ID)
	if err != nil {
		return i.presenter.Output(domain.Advertisment{}), err
	}

	return i.presenter.Output(adv), nil
}
