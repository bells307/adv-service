package usecase

import (
	"context"
	"fmt"

	"github.com/bells307/adv-service/internal/domain"
	_ "github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type (
	// Создание объявления
	CreateAdvertismentUseCase interface {
		Execute(context.Context, CreateAdvertismentInput) (CreateAdvertismentOutput, error)
	}

	// Создать объявление
	CreateAdvertismentInput struct {
		// Имя
		Name string `json:"name" validate:"required"`
		// ID категорий
		Categories []string `json:"categories" validate:"required"`
		// Описание
		Description string `json:"description"`
		// Цена
		Price domain.Price `json:"price" validate:"required"`
		// Ссылка на главное фото
		MainPhotoURL string `json:"mainPhotoURL"`
		// Ссылки на дополнительные фото
		AdditionalPhotoURLs []string `json:"additionalPhotoURLs"`
	}

	// Порт выхода презентера
	CreateAdvertismentPresenter interface {
		Output(domain.Advertisment) CreateAdvertismentOutput
	}

	// Созданное объявление
	CreateAdvertismentOutput struct {
		// Идентификатор
		ID string `json:"id"`
		// Имя
		Name string `json:"name"`
		// Идентификаторы категорий
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

	createAdvertismentInteractor struct {
		advRepo   domain.AdvertismentRepository
		catRepo   domain.CategoryRepository
		presenter CreateAdvertismentPresenter
	}
)

func NewCreateAdvertismentInteractor(
	advRepo domain.AdvertismentRepository,
	catRepo domain.CategoryRepository,
	presenter CreateAdvertismentPresenter,
) CreateAdvertismentUseCase {
	return createAdvertismentInteractor{
		advRepo,
		catRepo,
		presenter,
	}
}

func (i createAdvertismentInteractor) Execute(ctx context.Context, input CreateAdvertismentInput) (CreateAdvertismentOutput, error) {
	categories, err := i.catRepo.FindAllByID(ctx, input.Categories)
	if err != nil {
		return CreateAdvertismentOutput{}, fmt.Errorf("can't find categories %s: %v", input.Categories, err)
	}

	adv := domain.Advertisment{
		ID:                  uuid.NewString(),
		Name:                input.Name,
		Categories:          categories,
		Description:         input.Description,
		Price:               input.Price,
		MainPhotoURL:        input.MainPhotoURL,
		AdditionalPhotoURLs: input.AdditionalPhotoURLs,
	}

	if err := adv.Validate(); err != nil {
		return CreateAdvertismentOutput{}, err
	}

	if err := i.advRepo.Create(ctx, adv); err != nil {
		return CreateAdvertismentOutput{}, fmt.Errorf("failed to create advertisment: %v", err)
	}

	return i.presenter.Output(adv), nil
}
