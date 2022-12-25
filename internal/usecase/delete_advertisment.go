package usecase

import (
	"context"

	"github.com/bells307/adv-service/internal/domain"
)

type (
	// Удаление объявления
	DeleteAdvertismentUseCase interface {
		Execute(context.Context, DeleteAdvertismentInput) error
	}

	// Удалить объявление
	DeleteAdvertismentInput struct {
		ID string `json:"id"`
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
	return i.advRepo.Delete(ctx, input.ID)
}
