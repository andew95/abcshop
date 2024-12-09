package app

import (
	"abcShop/repository/userRepository"

	"gorm.io/gorm"
)

type SetupRepository struct {
	UserRepo userRepository.IUserRepository
}

func NewSetupRepository(db *gorm.DB) *SetupRepository {
	userRepo := userRepository.NewUserRepository(db)
	return &SetupRepository{
		UserRepo: userRepo,
	}
}
