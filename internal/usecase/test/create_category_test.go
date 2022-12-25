package test

import (
	"context"
	"testing"

	"github.com/bells307/adv-service/internal/domain"
	"github.com/bells307/adv-service/internal/usecase"
	"github.com/bells307/adv-service/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Тест создания категории
func TestCreateCategoty(t *testing.T) {
	ctx := context.Background()

	cat := domain.Category{
		ID:   uuid.NewString(),
		Name: "some category",
	}

	in := usecase.CreateCategoryInput{
		Name: cat.Name,
	}

	expOut := usecase.CreateCategoryOutput{
		ID:   mock.Anything,
		Name: cat.Name,
	}

	catRepoMock := mocks.NewCategoryRepository(t)
	catRepoMock.On("Create", ctx, mock.MatchedBy(func(c domain.Category) bool {
		return c.Name == cat.Name
	})).Return(nil)

	createCatPresMock := mocks.NewCreateCategoryPresenter(t)
	createCatPresMock.On("Output", mock.MatchedBy(func(c domain.Category) bool {
		return c.Name == cat.Name
	})).Return(expOut)

	uc := usecase.NewCreateCategoryInteractor(catRepoMock, createCatPresMock)

	out, err := uc.Execute(ctx, in)

	assert.Nil(t, err)
	assert.Equal(t, out, expOut)
}
