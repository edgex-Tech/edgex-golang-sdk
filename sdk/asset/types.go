package asset

// Response types for asset API

// BaseResponse is the base response structure
type BaseResponse struct {
	Code       string        `json:"code"`
	Data       interface{}   `json:"data"`
	ErrorParam []interface{} `json:"errorParam"`
}

// ResultPageDataAssetOrder represents paginated asset orders
type ResultPageDataAssetOrder struct {
	Code       string        `json:"code"`
	Data       *PageData     `json:"data"`
	ErrorParam []interface{} `json:"errorParam"`
}

// PageData represents pagination data
type PageData struct {
	List       []interface{} `json:"list"`
	OffsetData string        `json:"offsetData"`
}

// ResultGetCoinRate represents coin rate information
type ResultGetCoinRate struct {
	Code       string        `json:"code"`
	Data       interface{}   `json:"data"`
	ErrorParam []interface{} `json:"errorParam"`
}

// ResultListCrossWithdraw represents list of cross withdrawals
type ResultListCrossWithdraw struct {
	Code       string        `json:"code"`
	Data       []interface{} `json:"data"`
	ErrorParam []interface{} `json:"errorParam"`
}

// ResultGetCrossWithdrawSignInfo represents cross withdraw sign info
type ResultGetCrossWithdrawSignInfo struct {
	Code       string        `json:"code"`
	Data       interface{}   `json:"data"`
	ErrorParam []interface{} `json:"errorParam"`
}

// ResultListFastWithdraw represents list of fast withdrawals
type ResultListFastWithdraw struct {
	Code       string        `json:"code"`
	Data       []interface{} `json:"data"`
	ErrorParam []interface{} `json:"errorParam"`
}

// ResultGetFastWithdrawSignInfo represents fast withdraw sign info
type ResultGetFastWithdrawSignInfo struct {
	Code       string        `json:"code"`
	Data       interface{}   `json:"data"`
	ErrorParam []interface{} `json:"errorParam"`
}

// ResultListNormalWithdraw represents list of normal withdrawals
type ResultListNormalWithdraw struct {
	Code       string        `json:"code"`
	Data       []interface{} `json:"data"`
	ErrorParam []interface{} `json:"errorParam"`
}

// ResultGetNormalWithdrawableAmount represents normal withdrawable amount
type ResultGetNormalWithdrawableAmount struct {
	Code       string        `json:"code"`
	Data       interface{}   `json:"data"`
	ErrorParam []interface{} `json:"errorParam"`
}

// ResultCreateNormalWithdraw represents result of creating normal withdrawal
type ResultCreateNormalWithdraw struct {
	Code       string        `json:"code"`
	Data       interface{}   `json:"data"`
	ErrorParam []interface{} `json:"errorParam"`
}

// ResultCreateCrossWithdraw represents result of creating cross withdrawal
type ResultCreateCrossWithdraw struct {
	Code       string        `json:"code"`
	Data       interface{}   `json:"data"`
	ErrorParam []interface{} `json:"errorParam"`
}

// ResultCreateFastWithdraw represents result of creating fast withdrawal
type ResultCreateFastWithdraw struct {
	Code       string        `json:"code"`
	Data       interface{}   `json:"data"`
	ErrorParam []interface{} `json:"errorParam"`
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
	CoinId           string
	Amount           string
	EthAddress       string
	ClientWithdrawId string
	ExpireTime       string
	L2Signature      string
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
