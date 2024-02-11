package dto

type CreatePerson struct {
	Firstname  string `json:"firstname"`
	Middlename string `json:"middlename"`
	Surname    string `json:"surname"`
}
