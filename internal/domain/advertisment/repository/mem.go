package repository

import (
	"context"
	"sync"

	"github.com/bells307/adv-service/internal/domain/advertisment/model"
)

type advertismentMemoryRepository struct {
	mtx           sync.Mutex
	advertisments []*model.Advertisment
}

func NewAdvertismentMemoryRepository() *advertismentMemoryRepository {
	return &advertismentMemoryRepository{}
}

// Получить объявления
func (r *advertismentMemoryRepository) Get(ctx context.Context, limit int64, offset int64) ([]*model.Advertisment, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	return r.advertisments, nil
}

// Получить объявление
func (r *advertismentMemoryRepository) GetOne(ctx context.Context, id string) (*model.Advertisment, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	for i := 0; i < len(r.advertisments); i++ {
		if r.advertisments[i].ID == id {
			return r.advertisments[i], nil
		}
	}

	return nil, nil
}

// Создать объявление
func (r *advertismentMemoryRepository) Create(ctx context.Context, adv *model.Advertisment) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	r.advertisments = append(r.advertisments, adv)
	return nil
}
