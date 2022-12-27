package usecase

import (
	"context"
	"fmt"

	"github.com/bells307/adv-service/internal/domain"
	"github.com/go-playground/validator/v10"
)

type (
	// Удаление объявления
	DeleteAdvertismentUseCase interface {
		Execute(context.Context, DeleteAdvertismentInput) error
	}

	// Удалить объявление
	DeleteAdvertismentInput struct {
		ID string `json:"id" validate:"required" example:"2765cb06-f750-4d1f-b101-860289786469"`
	}

	DeleteAdvertismentInteractor struct {
		advRepo domain.AdvertismentRepository
	}
)

func NewDeleteAdvertismentInteractor(
	advRepo domain.AdvertismentRepository,
) DeleteAdvertismentUseCase {
	return DeleteAdvertismentInteractor{
		advRepo,
	}
}

func (i DeleteAdvertismentInteractor) Execute(ctx context.Context, input DeleteAdvertismentInput) error {
	validate := validator.New()
	err := validate.Struct(input)
	if err != nil {
		return fmt.Errorf("error validating delete advertisment input: %v", err)
	}

	return i.advRepo.Delete(ctx, input.ID)
}
