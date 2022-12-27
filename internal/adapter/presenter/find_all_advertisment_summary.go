package presenter

import (
	"github.com/bells307/adv-service/internal/domain"
	usecase "github.com/bells307/adv-service/internal/usecase"
)

type findAllAdvertismentSummaryPresenter struct{}

func NewFindAllAdvertismentSummaryPresenter() usecase.FindAllAdvertismentSummaryPresenter {
	return findAllAdvertismentSummaryPresenter{}
}

func (p findAllAdvertismentSummaryPresenter) Output(summary []domain.AdvertismentSummary) usecase.FindAllAdvertismentSummaryOutput {
	out := make(usecase.FindAllAdvertismentSummaryOutput, 0)

	for _, sum := range summary {
		val := struct {
			Name         string       `json:"name" example:"Selling the garage"`
			Categories   []string     `json:"categories" example:"real estate,auto,land"`
			Price        domain.Price `json:"price"`
			MainPhotoURL string       `json:"mainPhotoURL" example:"http://127.0.0.1/storage/main.jpg"`
		}{
			Name:         sum.Name,
			Categories:   sum.Categories,
			Price:        sum.Price,
			MainPhotoURL: sum.MainPhotoURL,
		}

		out = append(out, val)
	}

	return out
}
