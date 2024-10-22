package application

import (
	"log"
	"net/http"

	"github.com/ecommerce/product/domain/contracts"
	"github.com/ecommerce/product/domain/entities"
	"github.com/ecommerce/product/presentation/interfaces/http/data_objects"
)

type ProductService struct {
	contracts.ProductRepository
}

func NewProductService(repository contracts.ProductRepository) *ProductService {
	return &ProductService{
		ProductRepository: repository,
	}
}

func (service *ProductService) GetProducts(requestID string) ([]*entities.ProductEntity, error) {
	products, err := service.ProductRepository.FindAll()

	if err != nil {
		log.Printf("Request ID: %s, Error: %s", requestID, err.Error())

		return nil, data_objects.NewApiError(http.StatusInternalServerError, "internal server error")
	}

	return products, nil
}
