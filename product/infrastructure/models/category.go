package models

import (
	"github.com/ecommerce/product/domain/entities"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name        string
	Description string
	ParentID    *uint      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Children    []Category `gorm:"foreignKey:parent_id"`
}

func (model *Category) ToEntity() *entities.CategoryEntity {
	var children []entities.CategoryEntity
	for _, child := range model.Children {
		children = append(children, *child.ToEntity())
	}

	return &entities.CategoryEntity{
		Id:          model.ID,
		Name:        model.Name,
		Description: model.Description,
		Children:    children,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
		DeletedAt:   model.DeletedAt.Time,
	}
}
