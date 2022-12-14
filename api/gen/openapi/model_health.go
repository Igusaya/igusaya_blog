/*
 * igusaya_blog
 *
 * 個人用内製Blog
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type Health struct {
	Status string `json:"status,omitempty"`
}

// AssertHealthRequired checks if the required fields are not zero-ed
func AssertHealthRequired(obj Health) error {
	return nil
}

// AssertRecurseHealthRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of Health (e.g. [][]Health), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseHealthRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aHealth, ok := obj.(Health)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertHealthRequired(aHealth)
	})
}
