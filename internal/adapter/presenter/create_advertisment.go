package presenter

import (
	"github.com/bells307/adv-service/internal/domain"
	usecase "github.com/bells307/adv-service/internal/usecase"
)

type createAdvertismentPresenter struct{}

func NewCreateAdvertismentPresenter() usecase.CreateAdvertismentPresenter {
	return createAdvertismentPresenter{}
}

func (p createAdvertismentPresenter) Output(adv domain.Advertisment) usecase.CreateAdvertismentOutput {
	var categories []string
	for _, category := range adv.Categories {
		categories = append(categories, category.ID)
	}

	return usecase.CreateAdvertismentOutput{
		ID:          adv.ID,
		Name:        adv.Name,
		Categories:  categories,
		Description: adv.Description,
		Price: usecase.Price{
			Value:    adv.Price.Value,
			Currency: string(adv.Price.Currency),
		},
		MainPhotoURL:        adv.MainPhotoURL,
		AdditionalPhotoURLs: adv.AdditionalPhotoURLs,
	}
}
