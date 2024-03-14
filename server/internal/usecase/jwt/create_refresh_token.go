package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/imi-courses/scrumproject-ggteam/server/internal/dto"
)

type CustomRefreshTokenClaims struct {
	dto.RefreshTokenPayload
	jwt.RegisteredClaims
}

func (j *UseCase) CreateRefreshToken(payload dto.RefreshTokenPayload) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomRefreshTokenClaims{
		payload,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.RefreshExpiresAt)),
		},
	})
	t, err := token.SignedString([]byte(j.RefreshTokenSecret))
	if err != nil {
		return "", err
	}
	return t, err
}
