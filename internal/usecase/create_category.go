package usecase

import (
	"context"
	"fmt"

	"github.com/bells307/adv-service/internal/domain"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type (
	// Создание категории
	CreateCategoryUseCase interface {
		Execute(context.Context, CreateCategoryInput) (CreateCategoryOutput, error)
	}

	// Создать категорию
	CreateCategoryInput struct {
		Name string `json:"name" validate:"required" example:"car"`
	}

	// Порт выхода презентера
	CreateCategoryPresenter interface {
		Output(domain.Category) CreateCategoryOutput
	}

	// Созданная категория
	CreateCategoryOutput struct {
		// Идентификатор
		ID string `json:"id" example:"e15a4f3f-1549-466e-990a-4b44d10bd3aa"`
		// Имя
		Name string `json:"name" example:"car"`
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
	validate := validator.New()
	err := validate.Struct(input)
	if err != nil {
		return CreateCategoryOutput{}, fmt.Errorf("error validating create category input: %v", err)
	}

	cat := domain.Category{
		ID:   uuid.NewString(),
		Name: input.Name,
	}

	err = i.repo.Create(ctx, cat)
	if err != nil {
		return i.presenter.Output(domain.Category{}), err
	}

	return i.presenter.Output(cat), nil
}
