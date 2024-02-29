package exception

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type messageResponse struct {
	Message string `json:"message"`
}

func ErrorResponse(c *gin.Context, code int, message string) {
	c.AbortWithStatusJSON(code, messageResponse{message})
}

func BadRequest(c *gin.Context, message string) {
	c.AbortWithStatusJSON(http.StatusBadRequest, messageResponse{message})
}

func InternalServerError(c *gin.Context, message string) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, messageResponse{message})
}
