package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ecommerce/product/domain/entities"
	"github.com/ecommerce/product/infrastructure/models"
	"github.com/ecommerce/product/presentation/interfaces/http/data_objects"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	app, repoContainer := setup(t)

	products := []models.Product{
		{
			Name:        "test2 product",
			Description: "test product description",
			Price:       100,
		},
		{
			Name:        "second test product",
			Description: "second test product description",
			Price:       200,
		},
		{
			Name:        "third test product",
			Description: "third test product description",
			Price:       300,
		},
	}

	repoContainer.DB.Create(&products)

	t.Run("should get products", func(t *testing.T) {
		request := httptest.NewRequest("GET", "/api/products", nil)
		request.Header.Add("Content-Type", "application-json")

		res, err := app.Test(request)
		if err != nil {
			t.Error(err)
		}

		var productEntities []*entities.ProductEntity
		for _, model := range products {
			productEntities = append(productEntities, model.ToEntity())
		}

		encodedRes, err := io.ReadAll(res.Body)
		if err != nil {
			t.Error(err)
		}
		exceptedRes, _ := json.Marshal(data_objects.ListProductResponse{
			Success: true,
			Message: "Success",
			Data:    productEntities,
		})

		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.Equal(t, string(exceptedRes), string(encodedRes))
	})
}
