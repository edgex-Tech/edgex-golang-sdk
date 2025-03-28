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


// Class05DepositPrivateApiAPIService Class05DepositPrivateApiAPI service
type Class05DepositPrivateApiAPIService service

type ApiCreateDepositRequest struct {
	ctx context.Context
	ApiService *Class05DepositPrivateApiAPIService
	createDepositParam *CreateDepositParam
}

func (r ApiCreateDepositRequest) CreateDepositParam(createDepositParam CreateDepositParam) ApiCreateDepositRequest {
	r.createDepositParam = &createDepositParam
	return r
}

func (r ApiCreateDepositRequest) Execute() (*ResultCreateDeposit, *http.Response, error) {
	return r.ApiService.CreateDepositExecute(r)
}

/*
CreateDeposit 创建充值单

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateDepositRequest
*/
func (a *Class05DepositPrivateApiAPIService) CreateDeposit(ctx context.Context) ApiCreateDepositRequest {
	return ApiCreateDepositRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ResultCreateDeposit
func (a *Class05DepositPrivateApiAPIService) CreateDepositExecute(r ApiCreateDepositRequest) (*ResultCreateDeposit, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ResultCreateDeposit
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "Class05DepositPrivateApiAPIService.CreateDeposit")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/private/deposit/createDeposit"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createDepositParam == nil {
		return localVarReturnValue, nil, reportError("createDepositParam is required and must be specified")
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
	localVarPostBody = r.createDepositParam
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v Result
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
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

type ApiGetActiveDepositRequest struct {
	ctx context.Context
	ApiService *Class05DepositPrivateApiAPIService
	accountId *string
	size *string
	offsetData *string
}

// 账户id
func (r ApiGetActiveDepositRequest) AccountId(accountId string) ApiGetActiveDepositRequest {
	r.accountId = &accountId
	return r
}

// 获取数量。必须大于0且小于等于100
func (r ApiGetActiveDepositRequest) Size(size string) ApiGetActiveDepositRequest {
	r.size = &size
	return r
}

// 翻页获取偏移。如果不填或者为空串，则获取第一页
func (r ApiGetActiveDepositRequest) OffsetData(offsetData string) ApiGetActiveDepositRequest {
	r.offsetData = &offsetData
	return r
}

func (r ApiGetActiveDepositRequest) Execute() (*ResultPageDataDeposit, *http.Response, error) {
	return r.ApiService.GetActiveDepositExecute(r)
}

/*
GetActiveDeposit 翻页获取充值单

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetActiveDepositRequest
*/
func (a *Class05DepositPrivateApiAPIService) GetActiveDeposit(ctx context.Context) ApiGetActiveDepositRequest {
	return ApiGetActiveDepositRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ResultPageDataDeposit
func (a *Class05DepositPrivateApiAPIService) GetActiveDepositExecute(r ApiGetActiveDepositRequest) (*ResultPageDataDeposit, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ResultPageDataDeposit
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "Class05DepositPrivateApiAPIService.GetActiveDeposit")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/private/deposit/getActiveDeposit"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.accountId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "accountId", r.accountId, "form", "")
	}
	if r.size != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "size", r.size, "form", "")
	}
	if r.offsetData != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offsetData", r.offsetData, "form", "")
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v Result
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
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

type ApiGetDepositByClientDepositIdRequest struct {
	ctx context.Context
	ApiService *Class05DepositPrivateApiAPIService
	accountId *string
	clientDepositIdList *string
}

// 账户id
func (r ApiGetDepositByClientDepositIdRequest) AccountId(accountId string) ApiGetDepositByClientDepositIdRequest {
	r.accountId = &accountId
	return r
}

// 客户自定义id
func (r ApiGetDepositByClientDepositIdRequest) ClientDepositIdList(clientDepositIdList string) ApiGetDepositByClientDepositIdRequest {
	r.clientDepositIdList = &clientDepositIdList
	return r
}

func (r ApiGetDepositByClientDepositIdRequest) Execute() (*ResultListDeposit, *http.Response, error) {
	return r.ApiService.GetDepositByClientDepositIdExecute(r)
}

/*
GetDepositByClientDepositId 根据账户id和充值单clientId批量获取充值单

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetDepositByClientDepositIdRequest
*/
func (a *Class05DepositPrivateApiAPIService) GetDepositByClientDepositId(ctx context.Context) ApiGetDepositByClientDepositIdRequest {
	return ApiGetDepositByClientDepositIdRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ResultListDeposit
func (a *Class05DepositPrivateApiAPIService) GetDepositByClientDepositIdExecute(r ApiGetDepositByClientDepositIdRequest) (*ResultListDeposit, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ResultListDeposit
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "Class05DepositPrivateApiAPIService.GetDepositByClientDepositId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/private/deposit/getDepositByClientDepositId"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.accountId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "accountId", r.accountId, "form", "")
	}
	if r.clientDepositIdList != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "clientDepositIdList", r.clientDepositIdList, "form", "")
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v Result
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
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

type ApiGetDepositByIdRequest struct {
	ctx context.Context
	ApiService *Class05DepositPrivateApiAPIService
	accountId *string
	depositIdList *string
}

// 账户id
func (r ApiGetDepositByIdRequest) AccountId(accountId string) ApiGetDepositByIdRequest {
	r.accountId = &accountId
	return r
}

// 充值单id
func (r ApiGetDepositByIdRequest) DepositIdList(depositIdList string) ApiGetDepositByIdRequest {
	r.depositIdList = &depositIdList
	return r
}

func (r ApiGetDepositByIdRequest) Execute() (*ResultListDeposit, *http.Response, error) {
	return r.ApiService.GetDepositByIdExecute(r)
}

/*
GetDepositById 根据充值单id批量获取充值单

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetDepositByIdRequest
*/
func (a *Class05DepositPrivateApiAPIService) GetDepositById(ctx context.Context) ApiGetDepositByIdRequest {
	return ApiGetDepositByIdRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ResultListDeposit
func (a *Class05DepositPrivateApiAPIService) GetDepositByIdExecute(r ApiGetDepositByIdRequest) (*ResultListDeposit, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ResultListDeposit
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "Class05DepositPrivateApiAPIService.GetDepositById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/private/deposit/getDepositById"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.accountId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "accountId", r.accountId, "form", "")
	}
	if r.depositIdList != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "depositIdList", r.depositIdList, "form", "")
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v Result
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
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

type ApiRequestRelayerSignAndBroadcastRequest struct {
	ctx context.Context
	ApiService *Class05DepositPrivateApiAPIService
	requestRelayerSignAndBroadcastParam *RequestRelayerSignAndBroadcastParam
}

func (r ApiRequestRelayerSignAndBroadcastRequest) RequestRelayerSignAndBroadcastParam(requestRelayerSignAndBroadcastParam RequestRelayerSignAndBroadcastParam) ApiRequestRelayerSignAndBroadcastRequest {
	r.requestRelayerSignAndBroadcastParam = &requestRelayerSignAndBroadcastParam
	return r
}

func (r ApiRequestRelayerSignAndBroadcastRequest) Execute() (*ResultResultRequestRelayerSignAndBroadcast, *http.Response, error) {
	return r.ApiService.RequestRelayerSignAndBroadcastExecute(r)
}

/*
RequestRelayerSignAndBroadcast 创建 Relayer 充值单

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiRequestRelayerSignAndBroadcastRequest
*/
func (a *Class05DepositPrivateApiAPIService) RequestRelayerSignAndBroadcast(ctx context.Context) ApiRequestRelayerSignAndBroadcastRequest {
	return ApiRequestRelayerSignAndBroadcastRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ResultResultRequestRelayerSignAndBroadcast
func (a *Class05DepositPrivateApiAPIService) RequestRelayerSignAndBroadcastExecute(r ApiRequestRelayerSignAndBroadcastRequest) (*ResultResultRequestRelayerSignAndBroadcast, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ResultResultRequestRelayerSignAndBroadcast
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "Class05DepositPrivateApiAPIService.RequestRelayerSignAndBroadcast")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/private/deposit/requestRelayerSignAndBroadcast"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.requestRelayerSignAndBroadcastParam == nil {
		return localVarReturnValue, nil, reportError("requestRelayerSignAndBroadcastParam is required and must be specified")
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
	localVarPostBody = r.requestRelayerSignAndBroadcastParam
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v Result
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
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
