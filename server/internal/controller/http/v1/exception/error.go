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

func UnAuthorized(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, messageResponse{"unauthorized"})
}

func UnAuthorizedWithMessage(c *gin.Context, message string) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, messageResponse{message})
}
