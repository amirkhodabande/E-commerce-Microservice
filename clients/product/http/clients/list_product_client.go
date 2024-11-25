package clients

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/ecommerce/clients/product/http/data_objects"
	"github.com/google/go-querystring/query"
	"github.com/joho/godotenv"
)

func ListProduct(data data_objects.ListProductData, requestID string) (*data_objects.ListProductResponse, error) {
	err := godotenv.Load("../clients/.env")
	if err != nil {
		panic(err)
	}

	queryParams, err := query.Values(data)
	if err != nil {
		return nil, err
	}

	requestURL := "/api/products" + "?" + queryParams.Encode()
	request, err := http.NewRequest("GET", os.Getenv("PRODUCT_SERVICE_URL")+requestURL, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-Request-ID", requestID)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var listProductResponse *data_objects.ListProductResponse
	err = json.Unmarshal(body, &listProductResponse)
	if err != nil {
		return nil, err
	}
	listProductResponse.Status = response.StatusCode

	return listProductResponse, nil
}
