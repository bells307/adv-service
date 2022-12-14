package advertisment_handler

import (
	"net/http"

	"github.com/bells307/adv-service/internal/domain/advertisment/model"
	"github.com/gin-gonic/gin"
)

type GetAdvertismentResponse struct {
	Name                string      `json:"name"`
	Price               model.Price `json:"price"`
	MainPhotoURL        string      `json:"mainPhotoURL"`
	Description         *string     `json:"description,omitempty"`
	AdditionalPhotoURLs *[]string   `json:"additionalPhotoURLs,omitempty"`
}

type GetAdvertismentManyResponse struct {
	Name         string      `json:"name"`
	Price        model.Price `json:"price"`
	MainPhotoURL string      `json:"mainPhotoURL"`
}

type errorResponse struct {
	Message string `json:"message"`
}

func ErrorResponse(c *gin.Context, err error) {
	resp := errorResponse{
		Message: err.Error(),
	}
	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": resp})
}
