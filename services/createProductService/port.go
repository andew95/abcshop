package createProductService

import (
	"abcShop/enums"
	"abcShop/models"
	"time"

	"github.com/google/uuid"
)

type Request struct {
	Code        string    `json:"code"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Tags        []string  `json:"tags"`
	Amount      int       `json:"amount"`
	Price       float64   `json:"price"`
	CreatedBy   uuid.UUID `json:"createdBy"`
}

type Response struct {
	Id          string   `json:"id"`
	Code        string   `json:"code"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Category    string   `json:"category"`
	Tags        []string `json:"tags"`
	Amount      int      `json:"amount"`
	Price       float64  `json:"price"`
	Status      string   `json:"status"`
	CreatedAt   string   `json:"createdAt"`
	CreatedBy   string   `json:"createdBy"`
}

func (r *Request) ToCreateRequest() *models.Product {
	return &models.Product{
		Id:          uuid.New(),
		Code:        r.Code,
		Name:        r.Name,
		Description: r.Description,
		Category:    r.Category,
		Tags:        r.Tags,
		Status:      enums.PRODUCT_STATUS_ACTIVE,
		Amount:      r.Amount,
		Price:       r.Price,
		CreatedAt:   time.Now(),
		CreatedBy:   r.CreatedBy,
	}
}

func ModelToResponse(model *models.Product) *Response {
	return &Response{
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
