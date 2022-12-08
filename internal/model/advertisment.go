package model

import "net/url"

// Объявление
type Advertisment struct {
	// Идентификатор
	ID string
	// Имя
	Name string
	// Описание
	Description string
	// Цена
	Price Price
	// Ссылки на изображения
	ImageURLs []url.URL
}
