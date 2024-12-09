package models

import (
	"abcShop/enums"
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id           uuid.UUID
	Name         string
	Address      string
	Username     string
	Email        string
	Password     string
	HashPassword string
	SaltKey      string
	Role         enums.UserRole
	Status       enums.UserStatus
	CreatedAt    time.Time
}
