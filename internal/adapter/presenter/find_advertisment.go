package presenter

import (
	"github.com/bells307/adv-service/internal/domain"
	usecase "github.com/bells307/adv-service/internal/usecase"
)

type findAdvertismentPresenter struct{}

func NewFindAdvertismentPresenter() usecase.FindAdvertismentPresenter {
	return findAdvertismentPresenter{}
}

func (p findAdvertismentPresenter) Output(adv domain.Advertisment) usecase.FindAdvertismentOutput {
	var categoryNames []string
	for _, category := range adv.Categories {
		categoryNames = append(categoryNames, category.Name)
	}

	return usecase.FindAdvertismentOutput{
		ID:                  adv.ID,
		Name:                adv.Name,
		Categories:          categoryNames,
		Description:         adv.Description,
		Price:               adv.Price,
		MainPhotoURL:        adv.MainPhotoURL,
		AdditionalPhotoURLs: adv.AdditionalPhotoURLs,
	}
}
