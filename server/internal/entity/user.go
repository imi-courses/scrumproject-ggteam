package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role string

const (
	AdminRole    Role = "admin"
	EmployeeRole Role = "employee"
)

type User struct {
	ID        uuid.UUID `json:"id" gorm:"uuid;default:gen_random_uuid();primarykey"`
	Person    Person
	Role      Role
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
