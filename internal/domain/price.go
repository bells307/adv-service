package domain

type (
	// Валюта
	Currency string

	// Цена
	Price struct {
		// Значение
		Value float64 `json:"value" bson:"value" example:"1000"`
		// Валюта
		Currency Currency `json:"currency" bson:"currency" example:"rub"`
	}
)

const (
	// Рубль
	RUB Currency = "rub"
	// Доллар
	USD Currency = "usd"
)
