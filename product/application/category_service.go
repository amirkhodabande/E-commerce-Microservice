package application

import (
	"log"
	"net/http"

	"github.com/ecommerce/product/domain/contracts"
	"github.com/ecommerce/product/domain/entities"
	"github.com/ecommerce/product/presentation/interfaces/http/data_objects"
)

type CategoryService struct {
	contracts.CategoryRepository
}

func NewCategoryService(repository contracts.CategoryRepository) *CategoryService {
	return &CategoryService{
		CategoryRepository: repository,
	}
}

func (service *CategoryService) GetCategories(requestID string) ([]*entities.CategoryEntity, error) {
	categories, err := service.CategoryRepository.FindAll()

	if err != nil {
		log.Printf("Request ID: %s, Error: %s", requestID, err.Error())

		return nil, data_objects.NewApiError(http.StatusInternalServerError, "internal server error")
	}

	return categories, nil
}
