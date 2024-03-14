package jwt

import (
	"fmt"

	jwt "github.com/golang-jwt/jwt/v5"
)

func (j *UseCase) IsTokenValid(requestToken string, isAccessToken bool) (bool, error) {
	secret := j.RefreshTokenSecret
	if isAccessToken {
		secret = j.AccessTokenSecret
	}
	_, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return false, err
	}
	return true, nil
}
