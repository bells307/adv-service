package dto

import (
	"net/url"

	"github.com/bells307/adv-service/internal/model"
)

// Создать объявление
type CreateAdvertisment struct {
	// Имя
	Name string
	// Описание
	Description string
	// Цена
	Price model.Price
	// Ссылка на изображение
	ImageURLs []url.URL
}
