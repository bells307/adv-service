package v1

import (
	"github.com/bells307/adv-service/internal/domain"
)

// TODO: в презентер
type (
	GetAdvertismentResponse struct {
		Name                string       `json:"name"`
		Price               domain.Price `json:"price"`
		MainPhotoURL        string       `json:"mainPhotoURL"`
		Description         *string      `json:"description,omitempty"`
		AdditionalPhotoURLs *[]string    `json:"additionalPhotoURLs,omitempty"`
	}

	GetAdvertismentManyResponse struct {
		Name         string       `json:"name"`
		Price        domain.Price `json:"price"`
		MainPhotoURL string       `json:"mainPhotoURL"`
	}
)
