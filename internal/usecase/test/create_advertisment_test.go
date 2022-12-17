package test

import (
	"context"
	"fmt"
	"math/rand"
	"testing"

	"github.com/bells307/adv-service/internal/domain"
	"github.com/bells307/adv-service/internal/usecase"
	"github.com/bells307/adv-service/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Тест создания объявления
func TestCreateAdvertisment(t *testing.T) {
	ctx := context.Background()

	cat := domain.Category{
		ID:   uuid.NewString(),
		Name: "some category",
	}

	photoUrls := generatePhotoUrls(domain.MAX_PHOTO_COUNT)
	in := usecase.CreateAdvertismentInput{
		Name:        "myadv",
		CategoryID:  cat.ID,
		Description: "adv description",
		Price: domain.Price{
			Value:    1000,
			Currency: "rub",
		},
		MainPhotoURL:        photoUrls[0],
		AdditionalPhotoURLs: photoUrls[1:],
	}

	advRepoMock := mocks.NewAdvertismentRepository(t)
	advRepoMock.On("Create", ctx, mock.MatchedBy(matchInput(in))).Return(nil)

	catRepoMock := mocks.NewCategoryRepository(t)
	catRepoMock.On("FindByID", ctx, cat.ID).Return(cat, nil)

	createAdvPresMock := mocks.NewCreateAdvertismentPresenter(t)
	createAdvPresMock.On("Output", mock.MatchedBy(matchInput(in))).Return(usecase.CreateAdvertismentOutput{
		ID:                  mock.Anything,
		Name:                in.Name,
		Category:            cat.Name,
		Description:         in.Description,
		Price:               in.Price,
		MainPhotoURL:        in.MainPhotoURL,
		AdditionalPhotoURLs: in.AdditionalPhotoURLs,
	})

	uc := usecase.NewCreateAdvertismentInteractor(advRepoMock, catRepoMock, createAdvPresMock)
	out, err := uc.Execute(ctx, in)

	advRepoMock.AssertExpectations(t)
	catRepoMock.AssertExpectations(t)
	createAdvPresMock.AssertExpectations(t)

	assert.Nil(t, err)
	assert.Equal(t, out, usecase.CreateAdvertismentOutput{
		ID:                  mock.Anything,
		Name:                in.Name,
		Category:            cat.Name,
		Description:         in.Description,
		Price:               in.Price,
		MainPhotoURL:        in.MainPhotoURL,
		AdditionalPhotoURLs: in.AdditionalPhotoURLs,
	})
}

// Тест проверки валидности объявления с превышенным количеством дополнительных фото
func TestCreateAdvertismentWithMaxAdditionalPhotoCountExceeded(t *testing.T) {
	ctx := context.Background()

	cat := domain.Category{
		ID:   uuid.NewString(),
		Name: "some category",
	}

	photoUrls := generatePhotoUrls(domain.MAX_PHOTO_COUNT + 2)
	in := usecase.CreateAdvertismentInput{
		Name:        "myadv",
		CategoryID:  cat.ID,
		Description: "adv description",
		Price: domain.Price{
			Value:    1000,
			Currency: "rub",
		},
		MainPhotoURL:        photoUrls[0],
		AdditionalPhotoURLs: photoUrls[1:],
	}

	advRepoMock := mocks.NewAdvertismentRepository(t)

	catRepoMock := mocks.NewCategoryRepository(t)
	catRepoMock.On("FindByID", ctx, cat.ID).Return(cat, nil)

	createAdvPresMock := mocks.NewCreateAdvertismentPresenter(t)

	uc := usecase.NewCreateAdvertismentInteractor(advRepoMock, catRepoMock, createAdvPresMock)
	_, err := uc.Execute(ctx, in)

	catRepoMock.AssertExpectations(t)

	if assert.Error(t, err) {
		assert.Equal(t, domain.ErrAdvMaxPhotoCount, err)
	}
}

// Тест проверки валидности объявления с превышенной длиной имени
func TestCreateAdvertismentWithMaxNameLengthExceeded(t *testing.T) {
	ctx := context.Background()

	cat := domain.Category{
		ID:   uuid.NewString(),
		Name: "some category",
	}

	photoUrls := generatePhotoUrls(domain.MAX_PHOTO_COUNT)
	in := usecase.CreateAdvertismentInput{
		Name:        randSeq(domain.MAX_NAME_LENGTH + 1),
		CategoryID:  cat.ID,
		Description: "adv description",
		Price: domain.Price{
			Value:    1000,
			Currency: "rub",
		},
		MainPhotoURL:        photoUrls[0],
		AdditionalPhotoURLs: photoUrls[1:],
	}

	advRepoMock := mocks.NewAdvertismentRepository(t)

	catRepoMock := mocks.NewCategoryRepository(t)
	catRepoMock.On("FindByID", ctx, cat.ID).Return(cat, nil)

	createAdvPresMock := mocks.NewCreateAdvertismentPresenter(t)

	uc := usecase.NewCreateAdvertismentInteractor(advRepoMock, catRepoMock, createAdvPresMock)
	_, err := uc.Execute(ctx, in)

	catRepoMock.AssertExpectations(t)

	if assert.Error(t, err) {
		assert.Equal(t, domain.ErrAdvMaxNameLength, err)
	}
}

// Тест проверки валидности объявления с превышенной длиной описания
func TestCreateAdvertismentWithMaxDescLengthExceeded(t *testing.T) {
	ctx := context.Background()

	cat := domain.Category{
		ID:   uuid.NewString(),
		Name: "some category",
	}

	photoUrls := generatePhotoUrls(domain.MAX_PHOTO_COUNT)
	in := usecase.CreateAdvertismentInput{
		Name:        "myadv",
		CategoryID:  cat.ID,
		Description: randSeq(domain.MAX_DESC_LENGTH + 1),
		Price: domain.Price{
			Value:    1000,
			Currency: "rub",
		},
		MainPhotoURL:        photoUrls[0],
		AdditionalPhotoURLs: photoUrls[1:],
	}

	advRepoMock := mocks.NewAdvertismentRepository(t)

	catRepoMock := mocks.NewCategoryRepository(t)
	catRepoMock.On("FindByID", ctx, cat.ID).Return(cat, nil)

	createAdvPresMock := mocks.NewCreateAdvertismentPresenter(t)

	uc := usecase.NewCreateAdvertismentInteractor(advRepoMock, catRepoMock, createAdvPresMock)
	_, err := uc.Execute(ctx, in)

	catRepoMock.AssertExpectations(t)

	if assert.Error(t, err) {
		assert.Equal(t, domain.ErrAdvMaxDescLength, err)
	}
}

// Сопоставить вход юзкейса и модель объявления
func matchInput(in usecase.CreateAdvertismentInput) func(a domain.Advertisment) bool {
	return func(a domain.Advertisment) bool {
		if a.Name != in.Name {
			return false
		} else if a.Category.ID != in.CategoryID {
			return false
		} else if a.Description != in.Description {
			return false
		} else if a.Price != in.Price {
			return false
		} else if a.MainPhotoURL != in.MainPhotoURL {
			return false
		}
		// } else if a.AdditionalPhotoURLs != in.AdditionalPhotoURLs {
		// 	return false
		// }

		return true
	}
}

func generatePhotoUrls(count int) []string {
	var urls []string
	for i := 0; i < count; i++ {
		urls = append(urls, fmt.Sprintf("http://example.com/photo%d.jpg", i))
	}
	return urls
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
