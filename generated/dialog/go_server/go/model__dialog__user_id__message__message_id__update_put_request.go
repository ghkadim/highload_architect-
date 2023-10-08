/*
 * OTUS Highload Architect
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type DialogUserIdMessageMessageIdUpdatePutRequest struct {

	// Сообщение прочитано
	Read bool `json:"read"`
}

// AssertDialogUserIdMessageMessageIdUpdatePutRequestRequired checks if the required fields are not zero-ed
func AssertDialogUserIdMessageMessageIdUpdatePutRequestRequired(obj DialogUserIdMessageMessageIdUpdatePutRequest) error {
	elements := map[string]interface{}{
		"read": obj.Read,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseDialogUserIdMessageMessageIdUpdatePutRequestRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of DialogUserIdMessageMessageIdUpdatePutRequest (e.g. [][]DialogUserIdMessageMessageIdUpdatePutRequest), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseDialogUserIdMessageMessageIdUpdatePutRequestRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aDialogUserIdMessageMessageIdUpdatePutRequest, ok := obj.(DialogUserIdMessageMessageIdUpdatePutRequest)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertDialogUserIdMessageMessageIdUpdatePutRequestRequired(aDialogUserIdMessageMessageIdUpdatePutRequest)
	})
}
