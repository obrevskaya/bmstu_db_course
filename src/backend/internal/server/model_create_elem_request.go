/*
 * API for ppo project
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type CreateElemRequest struct {
	IdWine string `json:"idWine"`

	Count int32 `json:"count"`
}

// AssertCreateElemRequestRequired checks if the required fields are not zero-ed
func AssertCreateElemRequestRequired(obj CreateElemRequest) error {
	elements := map[string]interface{}{
		"idWine": obj.IdWine,
		"count":  obj.Count,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertCreateElemRequestConstraints checks if the values respects the defined constraints
func AssertCreateElemRequestConstraints(obj CreateElemRequest) error {
	return nil
}
