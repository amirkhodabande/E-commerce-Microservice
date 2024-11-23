package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"

	"github.com/ecommerce/product/domain/entities"
	"github.com/ecommerce/product/infrastructure/models"
	"github.com/ecommerce/product/presentation/interfaces/http/data_objects"
	"github.com/stretchr/testify/assert"
)

func TestGetProducts(t *testing.T) {
	app, repoContainer := setup(t)

	parentCategory := models.Category{
		Name:        "parent category",
		Description: "parent category description",
		ParentID:    nil,
	}
	repoContainer.DB.Create(&parentCategory)

	products := []models.Product{
		{
			Name:        "test product",
			Description: "test product description",
			Price:       100,
		},
		{
			Name:        "second test product",
			Description: "second test product description",
			Price:       200,
			CategoryID:  parentCategory.ID,
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

	t.Run("name filter is working", func(t *testing.T) {
		baseURL := "/api/products"
		params := url.Values{}
		params.Add("name", "third test")

		requestURL := baseURL + "?" + params.Encode()
		request := httptest.NewRequest("GET", requestURL, nil)
		request.Header.Add("Content-Type", "application-json")

		res, err := app.Test(request)
		if err != nil {
			t.Error(err)
		}

		var productEntities []*entities.ProductEntity
		productEntities = append(productEntities, products[2].ToEntity())

		encodedRes, err := io.ReadAll(res.Body)
		if err != nil {
			t.Error(err)
		}
		expectedRes, _ := json.Marshal(data_objects.ListProductResponse{
			Success: true,
			Message: "Success",
			Data:    productEntities,
		})

		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.Equal(t, string(expectedRes), string(encodedRes))
	})

	t.Run("category filter is working", func(t *testing.T) {
		baseURL := "/api/products"
		params := url.Values{}

		params.Add("category_id", strconv.Itoa(int(parentCategory.ID)))

		requestURL := baseURL + "?" + params.Encode()
		request := httptest.NewRequest("GET", requestURL, nil)
		request.Header.Add("Content-Type", "application-json")

		res, err := app.Test(request)
		if err != nil {
			t.Error(err)
		}

		var productEntities []*entities.ProductEntity
		productEntities = append(productEntities, products[1].ToEntity())

		encodedRes, err := io.ReadAll(res.Body)
		if err != nil {
			t.Error(err)
		}
		expectedRes, _ := json.Marshal(data_objects.ListProductResponse{
			Success: true,
			Message: "Success",
			Data:    productEntities,
		})

		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.Equal(t, string(expectedRes), string(encodedRes))
	})
}
