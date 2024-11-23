package models

import (
	"github.com/ecommerce/product/domain/entities"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name        string
	Description string
	CategoryID  uint
	Children    []Category `gorm:"foreignKey:category_id"`
}

func (model *Category) ToEntity() *entities.CategoryEntity {
	return &entities.CategoryEntity{
		Id:          model.ID,
		Name:        model.Name,
		Description: model.Description,
		Categories:  model.Children,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
		DeletedAt:   model.DeletedAt.Time,
	}
}
