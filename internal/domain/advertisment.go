package domain

import (
	"context"
	"errors"
	"fmt"
)

// Максимальная длина имени
const MAX_NAME_LENGTH = 200

// Максимальная длина описания для объявления
const MAX_DESC_LENGTH = 1000

// Максимальное количество фото
const MAX_PHOTO_COUNT = 3

// Количество объявлений на странице
const ADV_ON_PAGE_COUNT uint = 10

var (
	ErrAdvNotFound      = errors.New("advertisment not found")
	ErrAdvMaxNameLength = fmt.Errorf(
		"maximum advertisment name length %d exceeded",
		MAX_NAME_LENGTH,
	)
	ErrAdvMaxDescLength = fmt.Errorf(
		"maximum advertisment description length %d exceeded",
		MAX_DESC_LENGTH,
	)
	ErrAdvMaxPhotoCount = fmt.Errorf(
		"maximum advertisment photo count %d exceeded",
		MAX_PHOTO_COUNT,
	)
	ErrAdvPageNegative = errors.New("advertisment page can't be a negative number")
)

type (
	// Интерфейс репозитория объявлений
	AdvertismentRepository interface {
		// Получить объявления
		Find(ctx context.Context, limit uint, offset uint) ([]Advertisment, error)
		// Проверить, есть ли объявления с категорией
		ExistsWithCategory(ctx context.Context, categoryID string) (bool, error)
		// Получить объявление по ID
		FindByID(ctx context.Context, id string) (Advertisment, error)
		// Создать объявление
		Create(ctx context.Context, adv Advertisment) error
		// Удалить объявление
		Delete(ctx context.Context, id string) error
	}

	// Объявление
	Advertisment struct {
		// Идентификатор
		ID string `json:"id" bson:"_id"`
		// Имя
		Name string `json:"name" bson:"name"`
		// Категории
		Categories []Category `json:"categories" bson:"categories"`
		// Описание
		Description string `json:"description" bson:"description"`
		// Цена
		Price Price `json:"price" bson:"price"`
		// Ссылка на главное изображение
		MainPhotoURL string `json:"mainPhotoURL" bson:"mainPhotoURL"`
		// Ссылки на дополнительные изображения
		AdditionalPhotoURLs []string `json:"additionalPhotoURLs" bson:"additionalPhotoURLs"`
	}

	// Краткая информация об объявлении
	AdvertismentSummary struct {
		// Имя
		Name string `json:"name"`
		// Категории
		Categories []string `json:"categories"`
		// Цена
		Price Price `json:"price"`
		// Ссылка на главное изображение
		MainPhotoURL string `json:"mainPhotoURL"`
	}

	// Номер страницы объявлений
	AdvertismentPageNumber uint
)

func (a Advertisment) Validate() error {
	// Валидация полей
	if len(a.Name) > MAX_NAME_LENGTH {
		return ErrAdvMaxNameLength
	}

	if len(a.Description) > MAX_DESC_LENGTH {
		return ErrAdvMaxDescLength
	}

	if len(a.AdditionalPhotoURLs) > MAX_PHOTO_COUNT {
		return ErrAdvMaxPhotoCount
	}

	return nil
}

func (a Advertisment) Summarize() AdvertismentSummary {
	var categoryNames []string
	for _, category := range a.Categories {
		categoryNames = append(categoryNames, category.Name)
	}

	return AdvertismentSummary{
		Name:         a.Name,
		Categories:   categoryNames,
		Price:        a.Price,
		MainPhotoURL: a.MainPhotoURL,
	}
}
