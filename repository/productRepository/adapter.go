package productRepository

import (
	"abcShop/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IProductRepository interface {
	Create(productModel *models.Product) error
	Find(conds *models.Product) ([]models.Product, error)
	FindOne(conds *models.Product) (*models.Product, error)
	Update(productModel *models.Product) error
	Delete(id uuid.UUID) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) IProductRepository {
	return &productRepository{
		db: db,
	}
}

func (repo *productRepository) Create(productModel *models.Product) error {
	err := repo.db.Debug().Create(productModel).Error
	return err
}

func (repo *productRepository) Find(conds *models.Product) ([]models.Product, error) {
	var products []models.Product

	err := repo.db.Find(&products, conds).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (repo *productRepository) FindOne(conds *models.Product) (*models.Product, error) {
	var product models.Product
	err := repo.db.First(&product, conds).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (repo *productRepository) Update(productModel *models.Product) error {
	err := repo.db.Updates(&productModel).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *productRepository) Delete(id uuid.UUID) error {
	err := repo.db.Delete(&models.Product{}, &models.Product{Id: id}).Error
	return err
}
