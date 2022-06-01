package entity

import "time"

type userResponse struct {
	UserName          string    `json:"username"`
	FullName          string    `json:full_name`
	Email             string    `json:email`
	PasswordChangedAt time.Time `json:password_changed_at`
	CreateAt          time.Time `json:created_at`
}
type createUserRequest struct {
	UserName string `json:"username"`
	FullName string `json:full_name`
	Email    string `json:email`
	password string `json:"email"`
}
