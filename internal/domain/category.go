package domain

import (
	"context"
	"errors"
)

var ErrCategoryNotFound = errors.New("category not found")

type (
	// Репозиторий категорий
	CategoryRepository interface {
		// Создать категорию
		CreateCategory(context.Context, Category) error
		// Получить по ID
		FindByID(context.Context, string) (Category, error)
	}

	// Категория объявления
	Category struct {
		// Идентификатор
		ID string `bson:"_id"`
		// Имя
		Name string `bson:"name"`
	}
)
