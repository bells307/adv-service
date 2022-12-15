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
		ID string
		// Имя
		Name string
		// Категория
		CategoryID string
		// Описание
		Description string
		// Цена
		Price Price
		// Ссылка на главное изображение
		MainPhotoURL string
		// Ссылки на дополнительные изображения
		AdditionalPhotoURLs []string
	}

	AdvertismentSummary struct {
		// Имя
		Name string
		// Цена
		Price Price
		// Ссылка на главное изображение
		MainPhotoURL string
	}

	// Номер страницы объявлений
	AdvertismentPageNumber uint
)

func (a Advertisment) Summary() AdvertismentSummary {
	return AdvertismentSummary{
		Name:         a.Name,
		Price:        a.Price,
		MainPhotoURL: a.MainPhotoURL,
	}
}
