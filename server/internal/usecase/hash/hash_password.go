package hash

import (
	"golang.org/x/crypto/bcrypt"
)

func (h *UseCase) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
