package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Client struct {
	ID          uuid.UUID      `json:"id"           gorm:"uuid;default:gen_random_uuid();primarykey"`
	Surname     string         `json:"surname"`
	Firstname   string         `json:"firstname"`
	Middlename  string         `json:"middlename"`
	Email       string         `json:"email"`
	Phone       string         `json:"phone"        gorm:"unique"`
	RealEstates []RealEstate   `json:"real_estates" gorm:"foreignKey:ClientID"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"   gorm:"index"`
}
