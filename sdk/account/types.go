package account

// Response types for account API

// BaseResponse is the base response structure
type BaseResponse struct {
	Code       string        `json:"code"`
	Data       interface{}   `json:"data"`
	ErrorParam []interface{} `json:"errorParam"`
}

// GetAccountAssetResponse represents the response for GetAccountAsset
type GetAccountAssetResponse struct {
	Code       string            `json:"code"`
	Data       *AccountAssetData `json:"data"`
	ErrorParam []interface{}     `json:"errorParam"`
}

// AccountAssetData contains account asset information
type AccountAssetData struct {
	PositionList   []Position   `json:"positionList"`
	CollateralList []Collateral `json:"collateralList"`
}

// Position represents a position
type Position struct {
	ContractID string `json:"contractId"`
	Size       string `json:"size"`
	Price      string `json:"price"`
}

// Collateral represents collateral information
type Collateral struct {
	CoinID string `json:"coinId"`
	Amount string `json:"amount"`
}

// ListPositionResponse represents the response for GetAccountPositions
type ListPositionResponse struct {
	Code       string        `json:"code"`
	Data       []Position    `json:"data"`
	ErrorParam []interface{} `json:"errorParam"`
}

// PageDataPositionTransactionResponse represents paginated position transactions
type PageDataPositionTransactionResponse struct {
	Code       string        `json:"code"`
	Data       *PageData     `json:"data"`
	ErrorParam []interface{} `json:"errorParam"`
}

// PageData represents pagination data
type PageData struct {
	List       []interface{} `json:"list"`
	OffsetData string        `json:"offsetData"`
}

// PageDataCollateralTransactionResponse represents paginated collateral transactions
type PageDataCollateralTransactionResponse struct {
	Code       string        `json:"code"`
	Data       *PageData     `json:"data"`
	ErrorParam []interface{} `json:"errorParam"`
}

// PageDataPositionTermResponse represents paginated position terms
type PageDataPositionTermResponse struct {
	Code       string        `json:"code"`
	Data       *PageData     `json:"data"`
	ErrorParam []interface{} `json:"errorParam"`
}

// ListCollateralResponse represents the response for GetCollateralByCoinID
type ListCollateralResponse struct {
	Code       string        `json:"code"`
	Data       []Collateral  `json:"data"`
	ErrorParam []interface{} `json:"errorParam"`
}

// AccountResponse represents the response for GetAccountByID
type AccountResponse struct {
	Code       string        `json:"code"`
	Data       *Account      `json:"data"`
	ErrorParam []interface{} `json:"errorParam"`
}

// Account represents account information
type Account struct {
	ID         string `json:"id"`
	UserID     string `json:"userId"`
	EthAddress string `json:"ethAddress"`
	L2Key      string `json:"l2Key"`
}

// PageDataAccountAssetSnapshotResponse represents paginated account asset snapshots
type PageDataAccountAssetSnapshotResponse struct {
	Code       string        `json:"code"`
	Data       *PageData     `json:"data"`
	ErrorParam []interface{} `json:"errorParam"`
}

// ListPositionTransactionResponse represents the response for GetPositionTransactionByID
type ListPositionTransactionResponse struct {
	Code       string        `json:"code"`
	Data       []interface{} `json:"data"`
	ErrorParam []interface{} `json:"errorParam"`
}

// ListCollateralTransactionResponse represents the response for GetCollateralTransactionByID
type ListCollateralTransactionResponse struct {
	Code       string        `json:"code"`
	Data       []interface{} `json:"data"`
	ErrorParam []interface{} `json:"errorParam"`
}

// GetAccountDeleverageLightResponse represents the response for GetAccountDeleverageLight
type GetAccountDeleverageLightResponse struct {
	Code       string        `json:"code"`
	Data       interface{}   `json:"data"`
	ErrorParam []interface{} `json:"errorParam"`
}

// UpdateLeverageSettingResponse represents the response for UpdateLeverageSetting
type UpdateLeverageSettingResponse struct {
	Code       string        `json:"code"`
	Data       interface{}   `json:"data"`
	ErrorParam []interface{} `json:"errorParam"`
}

// Request parameter types

// GetPositionTransactionByIDParams represents the parameters for GetPositionTransactionByID
type GetPositionTransactionByIDParams struct {
	TransactionIDList []string
}

// GetCollateralTransactionByIDParams represents the parameters for GetCollateralTransactionByID
type GetCollateralTransactionByIDParams struct {
	TransactionIDList []string
}
