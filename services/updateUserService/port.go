package updateUserService

import "github.com/google/uuid"

type Request struct {
	Id      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Address string    `json:"address"`
	Email   string    `json:"email"`
}

type UserResponse struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Address   string `json:"address"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Status    string `json:"status"`
	CreatedAt string `json:"createdAt"`
}

type Response struct {
	User  UserResponse `json:"user"`
	Token string       `json:"token"`
}
