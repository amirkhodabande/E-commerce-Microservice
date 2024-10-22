package persistence

import (
	"fmt"

	"github.com/ecommerce/product/domain/entities"
	"github.com/ecommerce/product/infrastructure/models"
	"gorm.io/gorm"
)

type SqlProductRepository struct {
	db *gorm.DB
}

func NewSqlProductRepository(db *gorm.DB) *SqlProductRepository {
	return &SqlProductRepository{
		db: db,
	}
}

func (r *SqlProductRepository) FindAll() ([]*entities.ProductEntity, error) {
	var productModels []models.Product
	result := r.db.Find(&productModels)
	if result.Error != nil {
		fmt.Printf("Error fetching products: %+v\n", result.Error)
		return nil, result.Error
	}

	var products []*entities.ProductEntity
	for _, model := range productModels {
		products = append(products, model.ToEntity())
	}

	return products, nil
}
