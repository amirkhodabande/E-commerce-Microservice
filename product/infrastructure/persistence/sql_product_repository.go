package persistence

import (
	"fmt"

	"github.com/ecommerce/product/domain/entities"
	"github.com/ecommerce/product/infrastructure/models"
	"github.com/ecommerce/product/presentation/interfaces/http/data_objects"
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

func (r *SqlProductRepository) FindAll(queryParams *data_objects.ListProductParams) ([]*entities.ProductEntity, error) {
	limit := queryParams.Limit
	if limit == 0 {
		limit = 10
	}
	offset := (queryParams.Page - 1) * limit

	var productModels []models.Product
	query := r.db.Limit(limit).Offset(offset).
		Where("name LIKE ?", "%"+queryParams.Name+"%")

	switch queryParams.SortBy {
	case "most_expensive":
		query = query.Order("price DESC")
	case "cheapest":
		query = query.Order("price ASC")
	case "newest":
		query = query.Order("created_at DESC")
	}

	result := query.Find(&productModels)

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
