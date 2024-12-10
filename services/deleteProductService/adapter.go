package deleteProductService

import (
	"abcShop/models"
	"abcShop/repository/productRepository"

	"gorm.io/gorm"
)

type IDeleteProductService interface {
	Execute(request Request) error
}

type deleteProductService struct {
	repo productRepository.IProductRepository
}

func NewDeleteProductService(repo productRepository.IProductRepository) IDeleteProductService {
	return &deleteProductService{
		repo: repo,
	}
}

func (s *deleteProductService) Execute(request Request) error {
	product, err := s.repo.FindOne(&models.Product{Id: request.Id})
	if err != nil {
		return err
	}
	if product == nil {
		return gorm.ErrRecordNotFound
	}
	err = s.repo.Delete(request.Id)
	if err != nil {
		return err
	}
	return nil
}
