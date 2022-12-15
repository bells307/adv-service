package usecase

import (
	"context"
	"fmt"

	"github.com/bells307/adv-service/internal/domain"
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
		Name string `json:"name"`
		// ID категории
		CategoryID string `json:"categoryID"`
		// Описание
		Description string `json:"description"`
		// Цена
		Price domain.Price `json:"price"`
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
		// ID категории
		CategoryID string `json:"categoryID"`
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
		repo      domain.AdvertismentRepository
		presenter CreateAdvertismentPresenter
	}
)

func NewCreateAdvertismentInteractor(
	repo domain.AdvertismentRepository,
	presenter CreateAdvertismentPresenter,
) CreateAdvertismentUseCase {
	return createAdvertismentInteractor{
		repo,
		presenter,
	}
}

func (i createAdvertismentInteractor) Execute(ctx context.Context, input CreateAdvertismentInput) (CreateAdvertismentOutput, error) {
	// Валидация полей
	if len(input.Name) > domain.MAX_NAME_LENGTH {
		return CreateAdvertismentOutput{}, domain.ErrAdvMaxNameLength
	}

	if len(input.Description) > domain.MAX_DESC_LENGTH {
		return CreateAdvertismentOutput{}, domain.ErrAdvMaxDescLength
	}

	if len(input.AdditionalPhotoURLs) > domain.MAX_PHOTO_COUNT {
		return CreateAdvertismentOutput{}, domain.ErrAdvMaxPhotoCount
	}

	adv := domain.Advertisment{
		ID:                  uuid.NewString(),
		Name:                input.Name,
		CategoryID:          input.CategoryID,
		Description:         input.Description,
		Price:               input.Price,
		MainPhotoURL:        input.MainPhotoURL,
		AdditionalPhotoURLs: input.AdditionalPhotoURLs,
	}

	if err := i.repo.Create(ctx, adv); err != nil {
		return CreateAdvertismentOutput{}, fmt.Errorf("failed to create advertisment: %v", err)
	}

	return i.presenter.Output(adv), nil
}
