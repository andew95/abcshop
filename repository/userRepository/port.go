package userRepository

import "github.com/google/uuid"

type Request struct {
	Id       uuid.UUID
	Username string
}

type Response struct {
}
