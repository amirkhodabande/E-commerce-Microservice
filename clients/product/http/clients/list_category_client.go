package clients

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/ecommerce/clients/product/http/data_objects"
	"github.com/joho/godotenv"
)

func ListCategory() (*data_objects.ListCategoryResponse, error) {
	err := godotenv.Load("../clients/.env")
	if err != nil {
		panic(err)
	}

	requestURL := "/api/categories"
	request, err := http.NewRequest("GET", os.Getenv("PRODUCT_SERVICE_URL")+requestURL, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add("Content-Type", "application/json")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var listCategoryResponse *data_objects.ListCategoryResponse
	err = json.Unmarshal(body, &listCategoryResponse)
	if err != nil {
		return nil, err
	}
	listCategoryResponse.Status = response.StatusCode

	return listCategoryResponse, nil
}
