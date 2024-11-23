package data_objects

import "github.com/ecommerce/product/domain/entities"

type ListProductResponse struct {
	Success bool                      `json:"success"`
	Message string                    `json:"message"`
	Data    []*entities.ProductEntity `json:"data"`
}

type ListProductParams struct {
	Name       string `query:"name,omitempty" validate:"omitempty"`
	CategoryID uint   `query:"category_id,omitempty" validate:"omitempty,min=1"`
	Page       int    `query:"page,omitempty" validate:"omitempty,min=1"`
	Limit      int    `query:"limit,omitempty" validate:"omitempty,min=1,max=100"`
	SortBy     string `query:"sort_by,omitempty" validate:"omitempty,oneof=most_expensive cheapest newest"`
}
