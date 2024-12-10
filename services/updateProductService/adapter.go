package updateProductService

import (
	"abcShop/models"
	"abcShop/repository/productRepository"
)

type IUpdateProductService interface {
	Execute(request Request) (*Response, error)
}

type updateProductService struct {
	repo productRepository.IProductRepository
}

func NewUpdateProductService(repo productRepository.IProductRepository) IUpdateProductService {
	return &updateProductService{
		repo: repo,
	}
}

func (s *updateProductService) Execute(request Request) (*Response, error) {
	productModel, err := s.repo.FindOne(&models.Product{Id: request.Id})
	if err != nil {
		return nil, err
	}
	productModel = request.ToUpdateRequest(productModel)

	err = s.repo.Update(productModel)
	if err != nil {
		return nil, err
	}

	response := ToResponse(productModel)
	return &response, nil
}
