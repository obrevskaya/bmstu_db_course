package utils

import (
	"console/errors"
	openapi "console/internal/client"
	"encoding/json"
	"io"
	"net/http"
)

func ParseWinesBody(response *http.Response) ([]openapi.Wine, error) {
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, errors.ErrorReadBody
	}

	var result openapi.GetWinesResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, errors.ErrorParseBody
	}

	return result.Wines, nil
}

func ParseElemsBody(response *http.Response) ([]openapi.Elem, error) {
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, errors.ErrorReadBody
	}

	var result openapi.Elems
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, errors.ErrorParseBody
	}

	return result.Elems, nil
}

func ParseUserWinesBody(response *http.Response) ([]openapi.UserWine, error) {
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, errors.ErrorReadBody
	}

	var result openapi.UserWines
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, errors.ErrorParseBody
	}

	return result.UserWines, nil
}

func ParseOrderBody(response *http.Response) (*openapi.Order, error) {
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, errors.ErrorReadBody
	}

	var result openapi.Order
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, errors.ErrorParseBody
	}

	return &result, nil
}

func ParseWineBody(response *http.Response) (*openapi.Wine, error) {
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, errors.ErrorReadBody
	}

	var result openapi.Wine
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, errors.ErrorParseBody
	}

	return &result, nil
}

func ParseUserBody(response *http.Response) (*openapi.User, error) {
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, errors.ErrorReadBody
	}

	var result openapi.User
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, errors.ErrorParseBody
	}

	return &result, nil
}
