package getProductListService

import "abcShop/repository/productRepository"

type IGetProductListService interface {
	Execute(request Request) ([]Response, error)
}

type getProductListService struct {
	repo productRepository.IProductRepository
}

func NewGetProductListService(repo productRepository.IProductRepository) IGetProductListService {
	return &getProductListService{
		repo: repo,
	}
}

func (s *getProductListService) Execute(request Request) ([]Response, error) {
	products, err := s.repo.Find(nil)
	if err != nil {
		return nil, err
	}

	var resp []Response
	for _, productModel := range products {
		resp = append(resp, ToResponse(&productModel))
	}
	return resp, nil
}
