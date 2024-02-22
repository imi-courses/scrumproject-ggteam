package usecase

import (
	"golang.org/x/crypto/bcrypt"
)

type HashPasswordUseCase struct{}

func newHashPassword() *HashPasswordUseCase {
	return &HashPasswordUseCase{}
}

func (h *HashPasswordUseCase) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (h *HashPasswordUseCase) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
