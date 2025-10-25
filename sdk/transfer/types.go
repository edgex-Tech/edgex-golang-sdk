package transfer

import (
	"time"
)

// TransferReasonType represents the type of transfer reason
type TransferReasonType int

const (
	// USER_TRANSFER represents user initiated transfer
	USER_TRANSFER TransferReasonType = 1
	// FAST_WITHDRAW represents fast withdrawal
	FAST_WITHDRAW TransferReasonType = 2
	// CROSS_DEPOSIT represents cross-chain deposit
	CROSS_DEPOSIT TransferReasonType = 3
	// CROSS_WITHDRAW represents cross-chain withdrawal
	CROSS_WITHDRAW TransferReasonType = 4
)

// String returns the string representation of TransferReasonType
func (t TransferReasonType) String() string {
	switch t {
	case USER_TRANSFER:
		return "USER_TRANSFER"
	case FAST_WITHDRAW:
		return "FAST_WITHDRAW"
	case CROSS_DEPOSIT:
		return "CROSS_DEPOSIT"
	case CROSS_WITHDRAW:
		return "CROSS_WITHDRAW"
	default:
		return "UNKNOWN"
	}
}

// Value returns the integer value of TransferReasonType
func (t TransferReasonType) Value() int {
	return int(t)
}

// TransferOut represents a transfer out record
type TransferOut struct {
	Id                           *string      `json:"id,omitempty"`
	UserId                       *string      `json:"userId,omitempty"`
	AccountId                    *string      `json:"accountId,omitempty"`
	CoinId                       *string      `json:"coinId,omitempty"`
	Amount                       *string      `json:"amount,omitempty"`
	ReceiverAccountId            *string      `json:"receiverAccountId,omitempty"`
	ReceiverL2Key                *string      `json:"receiverL2Key,omitempty"`
	ClientTransferId             *string      `json:"clientTransferId,omitempty"`
	IsConditionTransfer          *bool        `json:"isConditionTransfer,omitempty"`
	ConditionFactRegistryAddress *string      `json:"conditionFactRegistryAddress,omitempty"`
	ConditionFactErc20Address    *string      `json:"conditionFactErc20Address,omitempty"`
	ConditionFactAmount          *string      `json:"conditionFactAmount,omitempty"`
	ConditionFact                *string      `json:"conditionFact,omitempty"`
	TransferReason               *string      `json:"transferReason,omitempty"`
	L2Nonce                      *string      `json:"l2Nonce,omitempty"`
	L2ExpireTime                 *string      `json:"l2ExpireTime,omitempty"`
	L2Signature                  *L2Signature `json:"l2Signature,omitempty"`
	ExtraType                    *string      `json:"extraType,omitempty"`
	ExtraDataJson                *string      `json:"extraDataJson,omitempty"`
	Status                       *string      `json:"status,omitempty"`
	ReceiverTransferInId         *string      `json:"receiverTransferInId,omitempty"`
	CollateralTransactionId      *string      `json:"collateralTransactionId,omitempty"`
	CensorTxId                   *string      `json:"censorTxId,omitempty"`
	CensorTime                   *string      `json:"censorTime,omitempty"`
	CensorFailCode               *string      `json:"censorFailCode,omitempty"`
	CensorFailReason             *string      `json:"censorFailReason,omitempty"`
	L2TxId                       *string      `json:"l2TxId,omitempty"`
	L2RejectTime                 *string      `json:"l2RejectTime,omitempty"`
	L2RejectCode                 *string      `json:"l2RejectCode,omitempty"`
	L2RejectReason               *string      `json:"l2RejectReason,omitempty"`
	L2ApprovedTime               *string      `json:"l2ApprovedTime,omitempty"`
	CreatedTime                  *string      `json:"createdTime,omitempty"`
	UpdatedTime                  *string      `json:"updatedTime,omitempty"`
}

// TransferIn represents a transfer in record
type TransferIn struct {
	Id                           *string `json:"id,omitempty"`
	UserId                       *string `json:"userId,omitempty"`
	AccountId                    *string `json:"accountId,omitempty"`
	CoinId                       *string `json:"coinId,omitempty"`
	Amount                       *string `json:"amount,omitempty"`
	SenderAccountId              *string `json:"senderAccountId,omitempty"`
	SenderL2Key                  *string `json:"senderL2Key,omitempty"`
	ClientTransferId             *string `json:"clientTransferId,omitempty"`
	IsConditionTransfer          *bool   `json:"isConditionTransfer,omitempty"`
	ConditionFactRegistryAddress *string `json:"conditionFactRegistryAddress,omitempty"`
	ConditionFactErc20Address    *string `json:"conditionFactErc20Address,omitempty"`
	ConditionFactAmount          *string `json:"conditionFactAmount,omitempty"`
	ConditionFact                *string `json:"conditionFact,omitempty"`
	TransferOutId                *string `json:"transferOutId,omitempty"`
	Status                       *string `json:"status,omitempty"`
	CollateralTransactionId      *string `json:"collateralTransactionId,omitempty"`
	CensorTxId                   *string `json:"censorTxId,omitempty"`
	CensorTime                   *string `json:"censorTime,omitempty"`
	CensorFailCode               *string `json:"censorFailCode,omitempty"`
	CensorFailReason             *string `json:"censorFailReason,omitempty"`
	L2TxId                       *string `json:"l2TxId,omitempty"`
	L2RejectTime                 *string `json:"l2RejectTime,omitempty"`
	L2RejectCode                 *string `json:"l2RejectCode,omitempty"`
	L2RejectReason               *string `json:"l2RejectReason,omitempty"`
	L2ApprovedTime               *string `json:"l2ApprovedTime,omitempty"`
	CreatedTime                  *string `json:"createdTime,omitempty"`
	UpdatedTime                  *string `json:"updatedTime,omitempty"`
}

// GetTransferAvailableAmount represents available transfer amount
type GetTransferAvailableAmount struct {
	AvailableAmount *string `json:"availableAmount,omitempty"`
}

// CreateTransferOut represents the result of creating a transfer out
type CreateTransferOut struct {
	Id                           *string      `json:"id,omitempty"`
	UserId                       *string      `json:"userId,omitempty"`
	AccountId                    *string      `json:"accountId,omitempty"`
	CoinId                       *string      `json:"coinId,omitempty"`
	Amount                       *string      `json:"amount,omitempty"`
	ReceiverAccountId            *string      `json:"receiverAccountId,omitempty"`
	ReceiverL2Key                *string      `json:"receiverL2Key,omitempty"`
	ClientTransferId             *string      `json:"clientTransferId,omitempty"`
	IsConditionTransfer          *bool        `json:"isConditionTransfer,omitempty"`
	ConditionFactRegistryAddress *string      `json:"conditionFactRegistryAddress,omitempty"`
	ConditionFactErc20Address    *string      `json:"conditionFactErc20Address,omitempty"`
	ConditionFactAmount          *string      `json:"conditionFactAmount,omitempty"`
	ConditionFact                *string      `json:"conditionFact,omitempty"`
	TransferReason               *string      `json:"transferReason,omitempty"`
	L2Nonce                      *string      `json:"l2Nonce,omitempty"`
	L2ExpireTime                 *string      `json:"l2ExpireTime,omitempty"`
	L2Signature                  *L2Signature `json:"l2Signature,omitempty"`
	ExtraType                    *string      `json:"extraType,omitempty"`
	ExtraDataJson                *string      `json:"extraDataJson,omitempty"`
}

// L2Signature represents a Layer 2 signature
type L2Signature struct {
	R *string `json:"r,omitempty"`
	S *string `json:"s,omitempty"`
	V *string `json:"v,omitempty"`
}

// ResultListTransferOut represents list of transfer out records
type ResultListTransferOut struct {
	Code     string        `json:"code"`
	Data     []TransferOut `json:"data"`
	ErrorMsg string        `json:"msg"`
}

// ResultListTransferIn represents list of transfer in records
type ResultListTransferIn struct {
	Code     string       `json:"code"`
	Data     []TransferIn `json:"data"`
	ErrorMsg string       `json:"msg"`
}

// ResultGetTransferOutAvailableAmount represents available transfer out amount
type ResultGetTransferOutAvailableAmount struct {
	Code     string                      `json:"code"`
	Data     *GetTransferAvailableAmount `json:"data"`
	ErrorMsg string                      `json:"msg"`
}

// ResultCreateTransferOut represents the result of creating a transfer out
type ResultCreateTransferOut struct {
	Code     string             `json:"code"`
	Data     *CreateTransferOut `json:"data"`
	ErrorMsg string             `json:"msg"`
}

// Request parameter types

// GetTransferOutByIdParams represents parameters for GetTransferOutById
type GetTransferOutByIdParams struct {
	TransferId string
}

// GetTransferInByIdParams represents parameters for GetTransferInById
type GetTransferInByIdParams struct {
	TransferId string
}

// GetWithdrawAvailableAmountParams represents parameters for GetWithdrawAvailableAmount
type GetWithdrawAvailableAmountParams struct {
	CoinId string
}

// CreateTransferOutParams represents parameters for CreateTransferOut
type CreateTransferOutParams struct {
	CoinId            string
	Amount            string
	ReceiverAccountId string
	ReceiverL2Key     string
	TransferReason    string
	ExpireTime        time.Time
	ExtraType         *string
	ExtraDataJson     *string
}
