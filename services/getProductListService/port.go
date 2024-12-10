package getProductListService

import (
	"abcShop/models"
	"time"
)

type Request struct {
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
