package usecase

import (
	"context"
	"fmt"

	"github.com/bells307/adv-service/internal/domain"
	"github.com/go-playground/validator/v10"
)

type (
	// Получение объявления
	FindCategoryUseCase interface {
		Execute(context.Context, FindCategoryInput) (FindCategoryOutput, error)
	}

	// Найти категорию
	FindCategoryInput struct {
		// Идентификатор категории
		ID string `json:"id" validate:"required" example:"e15a4f3f-1549-466e-990a-4b44d10bd3aa"`
	}

	// Порт выхода презентера
	FindCategoryPresenter interface {
		Output(domain.Category) FindCategoryOutput
	}

	// Найденная категория
	FindCategoryOutput struct {
		// Идентификатор
		ID string `json:"id" example:"e15a4f3f-1549-466e-990a-4b44d10bd3aa"`
		// Имя
		Name string `json:"name" example:"car"`
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
	validate := validator.New()
	err := validate.Struct(input)
	if err != nil {
		return FindCategoryOutput{}, fmt.Errorf("error validating find category input: %v", err)
	}

	cat, err := i.repo.FindByID(ctx, input.ID)
	if err != nil {
		return i.presenter.Output(domain.Category{}), err
	}

	return i.presenter.Output(cat), nil
}
