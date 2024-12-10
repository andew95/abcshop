package app

import (
	"abcShop/repository/productRepository"
	"abcShop/repository/userRepository"

	"gorm.io/gorm"
)

type SetupRepository struct {
	UserRepo    userRepository.IUserRepository
	ProductRepo productRepository.IProductRepository
}

func NewSetupRepository(db *gorm.DB) *SetupRepository {
	userRepo := userRepository.NewUserRepository(db)
	productRepo := productRepository.NewProductRepository(db)
	return &SetupRepository{
		UserRepo:    userRepo,
		ProductRepo: productRepo,
	}
}
