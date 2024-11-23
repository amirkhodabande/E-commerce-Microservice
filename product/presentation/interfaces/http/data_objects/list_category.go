package data_objects

import "github.com/ecommerce/product/domain/entities"

type ListCategoryResponse struct {
	Success bool                       `json:"success"`
	Message string                     `json:"message"`
	Data    []*entities.CategoryEntity `json:"data"`
}
