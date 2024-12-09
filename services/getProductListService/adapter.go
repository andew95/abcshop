package getProductListService

import "abcShop/repository/productRepository"

type IGetProductListService interface {
	Execute(request Request) (*Response, error)
}

type getProductListService struct {
	repo productRepository.IProductRepository
}

func NewGetProductListService(repo productRepository.IProductRepository) IGetProductListService {
	return &getProductListService{
		repo: repo,
	}
}

func (s *getProductListService) Execute(request Request) (*Response, error) {
	return nil, nil
}
