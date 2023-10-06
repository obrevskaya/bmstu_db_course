/*
 * API for ppo project
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type RegisterRequest struct {
	Login string `json:"login"`

	Password string `json:"password"`

	Fio string `json:"fio"`

	Email string `json:"email"`

	Points int32 `json:"points,omitempty"`

	Status string `json:"status"`
}

// AssertRegisterRequestRequired checks if the required fields are not zero-ed
func AssertRegisterRequestRequired(obj RegisterRequest) error {
	elements := map[string]interface{}{
		"login":    obj.Login,
		"password": obj.Password,
		"fio":      obj.Fio,
		"email":    obj.Email,
		"status":   obj.Status,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRegisterRequestConstraints checks if the values respects the defined constraints
func AssertRegisterRequestConstraints(obj RegisterRequest) error {
	return nil
}