/*
Delphix DCT API

Delphix DCT API

API version: 1.0
Contact: support@delphix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package delphix_dct_api

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Linger please
var (
	_ context.Context
)

// VDBGroupsApiService VDBGroupsApi service
type VDBGroupsApiService service

type ApiCreateVdbGroupRequest struct {
	ctx context.Context
	ApiService *VDBGroupsApiService
	createVDBGroupRequest *CreateVDBGroupRequest
}

func (r ApiCreateVdbGroupRequest) CreateVDBGroupRequest(createVDBGroupRequest CreateVDBGroupRequest) ApiCreateVdbGroupRequest {
	r.createVDBGroupRequest = &createVDBGroupRequest
	return r
}

func (r ApiCreateVdbGroupRequest) Execute() (*CreateVDBGroupResponse, *http.Response, error) {
	return r.ApiService.CreateVdbGroupExecute(r)
}

/*
CreateVdbGroup Create a new VDBGroup.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateVdbGroupRequest
*/
func (a *VDBGroupsApiService) CreateVdbGroup(ctx context.Context) ApiCreateVdbGroupRequest {
	return ApiCreateVdbGroupRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CreateVDBGroupResponse
func (a *VDBGroupsApiService) CreateVdbGroupExecute(r ApiCreateVdbGroupRequest) (*CreateVDBGroupResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CreateVDBGroupResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VDBGroupsApiService.CreateVdbGroup")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/vdb-groups"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createVDBGroupRequest == nil {
		return localVarReturnValue, nil, reportError("createVDBGroupRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.createVDBGroupRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiDeleteVdbGroupRequest struct {
	ctx context.Context
	ApiService *VDBGroupsApiService
	vdbGroupId string
}


func (r ApiDeleteVdbGroupRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteVdbGroupExecute(r)
}

/*
DeleteVdbGroup Delete a VDBGoup.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param vdbGroupId The ID or name of the VDBGroup.
 @return ApiDeleteVdbGroupRequest
*/
func (a *VDBGroupsApiService) DeleteVdbGroup(ctx context.Context, vdbGroupId string) ApiDeleteVdbGroupRequest {
	return ApiDeleteVdbGroupRequest{
		ApiService: a,
		ctx: ctx,
		vdbGroupId: vdbGroupId,
	}
}

// Execute executes the request
func (a *VDBGroupsApiService) DeleteVdbGroupExecute(r ApiDeleteVdbGroupRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VDBGroupsApiService.DeleteVdbGroup")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/vdb-groups/{vdbGroupId}"
	localVarPath = strings.Replace(localVarPath, "{"+"vdbGroupId"+"}", url.PathEscape(parameterToString(r.vdbGroupId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.vdbGroupId) < 1 {
		return nil, reportError("vdbGroupId must have at least 1 elements")
	}
	if strlen(r.vdbGroupId) > 256 {
		return nil, reportError("vdbGroupId must have less than 256 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetBookmarksByVdbGroupRequest struct {
	ctx context.Context
	ApiService *VDBGroupsApiService
	vdbGroupId string
}


func (r ApiGetBookmarksByVdbGroupRequest) Execute() (*ListBookmarksByVDBGroupsResponse, *http.Response, error) {
	return r.ApiService.GetBookmarksByVdbGroupExecute(r)
}

/*
GetBookmarksByVdbGroup List bookmarks compatible with this VDB Group.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param vdbGroupId The ID or name of the VDBGroup.
 @return ApiGetBookmarksByVdbGroupRequest
*/
func (a *VDBGroupsApiService) GetBookmarksByVdbGroup(ctx context.Context, vdbGroupId string) ApiGetBookmarksByVdbGroupRequest {
	return ApiGetBookmarksByVdbGroupRequest{
		ApiService: a,
		ctx: ctx,
		vdbGroupId: vdbGroupId,
	}
}

// Execute executes the request
//  @return ListBookmarksByVDBGroupsResponse
func (a *VDBGroupsApiService) GetBookmarksByVdbGroupExecute(r ApiGetBookmarksByVdbGroupRequest) (*ListBookmarksByVDBGroupsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListBookmarksByVDBGroupsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VDBGroupsApiService.GetBookmarksByVdbGroup")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/vdb-groups/{vdbGroupId}/bookmarks"
	localVarPath = strings.Replace(localVarPath, "{"+"vdbGroupId"+"}", url.PathEscape(parameterToString(r.vdbGroupId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.vdbGroupId) < 1 {
		return localVarReturnValue, nil, reportError("vdbGroupId must have at least 1 elements")
	}
	if strlen(r.vdbGroupId) > 256 {
		return localVarReturnValue, nil, reportError("vdbGroupId must have less than 256 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetVdbGroupRequest struct {
	ctx context.Context
	ApiService *VDBGroupsApiService
	vdbGroupId string
}


func (r ApiGetVdbGroupRequest) Execute() (*VDBGroup, *http.Response, error) {
	return r.ApiService.GetVdbGroupExecute(r)
}

/*
GetVdbGroup Get a VDBGroup by name.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param vdbGroupId The ID or name of the VDBGroup.
 @return ApiGetVdbGroupRequest
*/
func (a *VDBGroupsApiService) GetVdbGroup(ctx context.Context, vdbGroupId string) ApiGetVdbGroupRequest {
	return ApiGetVdbGroupRequest{
		ApiService: a,
		ctx: ctx,
		vdbGroupId: vdbGroupId,
	}
}

// Execute executes the request
//  @return VDBGroup
func (a *VDBGroupsApiService) GetVdbGroupExecute(r ApiGetVdbGroupRequest) (*VDBGroup, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *VDBGroup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VDBGroupsApiService.GetVdbGroup")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/vdb-groups/{vdbGroupId}"
	localVarPath = strings.Replace(localVarPath, "{"+"vdbGroupId"+"}", url.PathEscape(parameterToString(r.vdbGroupId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.vdbGroupId) < 1 {
		return localVarReturnValue, nil, reportError("vdbGroupId must have at least 1 elements")
	}
	if strlen(r.vdbGroupId) > 256 {
		return localVarReturnValue, nil, reportError("vdbGroupId must have less than 256 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetVdbGroupsRequest struct {
	ctx context.Context
	ApiService *VDBGroupsApiService
}


func (r ApiGetVdbGroupsRequest) Execute() (*ListVDBGroupsResponse, *http.Response, error) {
	return r.ApiService.GetVdbGroupsExecute(r)
}

/*
GetVdbGroups List all VDBGroups.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetVdbGroupsRequest
*/
func (a *VDBGroupsApiService) GetVdbGroups(ctx context.Context) ApiGetVdbGroupsRequest {
	return ApiGetVdbGroupsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ListVDBGroupsResponse
func (a *VDBGroupsApiService) GetVdbGroupsExecute(r ApiGetVdbGroupsRequest) (*ListVDBGroupsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListVDBGroupsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VDBGroupsApiService.GetVdbGroups")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/vdb-groups"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiRefreshVdbGroupRequest struct {
	ctx context.Context
	ApiService *VDBGroupsApiService
	vdbGroupId string
	refreshVDBGroupParameters *RefreshVDBGroupParameters
}

// The parameters to refresh a VDB.
func (r ApiRefreshVdbGroupRequest) RefreshVDBGroupParameters(refreshVDBGroupParameters RefreshVDBGroupParameters) ApiRefreshVdbGroupRequest {
	r.refreshVDBGroupParameters = &refreshVDBGroupParameters
	return r
}

func (r ApiRefreshVdbGroupRequest) Execute() (*RefreshVDBGroupResponse, *http.Response, error) {
	return r.ApiService.RefreshVdbGroupExecute(r)
}

/*
RefreshVdbGroup Refresh a VDBGroup.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param vdbGroupId The ID or name of the VDBGroup.
 @return ApiRefreshVdbGroupRequest
*/
func (a *VDBGroupsApiService) RefreshVdbGroup(ctx context.Context, vdbGroupId string) ApiRefreshVdbGroupRequest {
	return ApiRefreshVdbGroupRequest{
		ApiService: a,
		ctx: ctx,
		vdbGroupId: vdbGroupId,
	}
}

// Execute executes the request
//  @return RefreshVDBGroupResponse
func (a *VDBGroupsApiService) RefreshVdbGroupExecute(r ApiRefreshVdbGroupRequest) (*RefreshVDBGroupResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RefreshVDBGroupResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VDBGroupsApiService.RefreshVdbGroup")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/vdb-groups/{vdbGroupId}/refresh"
	localVarPath = strings.Replace(localVarPath, "{"+"vdbGroupId"+"}", url.PathEscape(parameterToString(r.vdbGroupId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.vdbGroupId) < 1 {
		return localVarReturnValue, nil, reportError("vdbGroupId must have at least 1 elements")
	}
	if strlen(r.vdbGroupId) > 256 {
		return localVarReturnValue, nil, reportError("vdbGroupId must have less than 256 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.refreshVDBGroupParameters
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiRollbackVdbGroupRequest struct {
	ctx context.Context
	ApiService *VDBGroupsApiService
	vdbGroupId string
	rollbackVDBGroupParameters *RollbackVDBGroupParameters
}

// The parameters to rollback a VDB.
func (r ApiRollbackVdbGroupRequest) RollbackVDBGroupParameters(rollbackVDBGroupParameters RollbackVDBGroupParameters) ApiRollbackVdbGroupRequest {
	r.rollbackVDBGroupParameters = &rollbackVDBGroupParameters
	return r
}

func (r ApiRollbackVdbGroupRequest) Execute() (*RollbackVDBGroupResponse, *http.Response, error) {
	return r.ApiService.RollbackVdbGroupExecute(r)
}

/*
RollbackVdbGroup Rollback a VDBGroup.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param vdbGroupId The ID or name of the VDBGroup.
 @return ApiRollbackVdbGroupRequest
*/
func (a *VDBGroupsApiService) RollbackVdbGroup(ctx context.Context, vdbGroupId string) ApiRollbackVdbGroupRequest {
	return ApiRollbackVdbGroupRequest{
		ApiService: a,
		ctx: ctx,
		vdbGroupId: vdbGroupId,
	}
}

// Execute executes the request
//  @return RollbackVDBGroupResponse
func (a *VDBGroupsApiService) RollbackVdbGroupExecute(r ApiRollbackVdbGroupRequest) (*RollbackVDBGroupResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RollbackVDBGroupResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VDBGroupsApiService.RollbackVdbGroup")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/vdb-groups/{vdbGroupId}/rollback"
	localVarPath = strings.Replace(localVarPath, "{"+"vdbGroupId"+"}", url.PathEscape(parameterToString(r.vdbGroupId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.vdbGroupId) < 1 {
		return localVarReturnValue, nil, reportError("vdbGroupId must have at least 1 elements")
	}
	if strlen(r.vdbGroupId) > 256 {
		return localVarReturnValue, nil, reportError("vdbGroupId must have less than 256 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.rollbackVDBGroupParameters
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiSearchVdbGroupsRequest struct {
	ctx context.Context
	ApiService *VDBGroupsApiService
	searchBody *SearchBody
}

// A request body containing a filter expression. This enables searching for items matching arbitrarily complex conditions. The list of attributes which can be used in filter expressions is available in the x-filterable vendor extension.  # Filter Expression Overview **Note: All keywords are case-insensitive**  ## Comparison Operators | Operator | Description | Example | | --- | --- | --- | | CONTAINS | Substring or membership testing for string and list attributes respectively. | field3 CONTAINS &#39;foobar&#39;, field4 CONTAINS TRUE  | | IN | Tests if field is a member of a list literal. List can contain a maximum of 100 values | field2 IN [&#39;Goku&#39;, &#39;Vegeta&#39;] | | GE | Tests if a field is greater than or equal to a literal value | field1 GE 1.2e-2 | | GT | Tests if a field is greater than a literal value | field1 GT 1.2e-2 | | LE | Tests if a field is less than or equal to a literal value | field1 LE 9000 | | LT | Tests if a field is less than a literal value | field1 LT 9.02 | | NE | Tests if a field is not equal to a literal value | field1 NE 42 | | EQ | Tests if a field is equal to a literal value | field1 EQ 42 |  ## Search Operator The SEARCH operator filters for items which have any filterable attribute that contains the input string as a substring, comparison is done case-insensitively. This is not restricted to attributes with string values. Specifically &#x60;SEARCH &#39;12&#39;&#x60; would match an item with an attribute with an integer value of &#x60;123&#x60;.  ## Logical Operators Ordered by precedence. | Operator | Description | Example | | --- | --- | --- | | NOT | Logical NOT (Right associative) | NOT field1 LE 9000 | | AND | Logical AND (Left Associative) | field1 GT 9000 AND field2 EQ &#39;Goku&#39; | | OR | Logical OR (Left Associative) | field1 GT 9000 OR field2 EQ &#39;Goku&#39; |  ## Grouping Parenthesis &#x60;()&#x60; can be used to override operator precedence.  For example: NOT (field1 LT 1234 AND field2 CONTAINS &#39;foo&#39;)  ## Literal Values | Literal      | Description | Examples | | --- | --- | --- | | Nil | Represents the absence of a value | nil, Nil, nIl, NIL | | Boolean | true/false boolean | true, false, True, False, TRUE, FALSE | | Number | Signed integer and floating point numbers. Also supports scientific notation. | 0, 1, -1, 1.2, 0.35, 1.2e-2, -1.2e+2 | | String | Single or double quoted | \&quot;foo\&quot;, \&quot;bar\&quot;, \&quot;foo bar\&quot;, &#39;foo&#39;, &#39;bar&#39;, &#39;foo bar&#39; | | Datetime | Formatted according to [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339) | 2018-04-27T18:39:26.397237+00:00 | | List | Comma-separated literals wrapped in square brackets | [0], [0, 1], [&#39;foo&#39;, \&quot;bar\&quot;] |  ## Limitations - A maximum of 8 unique identifiers may be used inside a filter expression. 
func (r ApiSearchVdbGroupsRequest) SearchBody(searchBody SearchBody) ApiSearchVdbGroupsRequest {
	r.searchBody = &searchBody
	return r
}

func (r ApiSearchVdbGroupsRequest) Execute() (*SearchVDBGroupResponse, *http.Response, error) {
	return r.ApiService.SearchVdbGroupsExecute(r)
}

/*
SearchVdbGroups Search for VDB Groups.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSearchVdbGroupsRequest
*/
func (a *VDBGroupsApiService) SearchVdbGroups(ctx context.Context) ApiSearchVdbGroupsRequest {
	return ApiSearchVdbGroupsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SearchVDBGroupResponse
func (a *VDBGroupsApiService) SearchVdbGroupsExecute(r ApiSearchVdbGroupsRequest) (*SearchVDBGroupResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SearchVDBGroupResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VDBGroupsApiService.SearchVdbGroups")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/vdb-groups/search"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.searchBody
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUpdateVdbGroupByIdRequest struct {
	ctx context.Context
	ApiService *VDBGroupsApiService
	vdbGroupId string
	updateVDBGroupParameters *UpdateVDBGroupParameters
}

// The new data to update a VDB group.
func (r ApiUpdateVdbGroupByIdRequest) UpdateVDBGroupParameters(updateVDBGroupParameters UpdateVDBGroupParameters) ApiUpdateVdbGroupByIdRequest {
	r.updateVDBGroupParameters = &updateVDBGroupParameters
	return r
}

func (r ApiUpdateVdbGroupByIdRequest) Execute() (*VDBGroup, *http.Response, error) {
	return r.ApiService.UpdateVdbGroupByIdExecute(r)
}

/*
UpdateVdbGroupById Update values of a VDB group.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param vdbGroupId The ID or name of the VDBGroup.
 @return ApiUpdateVdbGroupByIdRequest
*/
func (a *VDBGroupsApiService) UpdateVdbGroupById(ctx context.Context, vdbGroupId string) ApiUpdateVdbGroupByIdRequest {
	return ApiUpdateVdbGroupByIdRequest{
		ApiService: a,
		ctx: ctx,
		vdbGroupId: vdbGroupId,
	}
}

// Execute executes the request
//  @return VDBGroup
func (a *VDBGroupsApiService) UpdateVdbGroupByIdExecute(r ApiUpdateVdbGroupByIdRequest) (*VDBGroup, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *VDBGroup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VDBGroupsApiService.UpdateVdbGroupById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/vdb-groups/{vdbGroupId}"
	localVarPath = strings.Replace(localVarPath, "{"+"vdbGroupId"+"}", url.PathEscape(parameterToString(r.vdbGroupId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.vdbGroupId) < 1 {
		return localVarReturnValue, nil, reportError("vdbGroupId must have at least 1 elements")
	}
	if strlen(r.vdbGroupId) > 256 {
		return localVarReturnValue, nil, reportError("vdbGroupId must have less than 256 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.updateVDBGroupParameters
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
