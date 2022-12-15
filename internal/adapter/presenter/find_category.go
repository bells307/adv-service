package presenter

import (
	"github.com/bells307/adv-service/internal/domain"
	usecase "github.com/bells307/adv-service/internal/usecase"
)

type findCategoryPresenter struct{}

func NewFindCategoryPresenter() usecase.FindCategoryPresenter {
	return findCategoryPresenter{}
}

func (p findCategoryPresenter) Output(adv domain.Category) usecase.FindCategoryOutput {
	return usecase.FindCategoryOutput{
		ID:   adv.ID,
		Name: adv.Name,
	}
}
