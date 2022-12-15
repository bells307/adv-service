package presenter

import (
	"github.com/bells307/adv-service/internal/domain"
	usecase "github.com/bells307/adv-service/internal/usecase"
)

type findAllAdvertismentSummaryByPagePresenter struct{}

func NewFindAllAdvertismentSummaryByPage() usecase.FindAllAdvertismentSummaryByPagePresenter {
	return findAllAdvertismentSummaryByPagePresenter{}
}

func (p findAllAdvertismentSummaryByPagePresenter) Output(summary []domain.AdvertismentSummary) usecase.FindAllAdvertismentSummaryByPageOutput {
	out := make(usecase.FindAllAdvertismentSummaryByPageOutput, 0)

	for _, sum := range summary {
		val := struct {
			Name         string       `json:"name"`
			Price        domain.Price `json:"price"`
			MainPhotoURL string       `json:"mainPhotoURL"`
		}{
			Name:         sum.Name,
			Price:        sum.Price,
			MainPhotoURL: sum.MainPhotoURL,
		}

		out = append(out, val)
	}

	return out
}
