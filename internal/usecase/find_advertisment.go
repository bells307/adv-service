package usecase

import (
	"context"
	"fmt"

	"github.com/bells307/adv-service/internal/domain"
	"github.com/go-playground/validator/v10"
)

type (
	// Получение объявления
	FindAdvertismentUseCase interface {
		Execute(context.Context, FindAdvertismentInput) (FindAdvertismentOutput, error)
	}

	// Найти объявление
	FindAdvertismentInput struct {
		// Идентификатор объявления
		ID string `json:"id" validate:"required" example:"c716e447-b1b8-4266-a64b-d79559da8bf4"`
	}

	// Порт выхода презентера
	FindAdvertismentPresenter interface {
		Output(domain.Advertisment) FindAdvertismentOutput
	}

	// Найденное объявление
	FindAdvertismentOutput struct {
		// Идентификатор
		ID string `json:"id" example:"2765cb06-f750-4d1f-b101-860289786469"`
		// Имя
		Name string `json:"name" example:"Selling the garage"`
		// Имена категорий
		Categories []string `json:"categories" example:"real estate,auto,land"`
		// Описание
		Description string `json:"description" example:"Very big"`
		// Цена
		Price domain.Price `json:"price"`
		// Ссылка на главное фото
		MainPhotoURL string `json:"mainPhotoURL" example:"http://127.0.0.1/storage/main.jpg"`
		// Ссылки на дополнительные фото
		AdditionalPhotoURLs []string `json:"additionalPhotoURLs" example:"http://127.0.0.1/storage/add1.jpg,http://127.0.0.1/storage/add2.jpg"`
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
	validate := validator.New()
	err := validate.Struct(input)
	if err != nil {
		return FindAdvertismentOutput{}, fmt.Errorf("error validating find advertisment input: %v", err)
	}

	adv, err := i.repo.FindByID(ctx, input.ID)
	if err != nil {
		return i.presenter.Output(domain.Advertisment{}), err
	}

	return i.presenter.Output(adv), nil
}
