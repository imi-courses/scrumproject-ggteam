package dto

type CreateClient struct {
	Surname    string `json:"surname"    binding:"required,alphaunicode"`
	Firstname  string `json:"firstname"  binding:"required,alphaunicode"`
	Middlename string `json:"middlename" binding:"omitempty,alphaunicode"`
	Email      string `json:"email"      binding:"omitempty,email"`
	Phone      string `json:"phone"      binding:"required,e164"`
}

type UpdateClient struct {
	Surname    string `json:"surname"    binding:"omitempty,alphaunicode"`
	Firstname  string `json:"firstname"  binding:"omitempty,alphaunicode"`
	Middlename string `json:"middlename" binding:"omitempty,alphaunicode"`
	Email      string `json:"email"      binding:"omitempty,email"`
	Phone      string `json:"phone"      binding:"omitempty,e164"`
}
