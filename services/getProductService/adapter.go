package getProductService

import "abcShop/repository/productRepository"

type IGetProductService interface{}

type getProductService struct {
	repo productRepository.IProductRepository
}

func NewCreateProductService(repo productRepository.IProductRepository) IGetProductService {
	return &getProductService{
		repo: repo,
	}
}

func (s *getProductService) Execute(request Response) (*Response, error) {
	return nil, nil
}
