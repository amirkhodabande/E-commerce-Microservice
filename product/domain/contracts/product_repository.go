package contracts

import "github.com/ecommerce/product/domain/entities"

type ProductRepository interface {
	FindAll() ([]*entities.ProductEntity, error)
}
