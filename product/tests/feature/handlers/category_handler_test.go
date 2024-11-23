package handlers

import (
	"encoding/json"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/ecommerce/product/infrastructure/models"
	"github.com/ecommerce/product/presentation/interfaces/http/data_objects"
	"github.com/stretchr/testify/assert"
)

func TestGetCategories(t *testing.T) {
	app, repoContainer := setup(t)

	parentCategories := []models.Category{
		{
			Name:        "parent category",
			Description: "parent category description",
			ParentID:    nil,
		},
		{
			Name:        "second parent category",
			Description: "second parent category description",
			ParentID:    nil,
		},
	}
	repoContainer.DB.Create(&parentCategories)

	childCategories := []models.Category{
		{
			Name:        "child category",
			Description: "child category description",
			ParentID:    &parentCategories[0].ID,
		},
	}
	repoContainer.DB.Create(&childCategories)

	t.Run("should get all categories", func(t *testing.T) {
		request := httptest.NewRequest("GET", "/api/categories", nil)
		request.Header.Add("Content-Type", "application-json")

		res, err := app.Test(request)
		if err != nil {
			t.Error(err)
		}

		encodedRes, err := io.ReadAll(res.Body)
		if err != nil {
			t.Error(err)
		}

		var jsonResponse data_objects.ListCategoryResponse
		err = json.Unmarshal(encodedRes, &jsonResponse)
		if err != nil {
			t.Error(err)
		}

		assert.True(t, jsonResponse.Success, "Expected success to be true")

		assert.Equal(t, 2, len(jsonResponse.Data), "Expected 2 categories in the response")

		if len(jsonResponse.Data) > 0 {
			assert.Equal(t, 1, len(jsonResponse.Data[0].Children), "Expected first category to have 1 child")
		}
	})
}
