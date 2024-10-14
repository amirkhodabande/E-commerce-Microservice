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

func LoginUser(data data_objects.LoginUserData) (*data_objects.LoginUserResponse, error) {
	err := godotenv.Load("../clients/.env")
	if err != nil {
		panic(err)
	}

	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", os.Getenv("USER_SERVICE_URL")+"/api/user/login", bytes.NewReader(payload))
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

	var loginResponse *data_objects.LoginUserResponse
	err = json.Unmarshal(body, &loginResponse)
	if err != nil {
		return nil, err
	}
	loginResponse.Status = response.StatusCode

	return loginResponse, nil
}
