package dto

import "github.com/google/uuid"

type CreateRealEstate struct {
	Address  string    `json:"address"`
	Type     string    `json:"type"`
	ClientID uuid.UUID `json:"client_id"`
}

type CreateRealEstateRequest struct {
	Address  string `json:"address"   binding:"required"`
	Type     string `json:"type"      binding:"required"`
	ClientID string `json:"client_id" binding:"required,uuid"`
}

type UpdateRealEstate struct {
	Address  string    `json:"address"   binding:"omitempty"`
	Type     string    `json:"type"      binding:"omitempty"`
	ClientID uuid.UUID `json:"client_id" binding:"omitempty,uuid"`
}
