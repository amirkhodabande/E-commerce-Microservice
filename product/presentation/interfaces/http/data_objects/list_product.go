package data_objects

import "github.com/ecommerce/product/domain/entities"

type ListProductResponse struct {
	Success bool                      `json:"success"`
	Message string                    `json:"message"`
	Data    []*entities.ProductEntity `json:"data"`
}
