package domain

import "fmt"

type (
	// Валюта
	Currency string

	// Цена
	Price struct {
		Value    float64  `json:"value" bson:"value"`
		Currency Currency `json:"currency" bson:"currency"`
	}
)

const (
	RUB Currency = "rub"
	USD Currency = "usd"
)

func CurrencyFromString(str string) (Currency, error) {
	if str == "rub" {
		return RUB, nil
	} else if str == "usd" {
		return USD, nil
	} else {
		return Currency(""), fmt.Errorf("currency %s not found", str)
	}
}
