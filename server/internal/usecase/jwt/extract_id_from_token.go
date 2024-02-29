package jwt

import (
	"fmt"

	_jwt "github.com/golang-jwt/jwt/v5"
)

func (j *UseCase) ExtractIDFromToken(requestToken string) (string, error) {
	token, err := _jwt.Parse(requestToken, func(token *_jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*_jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(j.Secret), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(_jwt.MapClaims)

	if !ok && !token.Valid {
		return "", fmt.Errorf("Invalid Token")
	}

	return claims["id"].(string), nil
}
