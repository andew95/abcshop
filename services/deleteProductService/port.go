package deleteProductService

import "github.com/google/uuid"

type Request struct {
	Id uuid.UUID `json:"id" param:"id"`
}

type Response struct {
	Message string `json:"message"`
}
