package userRepository

import (
	"abcShop/models"

	"gorm.io/gorm"
)

type IUserRepository interface {
	Find() ([]models.User, error)
	FindOne() (*models.User, error)
	Create(model *models.User) error
	Update(model *models.User) error
	Delete(model *models.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{
		db: db,
	}
}

func (repo *userRepository) Find() ([]models.User, error) {
	return nil, nil
}

func (repo *userRepository) FindOne() (*models.User, error) {
	return nil, nil
}

func (repo *userRepository) Create(model *models.User) error {
	return nil
}

func (repo *userRepository) Update(model *models.User) error {
	return nil
}

func (repo *userRepository) Delete(model *models.User) error {
	return nil
}
