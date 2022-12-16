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
			Name         string       `json:"name"`
			Category     string       `json:"category"`
			Price        domain.Price `json:"price"`
			MainPhotoURL string       `json:"mainPhotoURL"`
		}{
			Name:         sum.Name,
			Category:     sum.Category.Name,
			Price:        sum.Price,
			MainPhotoURL: sum.MainPhotoURL,
		}

		out = append(out, val)
	}

	return out
}
