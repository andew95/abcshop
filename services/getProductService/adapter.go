package getProductService

import "abcShop/repository/productRepository"

type IGetProductService interface {
	Execute(request Request) (*Response, error)
}

type getProductService struct {
	repo productRepository.IProductRepository
}

func NewCreateProductService(repo productRepository.IProductRepository) IGetProductService {
	return &getProductService{
		repo: repo,
	}
}

func (s *getProductService) Execute(request Request) (*Response, error) {
	conds := request.ToFindOneRequest()
	productModel, err := s.repo.FindOne(conds)
	if err != nil {
		return nil, err
	}

	response := ToResponse(productModel)
	return &response, nil
}
