/*
 * OTUS Highload Architect
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type LoginPostRequest struct {

	// Идентификатор пользователя
	Id string `json:"id,omitempty"`

	Password string `json:"password,omitempty"`
}

// AssertLoginPostRequestRequired checks if the required fields are not zero-ed
func AssertLoginPostRequestRequired(obj LoginPostRequest) error {
	return nil
}

// AssertRecurseLoginPostRequestRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of LoginPostRequest (e.g. [][]LoginPostRequest), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseLoginPostRequestRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aLoginPostRequest, ok := obj.(LoginPostRequest)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertLoginPostRequestRequired(aLoginPostRequest)
	})
}
