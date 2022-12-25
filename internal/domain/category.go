package domain

import (
	"context"
	"errors"
)

var (
	ErrCategoryNotFound                 = errors.New("category not found")
	ErrDeletingCategoryWithAdvertisment = errors.New("can't delete category with existing advertisments")
)

type (
	// Репозиторий категорий
	CategoryRepository interface {
		// Создать категорию
		Create(context.Context, Category) error
		// Удалить категорию
		Delete(context.Context, string) error
		// Получить по ID
		FindByID(context.Context, string) (Category, error)
		// Получить все
		FindAllByID(ctx context.Context, ids []string) ([]Category, error)
	}

	// Категория объявления
	Category struct {
		// Идентификатор
		ID string `json:"id" bson:"_id"`
		// Имя
		Name string `json:"name" bson:"name"`
	}
)
