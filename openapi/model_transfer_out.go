/*
Http Gateway

Contains interface documents such as accounts, assets, transactions, etc.

API version: v1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the TransferOut type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransferOut{}

// TransferOut 转账转出单
type TransferOut struct {
	// 转账转出单id
	Id *string `json:"id,omitempty"`
	// 所属用户id
	UserId *string `json:"userId,omitempty"`
	// 账户id
	AccountId *string `json:"accountId,omitempty"`
	// 所属抵押品币种id
	CoinId *string `json:"coinId,omitempty"`
	// 转账数量
	Amount *string `json:"amount,omitempty"`
	// 收款账户id
	ReceiverAccountId *string `json:"receiverAccountId,omitempty"`
	// 收款账户L2 key. bigint for hex str
	ReceiverL2Key *string `json:"receiverL2Key,omitempty"`
	// 客户自定义id，用于幂等校验和生成签名nonce
	ClientTransferId *string `json:"clientTransferId,omitempty"`
	// 是否为条件转账
	IsConditionTransfer *bool `json:"isConditionTransfer,omitempty"`
	// 条件转账 fact注册合约地址，当is_conditional_transfer=true时必填。
	ConditionFactRegistryAddress *string `json:"conditionFactRegistryAddress,omitempty"`
	// 条件转账fact生成所使用的erc20地址，当is_conditional_transfer=true时必填。
	ConditionFactErc20Address *string `json:"conditionFactErc20Address,omitempty"`
	// 条件转账fact生成所使用的amount，当is_conditional_transfer=true时必填。
	ConditionFactAmount *string `json:"conditionFactAmount,omitempty"`
	// 条件转账 fact，当is_conditional_transfer=true时必填。
	ConditionFact *string `json:"conditionFact,omitempty"`
	// 转账原因
	TransferReason *string `json:"transferReason,omitempty"`
	// l2签名nonce。取sha256(client_transfer_id) 前32个bit
	L2Nonce *string `json:"l2Nonce,omitempty"`
	// l2签名过期时间，单位毫秒。参与签名生成/校验时要取小时数，即 l2_expire_time / 3600000
	L2ExpireTime *string `json:"l2ExpireTime,omitempty"`
	L2Signature *L2Signature `json:"l2Signature,omitempty"`
	// 附加类型，供上层业务使用
	ExtraType *string `json:"extraType,omitempty"`
	// 额外数据，json格式，默认为空串
	ExtraDataJson *string `json:"extraDataJson,omitempty"`
	// 转账状态
	Status *string `json:"status,omitempty"`
	// 收款方转账转入单id
	ReceiverTransferInId *string `json:"receiverTransferInId,omitempty"`
	// 关联的抵押品明细id。当 status=SUCCESS_XXX/FAILED_L2_REJECT/FAILED_L2_REJECT_APPROVED 时存在
	CollateralTransactionId *string `json:"collateralTransactionId,omitempty"`
	// 审查处理序号。当 status=SUCCESS_XXX/FAILED_CENSOR_FAILURE/FAILED_L2_REJECT/FAILED_L2_REJECT_APPROVED 时存在
	CensorTxId *string `json:"censorTxId,omitempty"`
	// 审查处理时间。当 status=SUCCESS_XXX/FAILED_CENSOR_FAILURE/FAILED_L2_REJECT/FAILED_L2_REJECT_APPROVED 时存在
	CensorTime *string `json:"censorTime,omitempty"`
	// 审查失败错误码。当 status=FAILED_CENSOR_FAILURE 时存在
	CensorFailCode *string `json:"censorFailCode,omitempty"`
	// 审查失败原因。当 status=FAILED_CENSOR_FAILURE 时存在
	CensorFailReason *string `json:"censorFailReason,omitempty"`
	// l2推送交易id。当 censor_status=CENSOR_SUCCESS/L2_APPROVED/L2_REJECT/L2_REJECT_APPROVED 时存在
	L2TxId *string `json:"l2TxId,omitempty"`
	// l2拒绝时间。当 censor_status=L2_REJECT/L2_REJECT_APPROVED 时存在
	L2RejectTime *string `json:"l2RejectTime,omitempty"`
	// l2拒绝错误码。当 censor_status=L2_REJECT/L2_REJECT_APPROVED 时存在
	L2RejectCode *string `json:"l2RejectCode,omitempty"`
	// l2拒绝原因。当 censor_status=L2_REJECT/L2_REJECT_APPROVED 时存在
	L2RejectReason *string `json:"l2RejectReason,omitempty"`
	// l2批次验证时间。当 status=L2_APPROVED/L2_REJECT_APPROVED 时存在
	L2ApprovedTime *string `json:"l2ApprovedTime,omitempty"`
	// 创建时间
	CreatedTime *string `json:"createdTime,omitempty"`
	// 更新时间
	UpdatedTime *string `json:"updatedTime,omitempty"`
}

// NewTransferOut instantiates a new TransferOut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferOut() *TransferOut {
	this := TransferOut{}
	return &this
}

// NewTransferOutWithDefaults instantiates a new TransferOut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferOutWithDefaults() *TransferOut {
	this := TransferOut{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TransferOut) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferOut) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TransferOut) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TransferOut) SetId(v string) {
	o.Id = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *TransferOut) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferOut) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *TransferOut) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *TransferOut) SetUserId(v string) {
	o.UserId = &v
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *TransferOut) GetAccountId() string {
	if o == nil || IsNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferOut) GetAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *TransferOut) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *TransferOut) SetAccountId(v string) {
	o.AccountId = &v
}

// GetCoinId returns the CoinId field value if set, zero value otherwise.
func (o *TransferOut) GetCoinId() string {
	if o == nil || IsNil(o.CoinId) {
		var ret string
		return ret
	}
	return *o.CoinId
}

// GetCoinIdOk returns a tuple with the CoinId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferOut) GetCoinIdOk() (*string, bool) {
	if o == nil || IsNil(o.CoinId) {
		return nil, false
	}
	return o.CoinId, true
}

// HasCoinId returns a boolean if a field has been set.
func (o *TransferOut) HasCoinId() bool {
	if o != nil && !IsNil(o.CoinId) {
		return true
	}

	return false
}

// SetCoinId gets a reference to the given string and assigns it to the CoinId field.
func (o *TransferOut) SetCoinId(v string) {
	o.CoinId = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *TransferOut) GetAmount() string {
	if o == nil || IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferOut) GetAmountOk() (*string, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *TransferOut) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *TransferOut) SetAmount(v string) {
	o.Amount = &v
}

// GetReceiverAccountId returns the ReceiverAccountId field value if set, zero value otherwise.
func (o *TransferOut) GetReceiverAccountId() string {
	if o == nil || IsNil(o.ReceiverAccountId) {
		var ret string
		return ret
	}
	return *o.ReceiverAccountId
}

// GetReceiverAccountIdOk returns a tuple with the ReceiverAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferOut) GetReceiverAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.ReceiverAccountId) {
		return nil, false
	}
	return o.ReceiverAccountId, true
}

// HasReceiverAccountId returns a boolean if a field has been set.
func (o *TransferOut) HasReceiverAccountId() bool {
	if o != nil && !IsNil(o.ReceiverAccountId) {
		return true
	}

	return false
}

// SetReceiverAccountId gets a reference to the given string and assigns it to the ReceiverAccountId field.
func (o *TransferOut) SetReceiverAccountId(v string) {
	o.ReceiverAccountId = &v
}

// GetReceiverL2Key returns the ReceiverL2Key field value if set, zero value otherwise.
func (o *TransferOut) GetReceiverL2Key() string {
	if o == nil || IsNil(o.ReceiverL2Key) {
		var ret string
		return ret
	}
	return *o.ReceiverL2Key
}

// GetReceiverL2KeyOk returns a tuple with the ReceiverL2Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferOut) GetReceiverL2KeyOk() (*string, bool) {
	if o == nil || IsNil(o.ReceiverL2Key) {
		return nil, false
	}
	return o.ReceiverL2Key, true
}

// HasReceiverL2Key returns a boolean if a field has been set.
func (o *TransferOut) HasReceiverL2Key() bool {
	if o != nil && !IsNil(o.ReceiverL2Key) {
		return true
	}

	return false
}

// SetReceiverL2Key gets a reference to the given string and assigns it to the ReceiverL2Key field.
func (o *TransferOut) SetReceiverL2Key(v string) {
	o.ReceiverL2Key = &v
}

// GetClientTransferId returns the ClientTransferId field value if set, zero value otherwise.
func (o *TransferOut) GetClientTransferId() string {
	if o == nil || IsNil(o.ClientTransferId) {
		var ret string
		return ret
	}
	return *o.ClientTransferId
}

// GetClientTransferIdOk returns a tuple with the ClientTransferId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferOut) GetClientTransferIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientTransferId) {
		return nil, false
	}
	return o.ClientTransferId, true
}

// HasClientTransferId returns a boolean if a field has been set.
func (o *TransferOut) HasClientTransferId() bool {
	if o != nil && !IsNil(o.ClientTransferId) {
		return true
	}

	return false
}

// SetClientTransferId gets a reference to the given string and assigns it to the ClientTransferId field.
func (o *TransferOut) SetClientTransferId(v string) {
	o.ClientTransferId = &v
}

// GetIsConditionTransfer returns the IsConditionTransfer field value if set, zero value otherwise.
func (o *TransferOut) GetIsConditionTransfer() bool {
	if o == nil || IsNil(o.IsConditionTransfer) {
		var ret bool
		return ret
	}
	return *o.IsConditionTransfer
}

// GetIsConditionTransferOk returns a tuple with the IsConditionTransfer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferOut) GetIsConditionTransferOk() (*bool, bool) {
	if o == nil || IsNil(o.IsConditionTransfer) {
		return nil, false
	}
	return o.IsConditionTransfer, true
}

// HasIsConditionTransfer returns a boolean if a field has been set.
func (o *TransferOut) HasIsConditionTransfer() bool {
	if o != nil && !IsNil(o.IsConditionTransfer) {
		return true
	}

	return false
}

// SetIsConditionTransfer gets a reference to the given bool and assigns it to the IsConditionTransfer field.
func (o *TransferOut) SetIsConditionTransfer(v bool) {
	o.IsConditionTransfer = &v
}

// GetConditionFactRegistryAddress returns the ConditionFactRegistryAddress field value if set, zero value otherwise.
func (o *TransferOut) GetConditionFactRegistryAddress() string {
	if o == nil || IsNil(o.ConditionFactRegistryAddress) {
		var ret string
		return ret
	}
	return *o.ConditionFactRegistryAddress
}

// GetConditionFactRegistryAddressOk returns a tuple with the ConditionFactRegistryAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferOut) GetConditionFactRegistryAddressOk() (*string, bool) {
	if o == nil || IsNil(o.ConditionFactRegistryAddress) {
		return nil, false
	}
	return o.ConditionFactRegistryAddress, true
}

// HasConditionFactRegistryAddress returns a boolean if a field has been set.
func (o *TransferOut) HasConditionFactRegistryAddress() bool {
	if o != nil && !IsNil(o.ConditionFactRegistryAddress) {
		return true
	}

	return false
}

// SetConditionFactRegistryAddress gets a reference to the given string and assigns it to the ConditionFactRegistryAddress field.
func (o *TransferOut) SetConditionFactRegistryAddress(v string) {
	o.ConditionFactRegistryAddress = &v
}

// GetConditionFactErc20Address returns the ConditionFactErc20Address field value if set, zero value otherwise.
func (o *TransferOut) GetConditionFactErc20Address() string {
	if o == nil || IsNil(o.ConditionFactErc20Address) {
		var ret string
		return ret
	}
	return *o.ConditionFactErc20Address
}

// GetConditionFactErc20AddressOk returns a tuple with the ConditionFactErc20Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferOut) GetConditionFactErc20AddressOk() (*string, bool) {
	if o == nil || IsNil(o.ConditionFactErc20Address) {
		return nil, false
	}
	return o.ConditionFactErc20Address, true
}

// HasConditionFactErc20Address returns a boolean if a field has been set.
func (o *TransferOut) HasConditionFactErc20Address() bool {
	if o != nil && !IsNil(o.ConditionFactErc20Address) {
		return true
	}

	return false
}

// SetConditionFactErc20Address gets a reference to the given string and assigns it to the ConditionFactErc20Address field.
func (o *TransferOut) SetConditionFactErc20Address(v string) {
	o.ConditionFactErc20Address = &v
}

// GetConditionFactAmount returns the ConditionFactAmount field value if set, zero value otherwise.
func (o *TransferOut) GetConditionFactAmount() string {
	if o == nil || IsNil(o.ConditionFactAmount) {
		var ret string
		return ret
	}
	return *o.ConditionFactAmount
}

// GetConditionFactAmountOk returns a tuple with the ConditionFactAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferOut) GetConditionFactAmountOk() (*string, bool) {
	if o == nil || IsNil(o.ConditionFactAmount) {
		return nil, false
	}
	return o.ConditionFactAmount, true
}

// HasConditionFactAmount returns a boolean if a field has been set.
func (o *TransferOut) HasConditionFactAmount() bool {
	if o != nil && !IsNil(o.ConditionFactAmount) {
		return true
	}

	return false
}

// SetConditionFactAmount gets a reference to the given string and assigns it to the ConditionFactAmount field.
func (o *TransferOut) SetConditionFactAmount(v string) {
	o.ConditionFactAmount = &v
}

// GetConditionFact returns the ConditionFact field value if set, zero value otherwise.
func (o *TransferOut) GetConditionFact() string {
	if o == nil || IsNil(o.ConditionFact) {
		var ret string
		return ret
	}
	return *o.ConditionFact
}

// GetConditionFactOk returns a tuple with the ConditionFact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferOut) GetConditionFactOk() (*string, bool) {
	if o == nil || IsNil(o.ConditionFact) {
		return nil, false
	}
	return o.ConditionFact, true
}

// HasConditionFact returns a boolean if a field has been set.
func (o *TransferOut) HasConditionFact() bool {
	if o != nil && !IsNil(o.ConditionFact) {
		return true
	}

	return false
}

// SetConditionFact gets a reference to the given string and assigns it to the ConditionFact field.
func (o *TransferOut) SetConditionFact(v string) {
	o.ConditionFact = &v
}

// GetTransferReason returns the TransferReason field value if set, zero value otherwise.
func (o *TransferOut) GetTransferReason() string {
	if o == nil || IsNil(o.TransferReason) {
		var ret string
		return ret
	}
	return *o.TransferReason
}

// GetTransferReasonOk returns a tuple with the TransferReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferOut) GetTransferReasonOk() (*string, bool) {
	if o == nil || IsNil(o.TransferReason) {
		return nil, false
	}
	return o.TransferReason, true
}

// HasTransferReason returns a boolean if a field has been set.
func (o *TransferOut) HasTransferReason() bool {
	if o != nil && !IsNil(o.TransferReason) {
		return true
	}

	return false
}

// SetTransferReason gets a reference to the given string and assigns it to the TransferReason field.
func (o *TransferOut) SetTransferReason(v string) {
	o.TransferReason = &v
}

// GetL2Nonce returns the L2Nonce field value if set, zero value otherwise.
func (o *TransferOut) GetL2Nonce() string {
	if o == nil || IsNil(o.L2Nonce) {
		var ret string
		return ret
	}
	return *o.L2Nonce
}

// GetL2NonceOk returns a tuple with the L2Nonce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferOut) GetL2NonceOk() (*string, bool) {
	if o == nil || IsNil(o.L2Nonce) {
		return nil, false
	}
	return o.L2Nonce, true
}

// HasL2Nonce returns a boolean if a field has been set.
func (o *TransferOut) HasL2Nonce() bool {
	if o != nil && !IsNil(o.L2Nonce) {
		return true
	}

	return false
}

// SetL2Nonce gets a reference to the given string and assigns it to the L2Nonce field.
func (o *TransferOut) SetL2Nonce(v string) {
	o.L2Nonce = &v
}

// GetL2ExpireTime returns the L2ExpireTime field value if set, zero value otherwise.
func (o *TransferOut) GetL2ExpireTime() string {
	if o == nil || IsNil(o.L2ExpireTime) {
		var ret string
		return ret
	}
	return *o.L2ExpireTime
}

// GetL2ExpireTimeOk returns a tuple with the L2ExpireTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferOut) GetL2ExpireTimeOk() (*string, bool) {
	if o == nil || IsNil(o.L2ExpireTime) {
		return nil, false
	}
	return o.L2ExpireTime, true
}

// HasL2ExpireTime returns a boolean if a field has been set.
func (o *TransferOut) HasL2ExpireTime() bool {
	if o != nil && !IsNil(o.L2ExpireTime) {
		return true
	}

	return false
}

// SetL2ExpireTime gets a reference to the given string and assigns it to the L2ExpireTime field.
func (o *TransferOut) SetL2ExpireTime(v string) {
	o.L2ExpireTime = &v
}

// GetL2Signature returns the L2Signature field value if set, zero value otherwise.
func (o *TransferOut) GetL2Signature() L2Signature {
	if o == nil || IsNil(o.L2Signature) {
		var ret L2Signature
		return ret
	}
	return *o.L2Signature
}

// GetL2SignatureOk returns a tuple with the L2Signature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferOut) GetL2SignatureOk() (*L2Signature, bool) {
	if o == nil || IsNil(o.L2Signature) {
		return nil, false
	}
	return o.L2Signature, true
}

// HasL2Signature returns a boolean if a field has been set.
func (o *TransferOut) HasL2Signature() bool {
	if o != nil && !IsNil(o.L2Signature) {
		return true
	}

	return false
}

// SetL2Signature gets a reference to the given L2Signature and assigns it to the L2Signature field.
func (o *TransferOut) SetL2Signature(v L2Signature) {
	o.L2Signature = &v
}

// GetExtraType returns the ExtraType field value if set, zero value otherwise.
func (o *TransferOut) GetExtraType() string {
	if o == nil || IsNil(o.ExtraType) {
		var ret string
		return ret
	}
	return *o.ExtraType
}

// GetExtraTypeOk returns a tuple with the ExtraType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferOut) GetExtraTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ExtraType) {
		return nil, false
	}
	return o.ExtraType, true
}

// HasExtraType returns a boolean if a field has been set.
func (o *TransferOut) HasExtraType() bool {
	if o != nil && !IsNil(o.ExtraType) {
		return true
	}

	return false
}

// SetExtraType gets a reference to the given string and assigns it to the ExtraType field.
func (o *TransferOut) SetExtraType(v string) {
	o.ExtraType = &v
}

// GetExtraDataJson returns the ExtraDataJson field value if set, zero value otherwise.
func (o *TransferOut) GetExtraDataJson() string {
	if o == nil || IsNil(o.ExtraDataJson) {
		var ret string
		return ret
	}
	return *o.ExtraDataJson
}

// GetExtraDataJsonOk returns a tuple with the ExtraDataJson field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferOut) GetExtraDataJsonOk() (*string, bool) {
	if o == nil || IsNil(o.ExtraDataJson) {
		return nil, false
	}
	return o.ExtraDataJson, true
}

// HasExtraDataJson returns a boolean if a field has been set.
func (o *TransferOut) HasExtraDataJson() bool {
	if o != nil && !IsNil(o.ExtraDataJson) {
		return true
	}

	return false
}

// SetExtraDataJson gets a reference to the given string and assigns it to the ExtraDataJson field.
func (o *TransferOut) SetExtraDataJson(v string) {
	o.ExtraDataJson = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *TransferOut) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferOut) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *TransferOut) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *TransferOut) SetStatus(v string) {
	o.Status = &v
}

// GetReceiverTransferInId returns the ReceiverTransferInId field value if set, zero value otherwise.
func (o *TransferOut) GetReceiverTransferInId() string {
	if o == nil || IsNil(o.ReceiverTransferInId) {
		var ret string
		return ret
	}
	return *o.ReceiverTransferInId
}

// GetReceiverTransferInIdOk returns a tuple with the ReceiverTransferInId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferOut) GetReceiverTransferInIdOk() (*string, bool) {
	if o == nil || IsNil(o.ReceiverTransferInId) {
		return nil, false
	}
	return o.ReceiverTransferInId, true
}

// HasReceiverTransferInId returns a boolean if a field has been set.
func (o *TransferOut) HasReceiverTransferInId() bool {
	if o != nil && !IsNil(o.ReceiverTransferInId) {
		return true
	}

	return false
}

// SetReceiverTransferInId gets a reference to the given string and assigns it to the ReceiverTransferInId field.
func (o *TransferOut) SetReceiverTransferInId(v string) {
	o.ReceiverTransferInId = &v
}

// GetCollateralTransactionId returns the CollateralTransactionId field value if set, zero value otherwise.
func (o *TransferOut) GetCollateralTransactionId() string {
	if o == nil || IsNil(o.CollateralTransactionId) {
		var ret string
		return ret
	}
	return *o.CollateralTransactionId
}

// GetCollateralTransactionIdOk returns a tuple with the CollateralTransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferOut) GetCollateralTransactionIdOk() (*string, bool) {
	if o == nil || IsNil(o.CollateralTransactionId) {
		return nil, false
	}
	return o.CollateralTransactionId, true
}

// HasCollateralTransactionId returns a boolean if a field has been set.
func (o *TransferOut) HasCollateralTransactionId() bool {
	if o != nil && !IsNil(o.CollateralTransactionId) {
		return true
	}

	return false
}

// SetCollateralTransactionId gets a reference to the given string and assigns it to the CollateralTransactionId field.
func (o *TransferOut) SetCollateralTransactionId(v string) {
	o.CollateralTransactionId = &v
}

// GetCensorTxId returns the CensorTxId field value if set, zero value otherwise.
func (o *TransferOut) GetCensorTxId() string {
	if o == nil || IsNil(o.CensorTxId) {
		var ret string
		return ret
	}
	return *o.CensorTxId
}

// GetCensorTxIdOk returns a tuple with the CensorTxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferOut) GetCensorTxIdOk() (*string, bool) {
	if o == nil || IsNil(o.CensorTxId) {
		return nil, false
	}
	return o.CensorTxId, true
}

// HasCensorTxId returns a boolean if a field has been set.
func (o *TransferOut) HasCensorTxId() bool {
	if o != nil && !IsNil(o.CensorTxId) {
		return true
	}

	return false
}

// SetCensorTxId gets a reference to the given string and assigns it to the CensorTxId field.
func (o *TransferOut) SetCensorTxId(v string) {
	o.CensorTxId = &v
}

// GetCensorTime returns the CensorTime field value if set, zero value otherwise.
func (o *TransferOut) GetCensorTime() string {
	if o == nil || IsNil(o.CensorTime) {
		var ret string
		return ret
	}
	return *o.CensorTime
}

// GetCensorTimeOk returns a tuple with the CensorTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferOut) GetCensorTimeOk() (*string, bool) {
	if o == nil || IsNil(o.CensorTime) {
		return nil, false
	}
	return o.CensorTime, true
}

// HasCensorTime returns a boolean if a field has been set.
func (o *TransferOut) HasCensorTime() bool {
	if o != nil && !IsNil(o.CensorTime) {
		return true
	}

	return false
}

// SetCensorTime gets a reference to the given string and assigns it to the CensorTime field.
func (o *TransferOut) SetCensorTime(v string) {
	o.CensorTime = &v
}

// GetCensorFailCode returns the CensorFailCode field value if set, zero value otherwise.
func (o *TransferOut) GetCensorFailCode() string {
	if o == nil || IsNil(o.CensorFailCode) {
		var ret string
		return ret
	}
	return *o.CensorFailCode
}

// GetCensorFailCodeOk returns a tuple with the CensorFailCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferOut) GetCensorFailCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CensorFailCode) {
		return nil, false
	}
	return o.CensorFailCode, true
}

// HasCensorFailCode returns a boolean if a field has been set.
func (o *TransferOut) HasCensorFailCode() bool {
	if o != nil && !IsNil(o.CensorFailCode) {
		return true
	}

	return false
}

// SetCensorFailCode gets a reference to the given string and assigns it to the CensorFailCode field.
func (o *TransferOut) SetCensorFailCode(v string) {
	o.CensorFailCode = &v
}

// GetCensorFailReason returns the CensorFailReason field value if set, zero value otherwise.
func (o *TransferOut) GetCensorFailReason() string {
	if o == nil || IsNil(o.CensorFailReason) {
		var ret string
		return ret
	}
	return *o.CensorFailReason
}

// GetCensorFailReasonOk returns a tuple with the CensorFailReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferOut) GetCensorFailReasonOk() (*string, bool) {
	if o == nil || IsNil(o.CensorFailReason) {
		return nil, false
	}
	return o.CensorFailReason, true
}

// HasCensorFailReason returns a boolean if a field has been set.
func (o *TransferOut) HasCensorFailReason() bool {
	if o != nil && !IsNil(o.CensorFailReason) {
		return true
	}

	return false
}

// SetCensorFailReason gets a reference to the given string and assigns it to the CensorFailReason field.
func (o *TransferOut) SetCensorFailReason(v string) {
	o.CensorFailReason = &v
}

// GetL2TxId returns the L2TxId field value if set, zero value otherwise.
func (o *TransferOut) GetL2TxId() string {
	if o == nil || IsNil(o.L2TxId) {
		var ret string
		return ret
	}
	return *o.L2TxId
}

// GetL2TxIdOk returns a tuple with the L2TxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferOut) GetL2TxIdOk() (*string, bool) {
	if o == nil || IsNil(o.L2TxId) {
		return nil, false
	}
	return o.L2TxId, true
}

// HasL2TxId returns a boolean if a field has been set.
func (o *TransferOut) HasL2TxId() bool {
	if o != nil && !IsNil(o.L2TxId) {
		return true
	}

	return false
}

// SetL2TxId gets a reference to the given string and assigns it to the L2TxId field.
func (o *TransferOut) SetL2TxId(v string) {
	o.L2TxId = &v
}

// GetL2RejectTime returns the L2RejectTime field value if set, zero value otherwise.
func (o *TransferOut) GetL2RejectTime() string {
	if o == nil || IsNil(o.L2RejectTime) {
		var ret string
		return ret
	}
	return *o.L2RejectTime
}

// GetL2RejectTimeOk returns a tuple with the L2RejectTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferOut) GetL2RejectTimeOk() (*string, bool) {
	if o == nil || IsNil(o.L2RejectTime) {
		return nil, false
	}
	return o.L2RejectTime, true
}

// HasL2RejectTime returns a boolean if a field has been set.
func (o *TransferOut) HasL2RejectTime() bool {
	if o != nil && !IsNil(o.L2RejectTime) {
		return true
	}

	return false
}

// SetL2RejectTime gets a reference to the given string and assigns it to the L2RejectTime field.
func (o *TransferOut) SetL2RejectTime(v string) {
	o.L2RejectTime = &v
}

// GetL2RejectCode returns the L2RejectCode field value if set, zero value otherwise.
func (o *TransferOut) GetL2RejectCode() string {
	if o == nil || IsNil(o.L2RejectCode) {
		var ret string
		return ret
	}
	return *o.L2RejectCode
}

// GetL2RejectCodeOk returns a tuple with the L2RejectCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferOut) GetL2RejectCodeOk() (*string, bool) {
	if o == nil || IsNil(o.L2RejectCode) {
		return nil, false
	}
	return o.L2RejectCode, true
}

// HasL2RejectCode returns a boolean if a field has been set.
func (o *TransferOut) HasL2RejectCode() bool {
	if o != nil && !IsNil(o.L2RejectCode) {
		return true
	}

	return false
}

// SetL2RejectCode gets a reference to the given string and assigns it to the L2RejectCode field.
func (o *TransferOut) SetL2RejectCode(v string) {
	o.L2RejectCode = &v
}

// GetL2RejectReason returns the L2RejectReason field value if set, zero value otherwise.
func (o *TransferOut) GetL2RejectReason() string {
	if o == nil || IsNil(o.L2RejectReason) {
		var ret string
		return ret
	}
	return *o.L2RejectReason
}

// GetL2RejectReasonOk returns a tuple with the L2RejectReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferOut) GetL2RejectReasonOk() (*string, bool) {
	if o == nil || IsNil(o.L2RejectReason) {
		return nil, false
	}
	return o.L2RejectReason, true
}

// HasL2RejectReason returns a boolean if a field has been set.
func (o *TransferOut) HasL2RejectReason() bool {
	if o != nil && !IsNil(o.L2RejectReason) {
		return true
	}

	return false
}

// SetL2RejectReason gets a reference to the given string and assigns it to the L2RejectReason field.
func (o *TransferOut) SetL2RejectReason(v string) {
	o.L2RejectReason = &v
}

// GetL2ApprovedTime returns the L2ApprovedTime field value if set, zero value otherwise.
func (o *TransferOut) GetL2ApprovedTime() string {
	if o == nil || IsNil(o.L2ApprovedTime) {
		var ret string
		return ret
	}
	return *o.L2ApprovedTime
}

// GetL2ApprovedTimeOk returns a tuple with the L2ApprovedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferOut) GetL2ApprovedTimeOk() (*string, bool) {
	if o == nil || IsNil(o.L2ApprovedTime) {
		return nil, false
	}
	return o.L2ApprovedTime, true
}

// HasL2ApprovedTime returns a boolean if a field has been set.
func (o *TransferOut) HasL2ApprovedTime() bool {
	if o != nil && !IsNil(o.L2ApprovedTime) {
		return true
	}

	return false
}

// SetL2ApprovedTime gets a reference to the given string and assigns it to the L2ApprovedTime field.
func (o *TransferOut) SetL2ApprovedTime(v string) {
	o.L2ApprovedTime = &v
}

// GetCreatedTime returns the CreatedTime field value if set, zero value otherwise.
func (o *TransferOut) GetCreatedTime() string {
	if o == nil || IsNil(o.CreatedTime) {
		var ret string
		return ret
	}
	return *o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferOut) GetCreatedTimeOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedTime) {
		return nil, false
	}
	return o.CreatedTime, true
}

// HasCreatedTime returns a boolean if a field has been set.
func (o *TransferOut) HasCreatedTime() bool {
	if o != nil && !IsNil(o.CreatedTime) {
		return true
	}

	return false
}

// SetCreatedTime gets a reference to the given string and assigns it to the CreatedTime field.
func (o *TransferOut) SetCreatedTime(v string) {
	o.CreatedTime = &v
}

// GetUpdatedTime returns the UpdatedTime field value if set, zero value otherwise.
func (o *TransferOut) GetUpdatedTime() string {
	if o == nil || IsNil(o.UpdatedTime) {
		var ret string
		return ret
	}
	return *o.UpdatedTime
}

// GetUpdatedTimeOk returns a tuple with the UpdatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferOut) GetUpdatedTimeOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedTime) {
		return nil, false
	}
	return o.UpdatedTime, true
}

// HasUpdatedTime returns a boolean if a field has been set.
func (o *TransferOut) HasUpdatedTime() bool {
	if o != nil && !IsNil(o.UpdatedTime) {
		return true
	}

	return false
}

// SetUpdatedTime gets a reference to the given string and assigns it to the UpdatedTime field.
func (o *TransferOut) SetUpdatedTime(v string) {
	o.UpdatedTime = &v
}

func (o TransferOut) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransferOut) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	if !IsNil(o.AccountId) {
		toSerialize["accountId"] = o.AccountId
	}
	if !IsNil(o.CoinId) {
		toSerialize["coinId"] = o.CoinId
	}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.ReceiverAccountId) {
		toSerialize["receiverAccountId"] = o.ReceiverAccountId
	}
	if !IsNil(o.ReceiverL2Key) {
		toSerialize["receiverL2Key"] = o.ReceiverL2Key
	}
	if !IsNil(o.ClientTransferId) {
		toSerialize["clientTransferId"] = o.ClientTransferId
	}
	if !IsNil(o.IsConditionTransfer) {
		toSerialize["isConditionTransfer"] = o.IsConditionTransfer
	}
	if !IsNil(o.ConditionFactRegistryAddress) {
		toSerialize["conditionFactRegistryAddress"] = o.ConditionFactRegistryAddress
	}
	if !IsNil(o.ConditionFactErc20Address) {
		toSerialize["conditionFactErc20Address"] = o.ConditionFactErc20Address
	}
	if !IsNil(o.ConditionFactAmount) {
		toSerialize["conditionFactAmount"] = o.ConditionFactAmount
	}
	if !IsNil(o.ConditionFact) {
		toSerialize["conditionFact"] = o.ConditionFact
	}
	if !IsNil(o.TransferReason) {
		toSerialize["transferReason"] = o.TransferReason
	}
	if !IsNil(o.L2Nonce) {
		toSerialize["l2Nonce"] = o.L2Nonce
	}
	if !IsNil(o.L2ExpireTime) {
		toSerialize["l2ExpireTime"] = o.L2ExpireTime
	}
	if !IsNil(o.L2Signature) {
		toSerialize["l2Signature"] = o.L2Signature
	}
	if !IsNil(o.ExtraType) {
		toSerialize["extraType"] = o.ExtraType
	}
	if !IsNil(o.ExtraDataJson) {
		toSerialize["extraDataJson"] = o.ExtraDataJson
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.ReceiverTransferInId) {
		toSerialize["receiverTransferInId"] = o.ReceiverTransferInId
	}
	if !IsNil(o.CollateralTransactionId) {
		toSerialize["collateralTransactionId"] = o.CollateralTransactionId
	}
	if !IsNil(o.CensorTxId) {
		toSerialize["censorTxId"] = o.CensorTxId
	}
	if !IsNil(o.CensorTime) {
		toSerialize["censorTime"] = o.CensorTime
	}
	if !IsNil(o.CensorFailCode) {
		toSerialize["censorFailCode"] = o.CensorFailCode
	}
	if !IsNil(o.CensorFailReason) {
		toSerialize["censorFailReason"] = o.CensorFailReason
	}
	if !IsNil(o.L2TxId) {
		toSerialize["l2TxId"] = o.L2TxId
	}
	if !IsNil(o.L2RejectTime) {
		toSerialize["l2RejectTime"] = o.L2RejectTime
	}
	if !IsNil(o.L2RejectCode) {
		toSerialize["l2RejectCode"] = o.L2RejectCode
	}
	if !IsNil(o.L2RejectReason) {
		toSerialize["l2RejectReason"] = o.L2RejectReason
	}
	if !IsNil(o.L2ApprovedTime) {
		toSerialize["l2ApprovedTime"] = o.L2ApprovedTime
	}
	if !IsNil(o.CreatedTime) {
		toSerialize["createdTime"] = o.CreatedTime
	}
	if !IsNil(o.UpdatedTime) {
		toSerialize["updatedTime"] = o.UpdatedTime
	}
	return toSerialize, nil
}

type NullableTransferOut struct {
	value *TransferOut
	isSet bool
}

func (v NullableTransferOut) Get() *TransferOut {
	return v.value
}

func (v *NullableTransferOut) Set(val *TransferOut) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferOut) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferOut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferOut(val *TransferOut) *NullableTransferOut {
	return &NullableTransferOut{value: val, isSet: true}
}

func (v NullableTransferOut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferOut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


