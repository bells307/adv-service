package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/bells307/adv-service/internal/dto"
	"github.com/bells307/adv-service/internal/model"
	"github.com/google/uuid"
)

var ErrNotFound = errors.New("advertisment not found")

// Сервис работы с объявлениями
type advertismentService struct {
	repo advertismentRepository
}

// Интерфейс репозитория объявлений
type advertismentRepository interface {
	// Получить объявления
	Get(ctx context.Context) ([]model.Advertisment, error)
	// Получить объявление
	GetOne(ctx context.Context, id string) (model.Advertisment, error)
	// Создать объявление
	Create(ctx context.Context, adv model.Advertisment) error
}

func NewAdvertismentService(repo advertismentRepository) *advertismentService {
	return &advertismentService{repo}
}

// Получить объявления
func (s *advertismentService) Get(ctx context.Context) ([]model.Advertisment, error) {
	return s.repo.Get(ctx)
}

// Получить объявление
func (s *advertismentService) GetOne(ctx context.Context, id string) (model.Advertisment, error) {
	return s.repo.GetOne(ctx, id)
}

// Создать объявление
func (s *advertismentService) Create(ctx context.Context, createAdv dto.CreateAdvertisment) (adv model.Advertisment, err error) {
	adv = model.Advertisment{
		ID:          uuid.NewString(),
		Name:        createAdv.Name,
		Description: createAdv.Description,
		Price:       createAdv.Price,
		ImageURLs:   createAdv.ImageURLs,
	}

	if err := s.repo.Create(ctx, adv); err != nil {
		return adv, fmt.Errorf("failed to create advertisment: %v", err)
	}

	return adv, nil
}
