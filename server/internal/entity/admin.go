package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Admin struct {
	ID        uuid.UUID      `json:"id" gorm:"uuid;default:gen_random_uuid();primarykey"`
	Email     string         `json:"email" gorm:"unique"`
	Password  string         `json:"password"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
