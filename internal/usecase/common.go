package usecase

import (
	"fmt"
	"github.com/bells307/adv-service/internal/domain"
)

// Цена
type Price struct {
	// Значение
	Value float64 `json:"value" bson:"value" example:"1000"`
	// Валюта
	Currency string `json:"currency" bson:"currency" example:"rub"`
}

func ValidatePrice(p Price) (domain.Price, error) {
	var currency domain.Currency
	var err error
	if p.Currency == "rub" {
		currency = domain.RUB
	} else if p.Currency == "usd" {
		currency = domain.USD
	} else {
		err = fmt.Errorf("currency %s not found", p.Currency)
	}

	if err != nil {
		return domain.Price{}, err
	}

	return domain.Price{
		Value:    p.Value,
		Currency: currency,
	}, nil
}
