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


// Class07TransferPrivateApiAPIService Class07TransferPrivateApiAPI service
type Class07TransferPrivateApiAPIService service

type ApiCreateTransferOutRequest struct {
	ctx context.Context
	ApiService *Class07TransferPrivateApiAPIService
	createTransferOutParam *CreateTransferOutParam
}

func (r ApiCreateTransferOutRequest) CreateTransferOutParam(createTransferOutParam CreateTransferOutParam) ApiCreateTransferOutRequest {
	r.createTransferOutParam = &createTransferOutParam
	return r
}

func (r ApiCreateTransferOutRequest) Execute() (*ResultCreateTransferOut, *http.Response, error) {
	return r.ApiService.CreateTransferOutExecute(r)
}

/*
CreateTransferOut 创建委托单

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateTransferOutRequest
*/
func (a *Class07TransferPrivateApiAPIService) CreateTransferOut(ctx context.Context) ApiCreateTransferOutRequest {
	return ApiCreateTransferOutRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ResultCreateTransferOut
func (a *Class07TransferPrivateApiAPIService) CreateTransferOutExecute(r ApiCreateTransferOutRequest) (*ResultCreateTransferOut, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ResultCreateTransferOut
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "Class07TransferPrivateApiAPIService.CreateTransferOut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/private/transfer/createTransferOut"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createTransferOutParam == nil {
		return localVarReturnValue, nil, reportError("createTransferOutParam is required and must be specified")
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
	localVarPostBody = r.createTransferOutParam
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

type ApiGetActiveTransferInRequest struct {
	ctx context.Context
	ApiService *Class07TransferPrivateApiAPIService
	accountId *string
	size *string
	offsetData *string
	filterCoinIdList *string
	filterStatusList *string
	filterTransferReasonList *string
	filterStartCreatedTimeInclusive *string
	filterEndCreatedTimeExclusive *string
}

// 账户id
func (r ApiGetActiveTransferInRequest) AccountId(accountId string) ApiGetActiveTransferInRequest {
	r.accountId = &accountId
	return r
}

// 获取数量。必须大于0且小于等于100
func (r ApiGetActiveTransferInRequest) Size(size string) ApiGetActiveTransferInRequest {
	r.size = &size
	return r
}

// 翻页获取偏移。如果不填或者为空串，则获取第一页
func (r ApiGetActiveTransferInRequest) OffsetData(offsetData string) ApiGetActiveTransferInRequest {
	r.offsetData = &offsetData
	return r
}

// 过滤获取对应抵押品coinId对应的转出单，如果为空则获取所有抵押品coinId的转入单
func (r ApiGetActiveTransferInRequest) FilterCoinIdList(filterCoinIdList string) ApiGetActiveTransferInRequest {
	r.filterCoinIdList = &filterCoinIdList
	return r
}

// 过滤获取指定状态的转入单，不填的话所有状态转出单
func (r ApiGetActiveTransferInRequest) FilterStatusList(filterStatusList string) ApiGetActiveTransferInRequest {
	r.filterStatusList = &filterStatusList
	return r
}

// 过滤获取指定转账原因的转入单，不填的话所有转账原因转出单
func (r ApiGetActiveTransferInRequest) FilterTransferReasonList(filterTransferReasonList string) ApiGetActiveTransferInRequest {
	r.filterTransferReasonList = &filterTransferReasonList
	return r
}

// 过滤获取指定开始时间创建的转出单 (包含)，不填或者为0的话就从最早开始
func (r ApiGetActiveTransferInRequest) FilterStartCreatedTimeInclusive(filterStartCreatedTimeInclusive string) ApiGetActiveTransferInRequest {
	r.filterStartCreatedTimeInclusive = &filterStartCreatedTimeInclusive
	return r
}

// 过滤获取指定结束时间创建的转出单 (不包含)，不填或者为0的话就到最近
func (r ApiGetActiveTransferInRequest) FilterEndCreatedTimeExclusive(filterEndCreatedTimeExclusive string) ApiGetActiveTransferInRequest {
	r.filterEndCreatedTimeExclusive = &filterEndCreatedTimeExclusive
	return r
}

func (r ApiGetActiveTransferInRequest) Execute() (*ResultPageDataTransferIn, *http.Response, error) {
	return r.ApiService.GetActiveTransferInExecute(r)
}

/*
GetActiveTransferIn 翻页获取转入单

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetActiveTransferInRequest
*/
func (a *Class07TransferPrivateApiAPIService) GetActiveTransferIn(ctx context.Context) ApiGetActiveTransferInRequest {
	return ApiGetActiveTransferInRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ResultPageDataTransferIn
func (a *Class07TransferPrivateApiAPIService) GetActiveTransferInExecute(r ApiGetActiveTransferInRequest) (*ResultPageDataTransferIn, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ResultPageDataTransferIn
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "Class07TransferPrivateApiAPIService.GetActiveTransferIn")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/private/transfer/getActiveTransferIn"

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
	if r.filterCoinIdList != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filterCoinIdList", r.filterCoinIdList, "form", "")
	}
	if r.filterStatusList != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filterStatusList", r.filterStatusList, "form", "")
	}
	if r.filterTransferReasonList != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filterTransferReasonList", r.filterTransferReasonList, "form", "")
	}
	if r.filterStartCreatedTimeInclusive != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filterStartCreatedTimeInclusive", r.filterStartCreatedTimeInclusive, "form", "")
	}
	if r.filterEndCreatedTimeExclusive != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filterEndCreatedTimeExclusive", r.filterEndCreatedTimeExclusive, "form", "")
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

type ApiGetActiveTransferOutRequest struct {
	ctx context.Context
	ApiService *Class07TransferPrivateApiAPIService
	accountId *string
	size *string
	offsetData *string
	filterCoinIdList *string
	filterStatusList *string
	filterTransferReasonList *string
	filterStartCreatedTimeInclusive *string
	filterEndCreatedTimeExclusive *string
}

// 账户id
func (r ApiGetActiveTransferOutRequest) AccountId(accountId string) ApiGetActiveTransferOutRequest {
	r.accountId = &accountId
	return r
}

// 获取数量。必须大于0且小于等于100
func (r ApiGetActiveTransferOutRequest) Size(size string) ApiGetActiveTransferOutRequest {
	r.size = &size
	return r
}

// 翻页获取偏移。如果不填或者为空串，则获取第一页
func (r ApiGetActiveTransferOutRequest) OffsetData(offsetData string) ApiGetActiveTransferOutRequest {
	r.offsetData = &offsetData
	return r
}

// 过滤获取对应抵押品coinId对应的转入单，如果为空则获取所有抵押品coinId的转入单
func (r ApiGetActiveTransferOutRequest) FilterCoinIdList(filterCoinIdList string) ApiGetActiveTransferOutRequest {
	r.filterCoinIdList = &filterCoinIdList
	return r
}

// 过滤获取指定状态的转入单，不填的话所有状态转入单
func (r ApiGetActiveTransferOutRequest) FilterStatusList(filterStatusList string) ApiGetActiveTransferOutRequest {
	r.filterStatusList = &filterStatusList
	return r
}

// 过滤获取指定转账原因的转入单，不填的话所有转账原因转入单
func (r ApiGetActiveTransferOutRequest) FilterTransferReasonList(filterTransferReasonList string) ApiGetActiveTransferOutRequest {
	r.filterTransferReasonList = &filterTransferReasonList
	return r
}

// 过滤获取指定开始时间创建的转入单 (包含)，不填或者为0的话就从最早开始
func (r ApiGetActiveTransferOutRequest) FilterStartCreatedTimeInclusive(filterStartCreatedTimeInclusive string) ApiGetActiveTransferOutRequest {
	r.filterStartCreatedTimeInclusive = &filterStartCreatedTimeInclusive
	return r
}

// 过滤获取指定结束时间创建的转入单 (不包含)，不填或者为0的话就到最近
func (r ApiGetActiveTransferOutRequest) FilterEndCreatedTimeExclusive(filterEndCreatedTimeExclusive string) ApiGetActiveTransferOutRequest {
	r.filterEndCreatedTimeExclusive = &filterEndCreatedTimeExclusive
	return r
}

func (r ApiGetActiveTransferOutRequest) Execute() (*ResultPageDataTransferOut, *http.Response, error) {
	return r.ApiService.GetActiveTransferOutExecute(r)
}

/*
GetActiveTransferOut 翻页获取转出单

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetActiveTransferOutRequest
*/
func (a *Class07TransferPrivateApiAPIService) GetActiveTransferOut(ctx context.Context) ApiGetActiveTransferOutRequest {
	return ApiGetActiveTransferOutRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ResultPageDataTransferOut
func (a *Class07TransferPrivateApiAPIService) GetActiveTransferOutExecute(r ApiGetActiveTransferOutRequest) (*ResultPageDataTransferOut, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ResultPageDataTransferOut
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "Class07TransferPrivateApiAPIService.GetActiveTransferOut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/private/transfer/getActiveTransferOut"

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
	if r.filterCoinIdList != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filterCoinIdList", r.filterCoinIdList, "form", "")
	}
	if r.filterStatusList != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filterStatusList", r.filterStatusList, "form", "")
	}
	if r.filterTransferReasonList != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filterTransferReasonList", r.filterTransferReasonList, "form", "")
	}
	if r.filterStartCreatedTimeInclusive != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filterStartCreatedTimeInclusive", r.filterStartCreatedTimeInclusive, "form", "")
	}
	if r.filterEndCreatedTimeExclusive != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filterEndCreatedTimeExclusive", r.filterEndCreatedTimeExclusive, "form", "")
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

type ApiGetTransferInByIdRequest struct {
	ctx context.Context
	ApiService *Class07TransferPrivateApiAPIService
	accountId *string
	transferInIdList *string
}

// 账户id
func (r ApiGetTransferInByIdRequest) AccountId(accountId string) ApiGetTransferInByIdRequest {
	r.accountId = &accountId
	return r
}

// 转入单id
func (r ApiGetTransferInByIdRequest) TransferInIdList(transferInIdList string) ApiGetTransferInByIdRequest {
	r.transferInIdList = &transferInIdList
	return r
}

func (r ApiGetTransferInByIdRequest) Execute() (*ResultListTransferIn, *http.Response, error) {
	return r.ApiService.GetTransferInByIdExecute(r)
}

/*
GetTransferInById 根据充值单id批量获取转入单

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetTransferInByIdRequest
*/
func (a *Class07TransferPrivateApiAPIService) GetTransferInById(ctx context.Context) ApiGetTransferInByIdRequest {
	return ApiGetTransferInByIdRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ResultListTransferIn
func (a *Class07TransferPrivateApiAPIService) GetTransferInByIdExecute(r ApiGetTransferInByIdRequest) (*ResultListTransferIn, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ResultListTransferIn
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "Class07TransferPrivateApiAPIService.GetTransferInById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/private/transfer/getTransferInById"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.accountId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "accountId", r.accountId, "form", "")
	}
	if r.transferInIdList != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "transferInIdList", r.transferInIdList, "form", "")
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

type ApiGetTransferOutByIdRequest struct {
	ctx context.Context
	ApiService *Class07TransferPrivateApiAPIService
	accountId *string
	transferOutIdList *string
}

// 账户id
func (r ApiGetTransferOutByIdRequest) AccountId(accountId string) ApiGetTransferOutByIdRequest {
	r.accountId = &accountId
	return r
}

// 转出单id
func (r ApiGetTransferOutByIdRequest) TransferOutIdList(transferOutIdList string) ApiGetTransferOutByIdRequest {
	r.transferOutIdList = &transferOutIdList
	return r
}

func (r ApiGetTransferOutByIdRequest) Execute() (*ResultListTransferOut, *http.Response, error) {
	return r.ApiService.GetTransferOutByIdExecute(r)
}

/*
GetTransferOutById 根据充值单id批量获取转出单

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetTransferOutByIdRequest
*/
func (a *Class07TransferPrivateApiAPIService) GetTransferOutById(ctx context.Context) ApiGetTransferOutByIdRequest {
	return ApiGetTransferOutByIdRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ResultListTransferOut
func (a *Class07TransferPrivateApiAPIService) GetTransferOutByIdExecute(r ApiGetTransferOutByIdRequest) (*ResultListTransferOut, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ResultListTransferOut
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "Class07TransferPrivateApiAPIService.GetTransferOutById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/private/transfer/getTransferOutById"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.accountId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "accountId", r.accountId, "form", "")
	}
	if r.transferOutIdList != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "transferOutIdList", r.transferOutIdList, "form", "")
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

type ApiGetWithdrawAvailableAmount1Request struct {
	ctx context.Context
	ApiService *Class07TransferPrivateApiAPIService
	accountId *string
	coinId *string
}

// 账户id
func (r ApiGetWithdrawAvailableAmount1Request) AccountId(accountId string) ApiGetWithdrawAvailableAmount1Request {
	r.accountId = &accountId
	return r
}

// 币id
func (r ApiGetWithdrawAvailableAmount1Request) CoinId(coinId string) ApiGetWithdrawAvailableAmount1Request {
	r.coinId = &coinId
	return r
}

func (r ApiGetWithdrawAvailableAmount1Request) Execute() (*ResultGetTransferOutAvailableAmount, *http.Response, error) {
	return r.ApiService.GetWithdrawAvailableAmount1Execute(r)
}

/*
GetWithdrawAvailableAmount1 翻页获取提现单

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetWithdrawAvailableAmount1Request
*/
func (a *Class07TransferPrivateApiAPIService) GetWithdrawAvailableAmount1(ctx context.Context) ApiGetWithdrawAvailableAmount1Request {
	return ApiGetWithdrawAvailableAmount1Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ResultGetTransferOutAvailableAmount
func (a *Class07TransferPrivateApiAPIService) GetWithdrawAvailableAmount1Execute(r ApiGetWithdrawAvailableAmount1Request) (*ResultGetTransferOutAvailableAmount, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ResultGetTransferOutAvailableAmount
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "Class07TransferPrivateApiAPIService.GetWithdrawAvailableAmount1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/private/transfer/getTransferOutAvailableAmount"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.accountId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "accountId", r.accountId, "form", "")
	}
	if r.coinId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "coinId", r.coinId, "form", "")
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
