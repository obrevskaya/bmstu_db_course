/*
 * API for ppo project
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type GetWinesResponse struct {
	Wines []Wine `json:"wines"`
}

// AssertGetWinesResponseRequired checks if the required fields are not zero-ed
func AssertGetWinesResponseRequired(obj GetWinesResponse) error {
	elements := map[string]interface{}{
		"wines": obj.Wines,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	for _, el := range obj.Wines {
		if err := AssertWineRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertGetWinesResponseConstraints checks if the values respects the defined constraints
func AssertGetWinesResponseConstraints(obj GetWinesResponse) error {
	return nil
}
