package dto

type CreateAdmin struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password"`
}

type SignUpAdmin CreateAdmin

type SignInAdmin struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password"`
}
