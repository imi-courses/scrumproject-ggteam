package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Person struct {
	ID         uuid.UUID `json:"id" gorm:"uuid;default:gen_random_uuid();primarykey"`
	Firstname  string
	Middlename string
	Surname    string
	UserID     uuid.UUID `json:"user_id" gorm:"uuid"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
