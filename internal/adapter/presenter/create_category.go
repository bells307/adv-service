package presenter

import (
	"github.com/bells307/adv-service/internal/domain"
	usecase "github.com/bells307/adv-service/internal/usecase"
)

type createCategoryPresenter struct{}

func NewCreateCategoryPresenter() usecase.CreateCategoryPresenter {
	return createCategoryPresenter{}
}

func (p createCategoryPresenter) Output(adv domain.Category) usecase.CreateCategoryOutput {
	return usecase.CreateCategoryOutput{
		ID:   adv.ID,
		Name: adv.Name,
	}
}
