package usecase

import (
	"context"

	"github.com/bells307/adv-service/internal/domain"
)

type (
	// Удаление категории
	DeleteCategoryUseCase interface {
		Execute(context.Context, DeleteCategoryInput) error
	}

	// Удалить категорию
	DeleteCategoryInput struct {
		ID string `json:"id"`
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
	panic("not implemented")
	// Сначала ищем объявления, которые используют удаляемую категорию
	// exists, err := i.advRepo.ExistsWithCategory(ctx, input.ID)
	// if err != nil {
	// 	return fmt.Errorf("error deleting category %s: %v", input.ID, err)
	// }

	// if exists {
	// 	return domain.ErrDeletingCategoryWithAdvertisment
	// }

	// return i.catRepo.Delete(ctx, input.ID)
}
