package advertisment_grpc

import (
	"context"
	"github.com/bells307/adv-service/internal/adapter/presenter"
	"github.com/bells307/adv-service/internal/domain"
	"github.com/bells307/adv-service/internal/usecase"
	"google.golang.org/grpc"
)

type advertismentHandler struct {
	advRepo domain.AdvertismentRepository
	catRepo domain.CategoryRepository
}

func NewAdvertismentHandler(
	advRepo domain.AdvertismentRepository,
	catRepo domain.CategoryRepository,
) *advertismentHandler {
	return &advertismentHandler{advRepo, catRepo}
}

func (a advertismentHandler) Register(grpcServer *grpc.Server) {
	RegisterAdvertismentHandlerServer(grpcServer, a)
}

func (a advertismentHandler) CreateAdvertisment(
	ctx context.Context,
	input *CreateAdvertismentInput,
) (*CreateAdvertismentOutput, error) {
	uc := usecase.NewCreateAdvertismentInteractor(
		a.advRepo,
		a.catRepo,
		presenter.NewCreateAdvertismentPresenter(),
	)

	ucInput := usecase.CreateAdvertismentInput{
		Name:        input.Name,
		Categories:  input.Categories,
		Description: input.Description,
		Price: usecase.Price{
			Value:    float64(input.Price.Value),
			Currency: input.Price.Currency,
		},
		MainPhotoURL:        input.MainPhotoURL,
		AdditionalPhotoURLs: input.AdditionalPhotoURLs,
	}

	out, err := uc.Execute(ctx, ucInput)
	if err != nil {
		return nil, err
	}

	return &CreateAdvertismentOutput{
		Id: out.ID,
	}, nil
}

func (a advertismentHandler) mustEmbedUnimplementedAdvertismentHandlerServer() {}
