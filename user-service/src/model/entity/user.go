package entity

import "time"

type User struct {
	UserId       string    `json:"user_id"`
	Email        string    `json:"email"`
	FullName     string    `json:"full_name"`
	Whatsapp     *string   `json:"whatsapp"`
	Role         string    `json:"role"`
	Password     string    `json:"password"`
	RefreshToken *string   `json:"refresh_token"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type SanitizedUser struct {
	UserId    string    `json:"user_id"`
	Email     string    `json:"email"`
	FullName  string    `json:"full_name"`
	Whatsapp  *string   `json:"whatsapp"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type RefreshToken struct {
	Token string `json:"refresh_token" validate:"required,max=500"`
}