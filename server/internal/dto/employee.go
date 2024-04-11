package dto

import (
	"github.com/google/uuid"
)

type CreateEmployee struct {
	Surname    string `json:"surname"    binding:"required,alphaunicode"`
	Firstname  string `json:"firstname"  binding:"required,alphaunicode"`
	Middlename string `json:"middlename" binding:"omitempty,alphaunicode"`
	Email      string `json:"email"      binding:"required,email"`
	Password   string `json:"password"`
}

type CreateEmployeeWithoutPassword struct {
	Surname    string `json:"surname"    binding:"required,alphaunicode"`
	Firstname  string `json:"firstname"  binding:"required,alphaunicode"`
	Middlename string `json:"middlename" binding:"omitempty,alphaunicode"`
	Email      string `json:"email"      binding:"required,email"`
}

type UpdateEmployee struct {
	Surname    string `json:"surname"    binding:"omitempty,alphaunicode"`
	Firstname  string `json:"firstname"  binding:"omitempty,alphaunicode"`
	Middlename string `json:"middlename" binding:"omitempty,alphaunicode"`
	Email      string `json:"email"      binding:"omitempty,email"`
}

type FindOneEmployee struct {
	ID    uuid.UUID `json:"id"    binding:"uuid"`
	Email string    `json:"email" binding:"email"`
}
