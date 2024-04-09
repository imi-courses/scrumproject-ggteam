package dto

import (
	"github.com/google/uuid"
)

type CreateAdmin struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password"`
}

type SignUpAdmin CreateAdmin

type SignInAdmin struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password"`
}

type FindOneAdmin struct {
	ID    uuid.UUID `json:"id" binding:"uuid"`
	Email string    `json:"email" binding:"email"`
}
