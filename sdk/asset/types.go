package asset

// AssetOrderType represents the type of asset order
type AssetOrderType string

const (
	AssetOrderTypeUnknown             AssetOrderType = "UNKNOWN_ORDER_TYPE"
	AssetOrderTypeNormalDeposit       AssetOrderType = "ORDER_TYPE_NORMAL_DEPOSIT"
	AssetOrderTypeNormalWithdraw      AssetOrderType = "ORDER_TYPE_NORMAL_WITHDRAW"
	AssetOrderTypeInternalTransferIn  AssetOrderType = "ORDER_TYPE_INTERNAL_TRANSFER_IN"
	AssetOrderTypeInternalTransferOut AssetOrderType = "ORDER_TYPE_INTERNAL_TRANSFER_OUT"
	AssetOrderTypeCrossDeposit        AssetOrderType = "ORDER_TYPE_CROSS_DEPOSIT"
	AssetOrderTypeCrossWithdraw       AssetOrderType = "ORDER_TYPE_CROSS_WITHDRAW"
	AssetOrderTypeFastWithdraw        AssetOrderType = "ORDER_TYPE_FAST_WITHDRAW"
	AssetOrderTypeAirDrop             AssetOrderType = "ORDER_TYPE_AIR_DROP"
	AssetOrderTypeTransferIn          AssetOrderType = "ORDER_TYPE_TRANSFER_IN"
	AssetOrderTypeTransferOut         AssetOrderType = "ORDER_TYPE_TRANSFER_OUT"
)

// PageDataAssetOrder represents paginated asset order data
type PageDataAssetOrder struct {
	DataList           []interface{} `json:"dataList,omitempty"`
	NextPageOffsetData *string       `json:"nextPageOffsetData,omitempty"`
}

// CreateCrossWithdraw represents cross withdrawal information
type CreateCrossWithdraw struct {
	Id               *string `json:"id,omitempty"`
	UserId           *string `json:"userId,omitempty"`
	AccountId        *string `json:"accountId,omitempty"`
	CoinId           *string `json:"coinId,omitempty"`
	Amount           *string `json:"amount,omitempty"`
	ReceiverAddress  *string `json:"receiverAddress,omitempty"`
	ReceiverChainId  *string `json:"receiverChainId,omitempty"`
	ClientWithdrawId *string `json:"clientWithdrawId,omitempty"`
	Status           *string `json:"status,omitempty"`
	CreatedTime      *string `json:"createdTime,omitempty"`
	UpdatedTime      *string `json:"updatedTime,omitempty"`
}

// CreateFastWithdraw represents fast withdrawal information
type CreateFastWithdraw struct {
	Id               *string `json:"id,omitempty"`
	UserId           *string `json:"userId,omitempty"`
	AccountId        *string `json:"accountId,omitempty"`
	CoinId           *string `json:"coinId,omitempty"`
	Amount           *string `json:"amount,omitempty"`
	ReceiverAddress  *string `json:"receiverAddress,omitempty"`
	ClientWithdrawId *string `json:"clientWithdrawId,omitempty"`
	Status           *string `json:"status,omitempty"`
	CreatedTime      *string `json:"createdTime,omitempty"`
	UpdatedTime      *string `json:"updatedTime,omitempty"`
}

// CreateNormalWithdraw represents normal withdrawal information
type CreateNormalWithdraw struct {
	Id               *string `json:"id,omitempty"`
	UserId           *string `json:"userId,omitempty"`
	AccountId        *string `json:"accountId,omitempty"`
	CoinId           *string `json:"coinId,omitempty"`
	Amount           *string `json:"amount,omitempty"`
	ReceiverAddress  *string `json:"receiverAddress,omitempty"`
	ClientWithdrawId *string `json:"clientWithdrawId,omitempty"`
	Status           *string `json:"status,omitempty"`
	CreatedTime      *string `json:"createdTime,omitempty"`
	UpdatedTime      *string `json:"updatedTime,omitempty"`
}

// ResultPageDataAssetOrder represents paginated asset orders
type ResultPageDataAssetOrder struct {
	Code       string              `json:"code"`
	Data       *PageDataAssetOrder `json:"data"`
	ErrorParam interface{}         `json:"errorParam"`
	ErrorMsg   string              `json:"msg"`
}

// ResultListNormalWithdraw represents list of normal withdrawals
type ResultListNormalWithdraw struct {
	Code       string                 `json:"code"`
	Data       []CreateNormalWithdraw `json:"data"`
	ErrorParam interface{}            `json:"errorParam"`
	ErrorMsg   string                 `json:"msg"`
}

// ResultCreateNormalWithdraw represents result of creating normal withdrawal
type ResultCreateNormalWithdraw struct {
	Code       string                `json:"code"`
	Data       *CreateNormalWithdraw `json:"data"`
	ErrorParam interface{}           `json:"errorParam"`
	ErrorMsg   string                `json:"msg"`
}

// ResultCreateCrossWithdraw represents result of creating cross withdrawal
type ResultCreateCrossWithdraw struct {
	Code       string               `json:"code"`
	Data       *CreateCrossWithdraw `json:"data"`
	ErrorParam interface{}          `json:"errorParam"`
	ErrorMsg   string               `json:"msg"`
}

// ResultCreateFastWithdraw represents result of creating fast withdrawal
type ResultCreateFastWithdraw struct {
	Code       string              `json:"code"`
	Data       *CreateFastWithdraw `json:"data"`
	ErrorParam interface{}         `json:"errorParam"`
	ErrorMsg   string              `json:"msg"`
}

// GetAllOrdersPageParams represents parameters for GetAllOrdersPage
type GetAllOrdersPageParams struct {
	StartTime  string
	EndTime    string
	ChainId    string
	TypeList   string
	Size       string
	OffsetData string
}

// GetWithdrawSignInfo represents withdraw sign info
type GetWithdrawSignInfo struct {
	PoolBalance *string `json:"poolBalance,omitempty"`
	Fee         *string `json:"fee,omitempty"`
}

// ResultGetWithdrawSignInfo represents withdraw sign info result
type ResultGetWithdrawSignInfo struct {
	Code       string               `json:"code"`
	Data       *GetWithdrawSignInfo `json:"data"`
	ErrorParam interface{}          `json:"errorParam"`
	ErrorMsg   string               `json:"msg"`
}

// GetWithdrawSignInfoParams represents parameters for GetWithdrawSignInfo
type GetWithdrawSignInfoParams struct {
	ChainId      string
	TokenAddress string
	Amount       string
}

// PrepareWithdrawSignInfoParams represents parameters for preparing withdraw sign info.
type PrepareWithdrawSignInfoParams struct {
	CoinId       string
	ChainId      string
	TokenAddress string
	Amount       string
}

// PreparedWithdrawSignInfo represents prepared withdraw sign context used by withdrawal APIs.
type PreparedWithdrawSignInfo struct {
	CoinId       string
	ChainId      string
	TokenAddress string
	Amount       string
	Fee          string
	SignInfo     *ResultGetWithdrawSignInfo
}

// WithdrawV2SignatureInfo represents generated V2 EIP-712 signature payload.
type WithdrawV2SignatureInfo struct {
	Signature    string
	Signer       string
	Nonce        string
	L2ExpireTime string
	ToAddress    string
}

// BuildWithdrawV2SignatureParams represents parameters for building V2 withdraw signature.
type BuildWithdrawV2SignatureParams struct {
	CoinId           string
	Amount           string
	Fee              string
	ClientWithdrawId string
	ToAddress        string
}

// CreateNormalWithdrawParams represents parameters for CreateNormalWithdraw
type CreateNormalWithdrawParams struct {
	CoinId     string
	Amount     string
	Fee        string // Withdrawal fee (optional)
	EthAddress string
}

// CreateCrossWithdrawAutoParams represents minimal parameters for creating cross withdraw with auto fee/signature.
type CreateCrossWithdrawAutoParams struct {
	CoinId           string
	ChainId          string
	TokenAddress     string
	Amount           string
	TargetAddress    string
	Fee              string
	ClientWithdrawId string
}

// CreateCrossWithdrawParams represents parameters for CreateCrossWithdraw (V2 EIP-712)
type CreateCrossWithdrawParams struct {
	CoinId           string // Coin ID
	Amount           string // Withdrawal amount
	EthAddress       string // Signer address (Privy address)
	TargetAddress    string // Final cross-chain withdrawal target address
	ClientWithdrawId string // Client-defined ID for idempotency
	Fee              string // Withdrawal fee
	Signature        string // EIP-712 signature (with 0x prefix)
	Signer           string // Signer address
	Nonce            string // Nonce (will be converted to integer)
	L2ExpireTime     string // L2 signature expiration time in milliseconds
	ChainId          string // Target chain ID
}

// CreateFastWithdrawParams represents parameters for CreateFastWithdraw
type CreateFastWithdrawParams struct {
	CoinId           string
	Amount           string
	EthAddress       string
	ClientWithdrawId string
	ExpireTime       string
	L2Signature      string
}
