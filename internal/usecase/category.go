package usecase

import (
	"context"

	"github.com/bells307/adv-service/internal/domain"
	"github.com/google/uuid"
)

type (
	// Юзкейс работы с категориями
	CategoryUsecase struct {
		repo domain.CategoryRepository
	}

	// Создать категорию
	CreateCategory struct {
		// Имя категории
		Name string `json:"name"`
	}
)

func NewCategoryUsecase(repo domain.CategoryRepository) *CategoryUsecase {
	return &CategoryUsecase{repo}
}

// Создать категорию
func (s *CategoryUsecase) CreateCategory(ctx context.Context, createCategory *CreateCategory) (*domain.Category, error) {
	cat := &domain.Category{
		ID:   uuid.NewString(),
		Name: createCategory.Name,
	}

	if err := s.repo.CreateCategory(ctx, cat); err != nil {
		return nil, err
	}

	return cat, nil
}

// Получить категорию по ID
func (s *CategoryUsecase) GetByID(ctx context.Context, id string) (*domain.Category, error) {
	return s.repo.GetByID(ctx, id)
}
