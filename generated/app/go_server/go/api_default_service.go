/*
 * OTUS Highload Architect
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"context"
	"net/http"
	"errors"
)

// DefaultApiService is a service that implements the logic for the DefaultApiServicer
// This service should implement the business logic for every endpoint for the DefaultApi API.
// Include any external packages or services that will be required by this service.
type DefaultApiService struct {
}

// NewDefaultApiService creates a default api service
func NewDefaultApiService() DefaultApiServicer {
	return &DefaultApiService{}
}

// DialogUserIdListGet - 
func (s *DefaultApiService) DialogUserIdListGet(ctx context.Context, userId string) (ImplResponse, error) {
	// TODO - update DialogUserIdListGet with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, []DialogMessage{}) or use other options such as http.Ok ...
	//return Response(200, []DialogMessage{}), nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(401, {}) or use other options such as http.Ok ...
	//return Response(401, nil),nil

	//TODO: Uncomment the next line to return response Response(500, LoginPost500Response{}) or use other options such as http.Ok ...
	//return Response(500, LoginPost500Response{}), nil

	//TODO: Uncomment the next line to return response Response(503, LoginPost500Response{}) or use other options such as http.Ok ...
	//return Response(503, LoginPost500Response{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("DialogUserIdListGet method not implemented")
}

// DialogUserIdSendPost - 
func (s *DefaultApiService) DialogUserIdSendPost(ctx context.Context, userId string, dialogUserIdSendPostRequest DialogUserIdSendPostRequest) (ImplResponse, error) {
	// TODO - update DialogUserIdSendPost with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, {}) or use other options such as http.Ok ...
	//return Response(200, nil),nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(401, {}) or use other options such as http.Ok ...
	//return Response(401, nil),nil

	//TODO: Uncomment the next line to return response Response(500, LoginPost500Response{}) or use other options such as http.Ok ...
	//return Response(500, LoginPost500Response{}), nil

	//TODO: Uncomment the next line to return response Response(503, LoginPost500Response{}) or use other options such as http.Ok ...
	//return Response(503, LoginPost500Response{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("DialogUserIdSendPost method not implemented")
}

// FriendDeleteUserIdPut - 
func (s *DefaultApiService) FriendDeleteUserIdPut(ctx context.Context, userId string) (ImplResponse, error) {
	// TODO - update FriendDeleteUserIdPut with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, {}) or use other options such as http.Ok ...
	//return Response(200, nil),nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(401, {}) or use other options such as http.Ok ...
	//return Response(401, nil),nil

	//TODO: Uncomment the next line to return response Response(500, LoginPost500Response{}) or use other options such as http.Ok ...
	//return Response(500, LoginPost500Response{}), nil

	//TODO: Uncomment the next line to return response Response(503, LoginPost500Response{}) or use other options such as http.Ok ...
	//return Response(503, LoginPost500Response{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("FriendDeleteUserIdPut method not implemented")
}

// FriendSetUserIdPut - 
func (s *DefaultApiService) FriendSetUserIdPut(ctx context.Context, userId string) (ImplResponse, error) {
	// TODO - update FriendSetUserIdPut with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, {}) or use other options such as http.Ok ...
	//return Response(200, nil),nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(401, {}) or use other options such as http.Ok ...
	//return Response(401, nil),nil

	//TODO: Uncomment the next line to return response Response(500, LoginPost500Response{}) or use other options such as http.Ok ...
	//return Response(500, LoginPost500Response{}), nil

	//TODO: Uncomment the next line to return response Response(503, LoginPost500Response{}) or use other options such as http.Ok ...
	//return Response(503, LoginPost500Response{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("FriendSetUserIdPut method not implemented")
}

// LoginPost - 
func (s *DefaultApiService) LoginPost(ctx context.Context, loginPostRequest LoginPostRequest) (ImplResponse, error) {
	// TODO - update LoginPost with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, LoginPost200Response{}) or use other options such as http.Ok ...
	//return Response(200, LoginPost200Response{}), nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	//return Response(404, nil),nil

	//TODO: Uncomment the next line to return response Response(500, LoginPost500Response{}) or use other options such as http.Ok ...
	//return Response(500, LoginPost500Response{}), nil

	//TODO: Uncomment the next line to return response Response(503, LoginPost500Response{}) or use other options such as http.Ok ...
	//return Response(503, LoginPost500Response{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("LoginPost method not implemented")
}

// PostCreatePost - 
func (s *DefaultApiService) PostCreatePost(ctx context.Context, postCreatePostRequest PostCreatePostRequest) (ImplResponse, error) {
	// TODO - update PostCreatePost with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, string{}) or use other options such as http.Ok ...
	//return Response(200, string{}), nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(401, {}) or use other options such as http.Ok ...
	//return Response(401, nil),nil

	//TODO: Uncomment the next line to return response Response(500, LoginPost500Response{}) or use other options such as http.Ok ...
	//return Response(500, LoginPost500Response{}), nil

	//TODO: Uncomment the next line to return response Response(503, LoginPost500Response{}) or use other options such as http.Ok ...
	//return Response(503, LoginPost500Response{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("PostCreatePost method not implemented")
}

// PostDeleteIdPut - 
func (s *DefaultApiService) PostDeleteIdPut(ctx context.Context, id string) (ImplResponse, error) {
	// TODO - update PostDeleteIdPut with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, {}) or use other options such as http.Ok ...
	//return Response(200, nil),nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(401, {}) or use other options such as http.Ok ...
	//return Response(401, nil),nil

	//TODO: Uncomment the next line to return response Response(500, LoginPost500Response{}) or use other options such as http.Ok ...
	//return Response(500, LoginPost500Response{}), nil

	//TODO: Uncomment the next line to return response Response(503, LoginPost500Response{}) or use other options such as http.Ok ...
	//return Response(503, LoginPost500Response{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("PostDeleteIdPut method not implemented")
}

// PostFeedGet - 
func (s *DefaultApiService) PostFeedGet(ctx context.Context, offset float32, limit float32) (ImplResponse, error) {
	// TODO - update PostFeedGet with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, []Post{}) or use other options such as http.Ok ...
	//return Response(200, []Post{}), nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(401, {}) or use other options such as http.Ok ...
	//return Response(401, nil),nil

	//TODO: Uncomment the next line to return response Response(500, LoginPost500Response{}) or use other options such as http.Ok ...
	//return Response(500, LoginPost500Response{}), nil

	//TODO: Uncomment the next line to return response Response(503, LoginPost500Response{}) or use other options such as http.Ok ...
	//return Response(503, LoginPost500Response{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("PostFeedGet method not implemented")
}

// PostGetIdGet - 
func (s *DefaultApiService) PostGetIdGet(ctx context.Context, id string) (ImplResponse, error) {
	// TODO - update PostGetIdGet with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, Post{}) or use other options such as http.Ok ...
	//return Response(200, Post{}), nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(401, {}) or use other options such as http.Ok ...
	//return Response(401, nil),nil

	//TODO: Uncomment the next line to return response Response(500, LoginPost500Response{}) or use other options such as http.Ok ...
	//return Response(500, LoginPost500Response{}), nil

	//TODO: Uncomment the next line to return response Response(503, LoginPost500Response{}) or use other options such as http.Ok ...
	//return Response(503, LoginPost500Response{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("PostGetIdGet method not implemented")
}

// PostUpdatePut - 
func (s *DefaultApiService) PostUpdatePut(ctx context.Context, postUpdatePutRequest PostUpdatePutRequest) (ImplResponse, error) {
	// TODO - update PostUpdatePut with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, {}) or use other options such as http.Ok ...
	//return Response(200, nil),nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(401, {}) or use other options such as http.Ok ...
	//return Response(401, nil),nil

	//TODO: Uncomment the next line to return response Response(500, LoginPost500Response{}) or use other options such as http.Ok ...
	//return Response(500, LoginPost500Response{}), nil

	//TODO: Uncomment the next line to return response Response(503, LoginPost500Response{}) or use other options such as http.Ok ...
	//return Response(503, LoginPost500Response{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("PostUpdatePut method not implemented")
}

// UserGetIdGet - 
func (s *DefaultApiService) UserGetIdGet(ctx context.Context, id string) (ImplResponse, error) {
	// TODO - update UserGetIdGet with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, User{}) or use other options such as http.Ok ...
	//return Response(200, User{}), nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	//return Response(404, nil),nil

	//TODO: Uncomment the next line to return response Response(500, LoginPost500Response{}) or use other options such as http.Ok ...
	//return Response(500, LoginPost500Response{}), nil

	//TODO: Uncomment the next line to return response Response(503, LoginPost500Response{}) or use other options such as http.Ok ...
	//return Response(503, LoginPost500Response{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("UserGetIdGet method not implemented")
}

// UserRegisterPost - 
func (s *DefaultApiService) UserRegisterPost(ctx context.Context, userRegisterPostRequest UserRegisterPostRequest) (ImplResponse, error) {
	// TODO - update UserRegisterPost with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, UserRegisterPost200Response{}) or use other options such as http.Ok ...
	//return Response(200, UserRegisterPost200Response{}), nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(500, LoginPost500Response{}) or use other options such as http.Ok ...
	//return Response(500, LoginPost500Response{}), nil

	//TODO: Uncomment the next line to return response Response(503, LoginPost500Response{}) or use other options such as http.Ok ...
	//return Response(503, LoginPost500Response{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("UserRegisterPost method not implemented")
}

// UserSearchGet - 
func (s *DefaultApiService) UserSearchGet(ctx context.Context, firstName string, lastName string) (ImplResponse, error) {
	// TODO - update UserSearchGet with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, []User{}) or use other options such as http.Ok ...
	//return Response(200, []User{}), nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(500, LoginPost500Response{}) or use other options such as http.Ok ...
	//return Response(500, LoginPost500Response{}), nil

	//TODO: Uncomment the next line to return response Response(503, LoginPost500Response{}) or use other options such as http.Ok ...
	//return Response(503, LoginPost500Response{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("UserSearchGet method not implemented")
}
