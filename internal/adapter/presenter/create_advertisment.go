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
	return usecase.CreateAdvertismentOutput{
		ID:                  adv.ID,
		Name:                adv.Name,
		CategoryID:          adv.CategoryID,
		Description:         adv.Description,
		Price:               adv.Price,
		MainPhotoURL:        adv.MainPhotoURL,
		AdditionalPhotoURLs: adv.AdditionalPhotoURLs,
	}
}
