package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RealEstate struct {
	ID        uuid.UUID      `json:"id"         gorm:"uuid;default:gen_random_uuid();primarykey"`
	Address   string         `json:"address"`
	Type      string         `json:"type"`
	ClientID  uuid.UUID      `json:"client_id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
