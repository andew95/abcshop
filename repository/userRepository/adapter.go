package userRepository

import (
	"abcShop/enums"
	"abcShop/models"

	"gorm.io/gorm"
)

type IUserRepository interface {
	Find() ([]models.User, error)
	FindOne(userModel *models.User) (*models.User, error)
	Create(userModel *models.User) error
	Update(userModel *models.User) error
	Delete(userModel *models.User) error
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
	var users []models.User
	err := repo.db.Not("status !=", enums.USER_STATUS_INACTIVE).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *userRepository) FindOne(userModel *models.User) (*models.User, error) {
	var record models.User

	err := repo.db.Debug().First(&record, userModel).Error
	if err != nil {
		return nil, err
	}

	return &record, nil
}

func (repo *userRepository) Create(userModel *models.User) error {
	err := repo.db.Debug().Create(userModel).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *userRepository) Update(model *models.User) error {
	err := repo.db.Debug().Updates(model).Error
	return err
}

func (repo *userRepository) Delete(model *models.User) error {
	return nil
}
