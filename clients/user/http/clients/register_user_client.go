package clients

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/ecommerce/clients/user/http/data_objects"
	"github.com/joho/godotenv"
)

func RegisterUser(data data_objects.RegisterUserData) (*data_objects.RegisterUserResponse, error) {
	err := godotenv.Load("../clients/.env")
	if err != nil {
		panic(err)
	}

	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", os.Getenv("USER_SERVICE_URL")+"/api/user/register", bytes.NewReader(payload))
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

	var registerResponse data_objects.RegisterUserResponse
	err = json.Unmarshal(body, &registerResponse)
	if err != nil {
		return nil, err
	}
	registerResponse.Status = response.StatusCode

	return &registerResponse, nil
}
