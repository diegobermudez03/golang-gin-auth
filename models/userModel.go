package models

import "time"

type User struct {
	ID           int 		`json:"id"`
	FirstName    string 	`json:"first_name" validate:"required, min=2, max=100"`
	LastName     string 	`json:"last_name" validate:"required, min=2, max=100"`
	Password     string 	`json:"password" validate:"required, min=6, max=12"`
	Email        string 	`json:"email" validate:"email, required"`
	Phone        string		`json:"phone" validate:"required"`
	Token        string		`json:"token"`
	UserType     string		`json:"user_type" validate:"required"`
	RefreshToken string		`json:"refresh_token"`
	CreatedAt    time.Time	`json:"created_at"`
	UpdatedAt	time.Time	`json:"updated_at"`
	UserId		int			`json:"user_id"`
}