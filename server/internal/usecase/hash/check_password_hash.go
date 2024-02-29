package hash

import (
	"golang.org/x/crypto/bcrypt"
)

func (h *UseCase) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
