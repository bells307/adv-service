package usecase

import (
	"context"

	"github.com/bells307/adv-service/internal/domain"
	"github.com/google/uuid"
)

type (
	// Создание категории
	CreateCategoryUseCase interface {
		Execute(context.Context, CreateCategoryInput) (CreateCategoryOutput, error)
	}

	// Создать категорию
	CreateCategoryInput struct {
		Name string `json:"name"`
	}

	// Порт выхода презентера
	CreateCategoryPresenter interface {
		Output(domain.Category) CreateCategoryOutput
	}

	// Созданная категория
	CreateCategoryOutput struct {
		// Идентификатор
		ID string `json:"id"`
		// Имя
		Name string `json:"name"`
	}

	createCategoryInteractor struct {
		repo      domain.CategoryRepository
		presenter CreateCategoryPresenter
	}
)

func NewCreateCategoryInteractor(
	repo domain.CategoryRepository,
	presenter CreateCategoryPresenter,
) CreateCategoryUseCase {
	return createCategoryInteractor{
		repo,
		presenter,
	}
}

func (i createCategoryInteractor) Execute(ctx context.Context, input CreateCategoryInput) (CreateCategoryOutput, error) {
	cat := domain.Category{
		ID:   uuid.NewString(),
		Name: input.Name,
	}

	err := i.repo.CreateCategory(ctx, cat)
	if err != nil {
		return i.presenter.Output(domain.Category{}), err
	}

	return i.presenter.Output(cat), nil
}
