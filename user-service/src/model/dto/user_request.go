package dto

type UpdateUserReq struct {
	UserId       string  `json:"user_id" validate:"required,max=21,min=21"`
	Email        string  `json:"email" validate:"omitempty,email"`
	FullName     string  `json:"full_name" validate:"omitempty,max=100"`
	Whatsapp     string `json:"whatsapp" validate:"omitempty,max=20"`
	Password     string  `json:"password" validate:"omitempty,max=100"`
	RefreshToken string  `json:"refresh_token" validate:"omitempty,max=500"`
}
