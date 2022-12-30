package usecase

import (
	"context"
	"fmt"

	"github.com/bells307/adv-service/internal/domain"
	"github.com/go-playground/validator/v10"
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
		Name string `json:"name" validate:"required" example:"Selling the garage"`
		// ID категорий
		Categories []string `json:"categories" validate:"required" example:"a9f9ecfe-25b5-4742-901a-21dee231f6cf,d39c79b6-78e0-41fc-939d-e60a64c0251e,d9ff247b-0469-4106-af73-89e7a966e4a4"`
		// Описание
		Description string `json:"description" example:"Very big"`
		// Цена
		Price Price `json:"price" validate:"required"`
		// Ссылка на главное фото
		MainPhotoURL string `json:"mainPhotoURL" example:"http://127.0.0.1/storage/main.jpg"`
		// Ссылки на дополнительные фото
		AdditionalPhotoURLs []string `json:"additionalPhotoURLs" example:"http://127.0.0.1/storage/add1.jpg,http://127.0.0.1/storage/add2.jpg"`
	}

	// Порт выхода презентера
	CreateAdvertismentPresenter interface {
		Output(domain.Advertisment) CreateAdvertismentOutput
	}

	// Созданное объявление
	CreateAdvertismentOutput struct {
		// Идентификатор
		ID string `json:"id" example:"2765cb06-f750-4d1f-b101-860289786469"`
		// Имя
		Name string `json:"name" example:"Selling the garage"`
		// Идентификаторы категорий
		Categories []string `json:"categories" example:"a9f9ecfe-25b5-4742-901a-21dee231f6cf,d39c79b6-78e0-41fc-939d-e60a64c0251e,d9ff247b-0469-4106-af73-89e7a966e4a4"`
		// Описание
		Description string `json:"description" example:"Very big"`
		// Цена
		Price Price `json:"price"`
		// Ссылка на главное фото
		MainPhotoURL string `json:"mainPhotoURL" example:"http://127.0.0.1/storage/main.jpg"`
		// Ссылки на дополнительные фото
		AdditionalPhotoURLs []string `json:"additionalPhotoURLs" example:"http://127.0.0.1/storage/add1.jpg,http://127.0.0.1/storage/add2.jpg"`
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
	validate := validator.New()
	err := validate.Struct(input)
	if err != nil {
		return CreateAdvertismentOutput{}, fmt.Errorf("error validating create advertisment input: %v", err)
	}

	categories, err := i.catRepo.FindAllByID(ctx, input.Categories)
	if err != nil {
		return CreateAdvertismentOutput{}, fmt.Errorf("can't find categories %s: %v", input.Categories, err)
	}

	price, err := ValidatePrice(input.Price)
	if err != nil {
		return CreateAdvertismentOutput{}, err
	}

	adv := domain.Advertisment{
		ID:                  uuid.NewString(),
		Name:                input.Name,
		Categories:          categories,
		Description:         input.Description,
		Price:               price,
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
