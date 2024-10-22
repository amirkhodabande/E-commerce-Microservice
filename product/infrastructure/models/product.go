package models

import (
	"github.com/ecommerce/product/domain/entities"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string
	Description string
	Price       float64
	CategoryID  uint
}

func (model *Product) ToEntity() *entities.ProductEntity {
	return &entities.ProductEntity{
		Id:          model.ID,
		Name:        model.Name,
		Description: model.Description,
		Price:       uint(model.Price),
	}
}
