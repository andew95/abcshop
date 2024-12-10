package updateProductService

import (
	"abcShop/enums"
	"abcShop/models"
	"time"

	"github.com/google/uuid"
)

type Request struct {
	Id          uuid.UUID
	Code        string
	Name        string
	Description string
	Category    string
	Tags        []string
	Status      string
	Amount      int
	Price       float64
}

type Response struct {
	Id          string   `json:"id"`
	Code        string   `json:"code"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Category    string   `json:"category"`
	Tags        []string `json:"tags"`
	Status      string   `json:"status"`
	Amount      int      `json:"amount"`
	Price       float64  `json:"price"`
	CreatedAt   string   `json:"createdAt"`
	CreatedBy   string   `json:"createdBy"`
}

func (r Request) ToUpdateRequest(model *models.Product) *models.Product {
	model.Code = r.Code
	model.Description = r.Description
	model.Category = r.Category
	model.Tags = r.Tags
	model.Status = enums.NewProductStatus(r.Status)
	model.Amount = r.Amount
	model.Price = r.Price
	return model
}

func ToResponse(model *models.Product) Response {
	return Response{
		Id:          model.Id.String(),
		Code:        model.Code,
		Name:        model.Name,
		Description: model.Description,
		Category:    model.Category,
		Tags:        model.Tags,
		Status:      model.Status.ToString(),
		Amount:      model.Amount,
		Price:       model.Price,
		CreatedAt:   model.CreatedAt.Format(time.ANSIC),
		CreatedBy:   model.CreatedBy.String(),
	}
}
