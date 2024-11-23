package contracts

import (
	"github.com/ecommerce/product/domain/entities"
)

type CategoryRepository interface {
	FindAll() ([]*entities.CategoryEntity, error)
}
