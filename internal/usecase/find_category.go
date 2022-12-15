package usecase

import (
	"context"

	"github.com/bells307/adv-service/internal/domain"
)

type (
	// Получение объявления
	FindCategoryUseCase interface {
		Execute(context.Context, FindCategoryInput) (FindCategoryOutput, error)
	}

	// Найти категорию
	FindCategoryInput struct {
		// Идентификатор категории
		ID string `json:"id"`
	}

	// Порт выхода презентера
	FindCategoryPresenter interface {
		Output(domain.Category) FindCategoryOutput
	}

	// Найденная категория
	FindCategoryOutput struct {
		// Идентификатор
		ID string `json:"id"`
		// Имя
		Name string `json:"name"`
	}

	findCategoryInteractor struct {
		repo      domain.CategoryRepository
		presenter FindCategoryPresenter
	}
)

func NewFindCategoryInteractor(
	repo domain.CategoryRepository,
	presenter FindCategoryPresenter,
) findCategoryInteractor {
	return findCategoryInteractor{
		repo,
		presenter,
	}
}

func (i findCategoryInteractor) Execute(ctx context.Context, input FindCategoryInput) (FindCategoryOutput, error) {
	cat, err := i.repo.FindByID(ctx, input.ID)
	if err != nil {
		return i.presenter.Output(domain.Category{}), err
	}

	return i.presenter.Output(cat), nil
}
