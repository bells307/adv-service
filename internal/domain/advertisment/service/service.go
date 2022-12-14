package service

import (
	"context"
	"fmt"

	"github.com/bells307/adv-service/internal/domain/advertisment/model"
	"github.com/bells307/adv-service/internal/domain/advertisment/service/constant"
	dto "github.com/bells307/adv-service/internal/domain/advertisment/service/dto"
	service_err "github.com/bells307/adv-service/internal/domain/advertisment/service/error"
	"github.com/google/uuid"
)

// Сервис работы с объявлениями
type AdvertismentService struct {
	repo advertismentRepository
}

// Интерфейс репозитория объявлений
type advertismentRepository interface {
	// Получить объявления
	Get(ctx context.Context, limit int64, offset int64) ([]*model.Advertisment, error)
	// Получить объявление
	GetOne(ctx context.Context, id string) (*model.Advertisment, error)
	// Создать объявление
	Create(ctx context.Context, adv *model.Advertisment) error
}

func NewAdvertismentService(repo advertismentRepository) *AdvertismentService {
	return &AdvertismentService{repo}
}

// Получить объявления
func (s *AdvertismentService) Get(ctx context.Context, getAdv *dto.GetAdvertisments) ([]*model.Advertisment, error) {
	return s.repo.Get(ctx, getAdv.Limit, getAdv.Offset)
}

// Получить объявление
func (s *AdvertismentService) GetOne(ctx context.Context, id string) (*model.Advertisment, error) {
	res, err := s.repo.GetOne(ctx, id)
	if err != nil {
		return nil, err
	}

	if res == nil {
		return nil, service_err.ErrNotFound
	}

	return res, nil
}

// Создать объявление
func (s *AdvertismentService) Create(ctx context.Context, createAdv *dto.CreateAdvertisment) (adv *model.Advertisment, err error) {
	// Валидация полей
	if len(createAdv.Name) > constant.MAX_NAME_LENGTH {
		return adv, service_err.ErrMaxNameLength
	}

	if len(createAdv.Description) > constant.MAX_DESC_LENGTH {
		return adv, service_err.ErrMaxDescLength
	}

	if len(createAdv.AdditionalPhotoURLs) > constant.MAX_PHOTO_COUNT {
		return adv, service_err.ErrMaxPhotoCount
	}

	adv = &model.Advertisment{
		ID:                  uuid.NewString(),
		Name:                createAdv.Name,
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
