package domain

import "fmt"

type (
	// Валюта
	Currency string

	// Цена
	Price struct {
		// Значение
		Value float64 `json:"value" bson:"value"`
		// Валюта
		Currency Currency `json:"currency" bson:"currency"`
	}
)

const (
	// Рубль
	RUB Currency = "rub"
	// Доллар
	USD Currency = "usd"
)

// Преобразование string => Currency
func CurrencyFromString(str string) (Currency, error) {
	if str == "rub" {
		return RUB, nil
	} else if str == "usd" {
		return USD, nil
	} else {
		return Currency(""), fmt.Errorf("currency %s not found", str)
	}
}
