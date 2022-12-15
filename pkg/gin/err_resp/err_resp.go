package err_resp

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Message string `json:"message"`
}

func ErrorResponse(c *gin.Context, err error) {
	resp := errorResponse{
		Message: err.Error(),
	}
	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": resp})
}
