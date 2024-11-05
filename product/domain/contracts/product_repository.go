package contracts

import (
	"github.com/ecommerce/product/domain/entities"
	"github.com/ecommerce/product/presentation/interfaces/http/data_objects"
)

type ProductRepository interface {
	FindAll(queryParams *data_objects.ListProductParams) ([]*entities.ProductEntity, error)
}
