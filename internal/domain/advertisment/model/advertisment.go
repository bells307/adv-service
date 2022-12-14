package model

// Объявление
type Advertisment struct {
	// Идентификатор
	ID string `bson:"_id"`
	// Имя
	Name string `bson:"name"`
	// Описание
	Description string `bson:"description"`
	// Цена
	Price Price `bson:"price"`
	// Ссылка на главное изображение
	MainPhotoURL string `bson:"mainPhotoURL"`
	// Ссылки на дополнительные изображения
	AdditionalPhotoURLs []string `bson:"additionalPhotoURLs"`
}
