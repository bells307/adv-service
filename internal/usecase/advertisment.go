package usecase

import (
	"context"
	"fmt"

	"github.com/bells307/adv-service/internal/domain"
	"github.com/google/uuid"
)

type (
	// Юзкейс работы с объявлениями
	AdvertismentUsecase struct {
		repo domain.AdvertismentRepository
	}

	// Создать объявление
	CreateAdvertisment struct {
		// Имя
		Name string
		// Категория
		CategoryID string
		// Описание
		Description string
		// Цена
		Price domain.Price
		// Ссылка на главное изображение
		MainPhotoURL string
		// Ссылки на дополнительные изображения
		AdditionalPhotoURLs []string
	}

	// Получить объявления
	GetAdvertisments struct {
		// Максимальное количество объявлений в результате
		Limit int64
		// Смещение от начала
		Offset int64
	}
)

func NewAdvertismentUsecase(repo domain.AdvertismentRepository) *AdvertismentUsecase {
	return &AdvertismentUsecase{repo}
}

// Получить объявления
func (s *AdvertismentUsecase) Get(ctx context.Context, getAdv *GetAdvertisments) ([]*domain.Advertisment, error) {
	return s.repo.Get(ctx, getAdv.Limit, getAdv.Offset)
}

// Получить объявление
func (s *AdvertismentUsecase) GetOne(ctx context.Context, id string) (*domain.Advertisment, error) {
	res, err := s.repo.GetOne(ctx, id)
	if err != nil {
		return nil, err
	}

	if res == nil {
		return nil, domain.ErrAdvNotFound
	}

	return res, nil
}

// Создать объявление
func (s *AdvertismentUsecase) Create(ctx context.Context, createAdv *CreateAdvertisment) (adv *domain.Advertisment, err error) {
	// Валидация полей
	if len(createAdv.Name) > domain.MAX_NAME_LENGTH {
		return adv, domain.ErrAdvMaxNameLength
	}

	if len(createAdv.Description) > domain.MAX_DESC_LENGTH {
		return adv, domain.ErrAdvMaxDescLength
	}

	if len(createAdv.AdditionalPhotoURLs) > domain.MAX_PHOTO_COUNT {
		return adv, domain.ErrAdvMaxPhotoCount
	}

	adv = &domain.Advertisment{
		ID:                  uuid.NewString(),
		Name:                createAdv.Name,
		CategoryID:          createAdv.CategoryID,
		Description:         createAdv.Description,
		Price:               createAdv.Price,
		MainPhotoURL:        createAdv.MainPhotoURL,
		AdditionalPhotoURLs: createAdv.AdditionalPhotoURLs,
	}

	if err := s.repo.Create(ctx, adv); err != nil {
		return adv, fmt.Errorf("failed to create advertisment: %v", err)
	}

	return adv, nil
}
