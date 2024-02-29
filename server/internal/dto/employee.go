package dto

type CreateEmployee struct {
	Surname    string `json:"surname"`
	Firstname  string `json:"firstname"`
	Middlename string `json:"middlename"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password"`
}
