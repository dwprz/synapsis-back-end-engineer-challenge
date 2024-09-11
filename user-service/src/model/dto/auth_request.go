package dto

type RegisterReq struct {
	UserId   string `json:"user_id" validate:"omitempty"`
	Email    string `json:"email" validate:"required,email,min=5,max=100"`
	FullName string `json:"full_name" validate:"required,min=3,max=100"`
	Password string `json:"password" validate:"required,min=5,max=100"`
}

type LoginReq struct {
	Email    string `json:"email" validate:"required,email,min=5,max=100"`
	Password string `json:"password" validate:"required,min=5,max=100"`
}