package updateUserPasswordService

import "github.com/google/uuid"

type Request struct {
	Id          uuid.UUID `json:"id"`
	NewPassword string    `json:"newPassword"`
	OldPassword string    `json:"oldPassword"`
}

type UserResponse struct {
	Id        string `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Role      string `json:"password"`
	CreatedAt string `json:"createdAt"`
}

type Response struct {
	User  UserResponse `json:"user"`
	Token string       `json:"token"`
}
