/*
 * OTUS Highload Architect
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type DialogUserIdSendPost500Response struct {

	// Описание ошибки
	Message string `json:"message"`

	// Идентификатор запроса. Предназначен для более быстрого поиска проблем.
	RequestId string `json:"request_id,omitempty"`

	// Код ошибки. Предназначен для классификации проблем и более быстрого решения проблем.
	Code int32 `json:"code,omitempty"`
}

// AssertDialogUserIdSendPost500ResponseRequired checks if the required fields are not zero-ed
func AssertDialogUserIdSendPost500ResponseRequired(obj DialogUserIdSendPost500Response) error {
	elements := map[string]interface{}{
		"message": obj.Message,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseDialogUserIdSendPost500ResponseRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of DialogUserIdSendPost500Response (e.g. [][]DialogUserIdSendPost500Response), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseDialogUserIdSendPost500ResponseRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aDialogUserIdSendPost500Response, ok := obj.(DialogUserIdSendPost500Response)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertDialogUserIdSendPost500ResponseRequired(aDialogUserIdSendPost500Response)
	})
}
