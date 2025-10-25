package asset

// PageDataAssetOrder represents paginated asset order data
type PageDataAssetOrder struct {
	DataList           []interface{} `json:"dataList,omitempty"`
	NextPageOffsetData *string       `json:"nextPageOffsetData,omitempty"`
}

// GetCoinRate represents coin rate information
type GetCoinRate struct {
	Rate *string `json:"rate,omitempty"`
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

// GetCrossWithdrawSignInfo represents cross withdraw sign info
type GetCrossWithdrawSignInfo struct {
	LpAccountId            *string `json:"lpAccountId,omitempty"`
	CrossWithdrawL2Key     *string `json:"crossWithdrawL2Key,omitempty"`
	CrossWithdrawMaxAmount *string `json:"crossWithdrawMaxAmount,omitempty"`
	Fee                    *string `json:"fee,omitempty"`
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

// GetFastWithdrawSignInfo represents fast withdraw sign info
type GetFastWithdrawSignInfo struct {
	LpAccountId                     *string `json:"lpAccountId,omitempty"`
	FastWithdrawL2Key               *string `json:"fastWithdrawL2Key,omitempty"`
	FastWithdrawFactRegisterAddress *string `json:"fastWithdrawFactRegisterAddress,omitempty"`
	FastWithdrawMaxAmount           *string `json:"fastWithdrawMaxAmount,omitempty"`
	Fee                             *string `json:"fee,omitempty"`
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

// GetNormalWithdrawableAmount represents normal withdrawable amount
type GetNormalWithdrawableAmount struct {
	Amount *string `json:"amount,omitempty"`
}

// ResultPageDataAssetOrder represents paginated asset orders
type ResultPageDataAssetOrder struct {
	Code       string              `json:"code"`
	Data       *PageDataAssetOrder `json:"data"`
	ErrorParam interface{}         `json:"errorParam"`
	ErrorMsg   string              `json:"msg"`
}

// ResultGetCoinRate represents coin rate information
type ResultGetCoinRate struct {
	Code       string       `json:"code"`
	Data       *GetCoinRate `json:"data"`
	ErrorParam interface{}  `json:"errorParam"`
	ErrorMsg   string       `json:"msg"`
}

// ResultListCrossWithdraw represents list of cross withdrawals
type ResultListCrossWithdraw struct {
	Code       string                `json:"code"`
	Data       []CreateCrossWithdraw `json:"data"`
	ErrorParam interface{}           `json:"errorParam"`
	ErrorMsg   string                `json:"msg"`
}

// ResultGetCrossWithdrawSignInfo represents cross withdraw sign info
type ResultGetCrossWithdrawSignInfo struct {
	Code       string                    `json:"code"`
	Data       *GetCrossWithdrawSignInfo `json:"data"`
	ErrorParam interface{}               `json:"errorParam"`
	ErrorMsg   string                    `json:"msg"`
}

// ResultListFastWithdraw represents list of fast withdrawals
type ResultListFastWithdraw struct {
	Code       string               `json:"code"`
	Data       []CreateFastWithdraw `json:"data"`
	ErrorParam interface{}          `json:"errorParam"`
	ErrorMsg   string               `json:"msg"`
}

// ResultGetFastWithdrawSignInfo represents fast withdraw sign info
type ResultGetFastWithdrawSignInfo struct {
	Code       string                   `json:"code"`
	Data       *GetFastWithdrawSignInfo `json:"data"`
	ErrorParam interface{}              `json:"errorParam"`
	ErrorMsg   string                   `json:"msg"`
}

// ResultListNormalWithdraw represents list of normal withdrawals
type ResultListNormalWithdraw struct {
	Code       string                 `json:"code"`
	Data       []CreateNormalWithdraw `json:"data"`
	ErrorParam interface{}            `json:"errorParam"`
	ErrorMsg   string                 `json:"msg"`
}

// ResultGetNormalWithdrawableAmount represents normal withdrawable amount
type ResultGetNormalWithdrawableAmount struct {
	Code       string                       `json:"code"`
	Data       *GetNormalWithdrawableAmount `json:"data"`
	ErrorParam interface{}                  `json:"errorParam"`
	ErrorMsg   string                       `json:"msg"`
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

// Request parameter types

// GetAllOrdersPageParams represents parameters for GetAllOrdersPage
type GetAllOrdersPageParams struct {
	StartTime  string
	EndTime    string
	ChainId    string
	TypeList   string
	Size       string
	OffsetData string
}

// GetCoinRateParams represents parameters for GetCoinRate
type GetCoinRateParams struct {
	ChainId string
	Coin    string
}

// GetCrossWithdrawByIdParams represents parameters for GetCrossWithdrawById
type GetCrossWithdrawByIdParams struct {
	CrossWithdrawIdList string
}

// GetCrossWithdrawSignInfoParams represents parameters for GetCrossWithdrawSignInfo
type GetCrossWithdrawSignInfoParams struct {
	ChainId string
	Amount  string
}

// GetFastWithdrawByIdParams represents parameters for GetFastWithdrawById
type GetFastWithdrawByIdParams struct {
	FastWithdrawIdList string
}

// GetFastWithdrawSignInfoParams represents parameters for GetFastWithdrawSignInfo
type GetFastWithdrawSignInfoParams struct {
	ChainId string
	Amount  string
}

// GetNormalWithdrawByIdParams represents parameters for GetNormalWithdrawById
type GetNormalWithdrawByIdParams struct {
	NormalWithdrawIdList string
}

// GetNormalWithdrawSignInfoParams represents parameters for GetNormalWithdrawSignInfo
type GetNormalWithdrawSignInfoParams struct {
	ChainId string
	Amount  string
}

// GetNormalWithdrawableAmountParams represents parameters for GetNormalWithdrawableAmount
type GetNormalWithdrawableAmountParams struct {
	Address string
}

// CreateNormalWithdrawParams represents parameters for CreateNormalWithdraw
type CreateNormalWithdrawParams struct {
	CoinId     string
	Amount     string
	EthAddress string
}

// CreateCrossWithdrawParams represents parameters for CreateCrossWithdraw
type CreateCrossWithdrawParams struct {
	CoinId                string
	Amount                string
	EthAddress            string
	Erc20Address          string
	LpAccountId           string
	ClientCrossWithdrawId string
	ExpireTime            string
	L2Signature           string
	Fee                   string
	ChainId               string
	MpcAddress            string
	MpcSignature          string
	MpcSignTime           string
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
