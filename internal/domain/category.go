package domain

import "context"

type (
	// Репозиторий категорий
	CategoryRepository interface {
		CreateCategory(context.Context, *Category) error
		GetByID(context.Context, string) (*Category, error)
	}

	// Категория объявления
	Category struct {
		// Идентификатор
		ID string `bson:"_id"`
		// Имя
		Name string `bson:"name"`
	}
)
