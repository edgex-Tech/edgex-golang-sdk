/*
Http Gateway

Contains interface documents such as accounts, assets, transactions, etc.

API version: v1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// Class00AppPublicApiAPIService Class00AppPublicApiAPI service
type Class00AppPublicApiAPIService service

type ApiGetAppUpdateRequest struct {
	ctx context.Context
	ApiService *Class00AppPublicApiAPIService
	platform *string
	language *string
}

// 客户端类型. 例如 Android / iOS
func (r ApiGetAppUpdateRequest) Platform(platform string) ApiGetAppUpdateRequest {
	r.platform = &platform
	return r
}

// 客户端请求语言. 例如 zh-CN / en-US
func (r ApiGetAppUpdateRequest) Language(language string) ApiGetAppUpdateRequest {
	r.language = &language
	return r
}

func (r ApiGetAppUpdateRequest) Execute() (*ResultAppUpdate, *http.Response, error) {
	return r.ApiService.GetAppUpdateExecute(r)
}

/*
GetAppUpdate 获取APP升级信息

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetAppUpdateRequest
*/
func (a *Class00AppPublicApiAPIService) GetAppUpdate(ctx context.Context) ApiGetAppUpdateRequest {
	return ApiGetAppUpdateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ResultAppUpdate
func (a *Class00AppPublicApiAPIService) GetAppUpdateExecute(r ApiGetAppUpdateRequest) (*ResultAppUpdate, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ResultAppUpdate
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "Class00AppPublicApiAPIService.GetAppUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/public/app/checkAppUpdate"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.platform != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "platform", r.platform, "form", "")
	}
	if r.language != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "language", r.language, "form", "")
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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
