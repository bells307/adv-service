package usecase

import (
	"context"
	"fmt"

	"github.com/bells307/adv-service/internal/domain"
	"github.com/go-playground/validator/v10"
)

type (
	// Удаление категории
	DeleteCategoryUseCase interface {
		Execute(context.Context, DeleteCategoryInput) error
	}

	// Удалить категорию
	DeleteCategoryInput struct {
		ID string `json:"id" validate:"required" example:"e15a4f3f-1549-466e-990a-4b44d10bd3aa"`
	}

	DeleteCategoryInteractor struct {
		advRepo domain.AdvertismentRepository
		catRepo domain.CategoryRepository
	}
)

func NewDeleteCategoryInteractor(
	advRepo domain.AdvertismentRepository,
	catRepo domain.CategoryRepository,
) DeleteCategoryUseCase {
	return DeleteCategoryInteractor{advRepo, catRepo}
}

func (i DeleteCategoryInteractor) Execute(ctx context.Context, input DeleteCategoryInput) error {
	validate := validator.New()
	err := validate.Struct(input)
	if err != nil {
		return fmt.Errorf("error validating delete category input: %v", err)
	}

	// Сначала ищем объявления, которые используют удаляемую категорию
	count, err := i.advRepo.CountByCategory(ctx, input.ID)
	if err != nil {
		return fmt.Errorf("error deleting category %s: %v", input.ID, err)
	}

	if count > 0 {
		return domain.ErrDeletingCategoryWithAdvertisment
	}

	return i.catRepo.Delete(ctx, input.ID)
}
