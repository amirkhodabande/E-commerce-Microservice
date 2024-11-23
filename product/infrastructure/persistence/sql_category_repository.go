package persistence

import (
	"fmt"

	"github.com/ecommerce/product/domain/entities"
	"github.com/ecommerce/product/infrastructure/models"
	"gorm.io/gorm"
)

type SqlCategoryRepository struct {
	db *gorm.DB
}

func NewSqlCategoryRepository(db *gorm.DB) *SqlCategoryRepository {
	return &SqlCategoryRepository{
		db: db,
	}
}

func (r *SqlCategoryRepository) FindAll() ([]*entities.CategoryEntity, error) {
	var categoryModels []models.Category

	result := r.db.Preload("Children").Where("parent_id IS NULL").Find(&categoryModels)

	if result.Error != nil {
		fmt.Printf("Error fetching categories: %+v\n", result.Error)
		return nil, result.Error
	}

	var categories []*entities.CategoryEntity

	for _, model := range categoryModels {
		categories = append(categories, model.ToEntity())
	}

	return categories, nil
}
