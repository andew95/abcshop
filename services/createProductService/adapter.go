package createProductService

import "abcShop/repository/productRepository"

type ICreateProductService interface {
	Execute(request Request) (*Response, error)
}

type createProductService struct {
	repo productRepository.IProductRepository
}

func NewCreateProductService(repo productRepository.IProductRepository) ICreateProductService {
	return &createProductService{
		repo: repo,
	}
}

func (s *createProductService) Execute(request Request) (*Response, error) {
	productModel := request.ToCreateRequest()
	err := s.repo.Create(productModel)
	if err != nil {
		return nil, err
	}
	
	response := ModelToResponse(productModel)
	return response, nil
}
