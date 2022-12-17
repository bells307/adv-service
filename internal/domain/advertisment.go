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
		// Получить объявление по ID
		FindByID(ctx context.Context, id string) (Advertisment, error)
		// Создать объявление
		Create(ctx context.Context, adv Advertisment) error
	}

	// Объявление
	Advertisment struct {
		// Идентификатор
		ID string `bson:"_id"`
		// Имя
		Name string `bson:"name"`
		// Категория
		Category Category `bson:"category"`
		// Описание
		Description string `bson:"description"`
		// Цена
		Price Price `bson:"price"`
		// Ссылка на главное изображение
		MainPhotoURL string `bson:"mainPhotoURL"`
		// Ссылки на дополнительные изображения
		AdditionalPhotoURLs []string `bson:"additionalPhotoURLs"`
	}

	// Краткая информация об объявлении
	AdvertismentSummary struct {
		// Имя
		Name string
		// Категория
		Category Category
		// Цена
		Price Price
		// Ссылка на главное изображение
		MainPhotoURL string
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
	return AdvertismentSummary{
		Name:         a.Name,
		Category:     a.Category,
		Price:        a.Price,
		MainPhotoURL: a.MainPhotoURL,
	}
}
