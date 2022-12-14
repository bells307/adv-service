package dto

import (
	"github.com/bells307/adv-service/internal/domain/advertisment/model"
)

// Создать объявление
type CreateAdvertisment struct {
	// Имя
	Name string
	// Описание
	Description string
	// Цена
	Price model.Price
	// Ссылка на главное изображение
	MainPhotoURL string
	// Ссылки на дополнительные изображения
	AdditionalPhotoURLs []string
}

// Получить объявления
type GetAdvertisments struct {
	// Максимальное количество объявлений в результате
	Limit int64
	// Смещение от начала
	Offset int64
}
