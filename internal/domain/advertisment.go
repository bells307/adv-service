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

var (
	ErrAdvNotFound      = errors.New("advertisment not found")
	ErrAdvMaxNameLength = fmt.Errorf("maximum name length %d exceeded", MAX_NAME_LENGTH)
	ErrAdvMaxDescLength = fmt.Errorf("maximum description length %d exceeded", MAX_DESC_LENGTH)
	ErrAdvMaxPhotoCount = fmt.Errorf("maximum advertisment photo count %d exceeded", MAX_PHOTO_COUNT)
)

type (
	// Интерфейс репозитория объявлений
	AdvertismentRepository interface {
		// Получить объявления
		Get(ctx context.Context, limit int64, offset int64) ([]*Advertisment, error)
		// Получить объявление
		GetOne(ctx context.Context, id string) (*Advertisment, error)
		// Создать объявление
		Create(ctx context.Context, adv *Advertisment) error
	}

	// Объявление
	Advertisment struct {
		// Идентификатор
		ID string `bson:"_id"`
		// Имя
		Name string `bson:"name"`
		// Категория
		CategoryID string `bson:"categoryID"`
		// Описание
		Description string `bson:"description"`
		// Цена
		Price Price `bson:"price"`
		// Ссылка на главное изображение
		MainPhotoURL string `bson:"mainPhotoURL"`
		// Ссылки на дополнительные изображения
		AdditionalPhotoURLs []string `bson:"additionalPhotoURLs"`
	}
)
