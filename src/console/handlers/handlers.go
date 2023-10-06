package handlers

import (
	"bytes"
	myErrors "console/errors"
	openapi "console/internal/client"
	"console/internal/consts"
	"fmt"
	"net/http"
)

const port = "8081"
const address = "localhost"

func DoRequest(client *http.Client, request *http.Request) (*http.Response, error) {
	response, err := client.Do(request)
	if err != nil {
		return nil, myErrors.ErrorResponse
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		fmt.Println(response.StatusCode)
		var p []byte
		_, err = response.Body.Read(p)
		fmt.Println(string(p))
		fmt.Println(err)
		return response, myErrors.ErrorResponseStatus
	}

	return response, nil
}

// func AuthorizeClient(ctx context.Context, client *openapi.DefaultAPIService, authRequest openapi.AuthRequest) (*http.Response, error) {
// user, response, err := client.Authorize(ctx).AuthRequest(authRequest).Execute()
func AuthorizeClient(client *http.Client, authRequest *openapi.AuthRequest) (*http.Response, error) {
	url := "http://" + address + ":" + port + "/authorize"
	params := fmt.Sprintf("{\"login\": \"%s\", \"password\": \"%s\"}", authRequest.Login, authRequest.Password)
	var jsonStr = []byte(params)

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, myErrors.ErrorNewRequest
	}

	request.Header.Set("Content-Type", "application/json")

	return DoRequest(client, request)
}

func CreateClient(client *http.Client, r *openapi.RegisterRequest) (*http.Response, error) {

	url := "http://" + address + ":" + port + "/register"
	params := fmt.Sprintf("{\"login\": \"%s\", \"password\": \"%s\", \"fio\": \"%s\", "+
		"\"email\": \"%s\", \"points\": %d, \"status\": \"%s\"}", r.Login, r.Password, r.Fio, r.Email,
		r.Points, r.Status)
	jsonStr := []byte(params)

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, myErrors.ErrorNewRequest
	}

	request.Header.Set("Content-Type", "application/json")

	return DoRequest(client, request)
}

func GetWines(client *http.Client, r *openapi.GetWinesRequest) (*http.Response, error) {
	url := "http://" + address + ":" + port + "/wines"
	params := fmt.Sprintf("{\"limit\": \"%s\", \"skip\": \"%s\"}", r.Limit, r.Skip)
	jsonStr := []byte(params)

	request, err := http.NewRequest("GET", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, myErrors.ErrorNewRequest
	}

	request.Header.Set("Content-Type", "application/json")

	response, err := DoRequest(client, request)

	return response, err
}

func CreateElem(client *http.Client, r *openapi.CreateElemRequest, login string, password string) (*http.Response, error) {
	url := "http://" + address + ":" + port + "/elems"
	params := fmt.Sprintf("{\"IdWine\": \"%s\", \"Count\": %d}", r.IdWine, r.Count)
	var jsonStr = []byte(params)

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, myErrors.ErrorNewRequest
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("login", login)
	request.Header.Set("password", password)

	response, err := DoRequest(client, request)

	return response, err
}

func AddElem(client *http.Client, r *openapi.AddElemRequest, login string, password string) (*http.Response, error) {
	url := "http://" + address + ":" + port + "/elems/" + r.Id + "/add"

	request, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		return nil, myErrors.ErrorNewRequest
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("login", login)
	request.Header.Set("password", password)

	response, err := DoRequest(client, request)

	return response, err
}

func DecreaseElem(client *http.Client, r *openapi.DecreaseElemRequest, login string, password string) (*http.Response, error) {
	url := "http://" + address + ":" + port + "/elems/" + r.Id + "/decrease"

	request, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		return nil, myErrors.ErrorNewRequest
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("login", login)
	request.Header.Set("password", password)

	response, err := DoRequest(client, request)

	return response, err
}

func DeleteElem(client *http.Client, id string, login string, password string) (*http.Response, error) {
	url := "http://" + address + ":" + port + "/elems/" + id

	request, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, myErrors.ErrorNewRequest
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("login", login)
	request.Header.Set("password", password)

	response, err := DoRequest(client, request)

	return response, err
}

func GetOrder(client *http.Client, id string, login string, password string) (*http.Response, error) {
	url := "http://" + address + ":" + port + "/orders/" + id

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, myErrors.ErrorNewRequest
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("login", login)
	request.Header.Set("password", password)

	response, err := DoRequest(client, request)

	return response, err
}

func GetByOrder(client *http.Client, id string, login string, password string) (*http.Response, error) {
	url := "http://" + address + ":" + port + "/elems"
	params := fmt.Sprintf("{\"id\": \"%s\"}", id)
	jsonStr := []byte(params)
	request, err := http.NewRequest("GET", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, myErrors.ErrorNewRequest
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("login", login)
	request.Header.Set("password", password)

	response, err := DoRequest(client, request)

	return response, err
}

func GetWine(client *http.Client, id string, login string, password string) (*http.Response, error) {
	url := "http://" + address + ":" + port + "/wines/" + id

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, myErrors.ErrorNewRequest
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("login", login)
	request.Header.Set("password", password)

	response, err := DoRequest(client, request)

	return response, err
}

func PlaceOrder(client *http.Client, order *openapi.Order, login string, password string) (*http.Response, error) {
	url := "http://" + address + ":" + port + "/orders"

	params := fmt.Sprintf("{\"id\": \"%s\", \"idUser\": \"%s\", \"totalPrice\": \"%s\", \"isPoints\": \"%s\", \"status\": \"%s\"}", order.Id,
		order.IdUser, order.TotalPrice, order.IsPoints, consts.PlacedOrder)
	jsonStr := []byte(params)

	request, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, myErrors.ErrorNewRequest
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("login", login)
	request.Header.Set("password", password)

	response, err := DoRequest(client, request)

	return response, err
}

func PayBill(client *http.Client, id string, login string, password string) (*http.Response, error) {
	url := "http://" + address + ":" + port + "/bills/" + id

	request, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		return nil, myErrors.ErrorNewRequest
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("login", login)
	request.Header.Set("password", password)

	response, err := DoRequest(client, request)

	return response, err
}

func AddWine(client *http.Client, r *openapi.AddWineRequest, login string, password string) (*http.Response, error) {
	url := "http://" + address + ":" + port + "/wines"
	params := fmt.Sprintf("{\"name\": \"%s\", \"count\": \"%s\", \"year\": %d, \"strength\": %d, \"price\": \"%s\", \"type\": \"%s\"}",
		r.Name, r.Count, r.Year, r.Strength, r.Price, r.Type)
	var jsonStr = []byte(params)

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, myErrors.ErrorNewRequest
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("login", login)
	request.Header.Set("password", password)

	response, err := DoRequest(client, request)

	return response, err
}

func DeleteWine(client *http.Client, id string, login string, password string) (*http.Response, error) {
	url := "http://" + address + ":" + port + "/wines/" + id

	request, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, myErrors.ErrorNewRequest
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("login", login)
	request.Header.Set("password", password)

	response, err := DoRequest(client, request)

	return response, err
}

func UpdateWine(client *http.Client, r *openapi.Wine, login string, password string) (*http.Response, error) {
	url := "http://" + address + ":" + port + "/wines"
	params := fmt.Sprintf("{\"id\": \"%s\", \"name\": \"%s\", \"count\": \"%s\", \"year\": %d, \"strength\": %d, \"price\": \"%s\", \"type\": \"%s\"}",
		r.Id, r.Name, r.Count, r.Year, r.Strength, r.Price, r.Type)
	var jsonStr = []byte(params)

	request, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, myErrors.ErrorNewRequest
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("login", login)
	request.Header.Set("password", password)

	response, err := DoRequest(client, request)

	return response, err
}

func CreateUserWine(client *http.Client, userID string, wineID string, login string, password string) (*http.Response, error) {
	url := "http://" + address + ":" + port + "/favourite"
	params := fmt.Sprintf("{\"idUser\": \"%s\", \"idWine\": \"%s\"}", userID, wineID)
	var jsonStr = []byte(params)

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, myErrors.ErrorNewRequest
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("login", login)
	request.Header.Set("password", password)

	response, err := DoRequest(client, request)

	return response, err
}

func DeleteUserWine(client *http.Client, userID string, wineID string, login string, password string) (*http.Response, error) {
	url := "http://" + address + ":" + port + "/favourite"
	params := fmt.Sprintf("{\"idUser\": \"%s\", \"idWine\": \"%s\"}", userID, wineID)
	var jsonStr = []byte(params)

	request, err := http.NewRequest("DELETE", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, myErrors.ErrorNewRequest
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("login", login)
	request.Header.Set("password", password)

	response, err := DoRequest(client, request)

	return response, err
}

func GetUserWines(client *http.Client, userID string, login string, password string) (*http.Response, error) {
	url := "http://" + address + ":" + port + "/favourite/" + userID

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, myErrors.ErrorNewRequest
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("login", login)
	request.Header.Set("password", password)

	response, err := DoRequest(client, request)

	return response, err
}
