package model

// Валюта
type Currency string

const (
	RUB Currency = "rub"
	USD Currency = "usd"
)

// Цена
type Price struct {
	Value    float64
	Currency Currency
}
