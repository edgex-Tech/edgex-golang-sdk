package account

// PositionTransaction represents a position transaction
type PositionTransaction struct {
	Id          *string `json:"id,omitempty"`
	UserId      *string `json:"userId,omitempty"`
	AccountId   *string `json:"accountId,omitempty"`
	ContractId  *string `json:"contractId,omitempty"`
	CoinId      *string `json:"coinId,omitempty"`
	Type        *string `json:"type,omitempty"`
	Size        *string `json:"size,omitempty"`
	Price       *string `json:"price,omitempty"`
	Fee         *string `json:"fee,omitempty"`
	CreatedTime *string `json:"createdTime,omitempty"`
}

// PageDataPositionTransaction represents paginated position transaction data
type PageDataPositionTransaction struct {
	DataList           []PositionTransaction `json:"dataList,omitempty"`
	NextPageOffsetData *string               `json:"nextPageOffsetData,omitempty"`
}

// CollateralTransaction represents a collateral transaction
type CollateralTransaction struct {
	Id          *string `json:"id,omitempty"`
	UserId      *string `json:"userId,omitempty"`
	AccountId   *string `json:"accountId,omitempty"`
	CoinId      *string `json:"coinId,omitempty"`
	Type        *string `json:"type,omitempty"`
	Amount      *string `json:"amount,omitempty"`
	CreatedTime *string `json:"createdTime,omitempty"`
}

// PageDataCollateralTransaction represents paginated collateral transaction data
type PageDataCollateralTransaction struct {
	DataList           []CollateralTransaction `json:"dataList,omitempty"`
	NextPageOffsetData *string                 `json:"nextPageOffsetData,omitempty"`
}

// PositionTerm represents a position term
type PositionTerm struct {
	Id             *string `json:"id,omitempty"`
	UserId         *string `json:"userId,omitempty"`
	AccountId      *string `json:"accountId,omitempty"`
	ContractId     *string `json:"contractId,omitempty"`
	CoinId         *string `json:"coinId,omitempty"`
	IsLongPosition *bool   `json:"isLongPosition,omitempty"`
	Size           *string `json:"size,omitempty"`
	Price          *string `json:"price,omitempty"`
	CreatedTime    *string `json:"createdTime,omitempty"`
}

// PageDataPositionTerm represents paginated position term data
type PageDataPositionTerm struct {
	DataList           []PositionTerm `json:"dataList,omitempty"`
	NextPageOffsetData *string        `json:"nextPageOffsetData,omitempty"`
}

// AccountAssetSnapshot represents an account asset snapshot
type AccountAssetSnapshot struct {
	Id          *string `json:"id,omitempty"`
	UserId      *string `json:"userId,omitempty"`
	AccountId   *string `json:"accountId,omitempty"`
	CoinId      *string `json:"coinId,omitempty"`
	Amount      *string `json:"amount,omitempty"`
	TimeTag     *int32  `json:"timeTag,omitempty"`
	CreatedTime *string `json:"createdTime,omitempty"`
}

// PageDataAccountAssetSnapshot represents paginated account asset snapshot data
type PageDataAccountAssetSnapshot struct {
	DataList           []AccountAssetSnapshot `json:"dataList,omitempty"`
	NextPageOffsetData *string                `json:"nextPageOffsetData,omitempty"`
}

// GetAccountDeleverageLight represents account deleverage light information
type GetAccountDeleverageLight struct {
	DeleverageLevel *string `json:"deleverageLevel,omitempty"`
}

// GetAccountAssetResponse represents the response for GetAccountAsset
type GetAccountAssetResponse struct {
	Code       string            `json:"code"`
	Data       *AccountAssetData `json:"data"`
	ErrorParam interface{}       `json:"errorParam"`
	ErrorMsg   string            `json:"msg"`
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
	Code       string      `json:"code"`
	Data       []Position  `json:"data"`
	ErrorParam interface{} `json:"errorParam"`
	ErrorMsg   string      `json:"msg"`
}

// GetPositionTransactionPageParams represents the parameters for GetPositionTransactionPage
type GetPositionTransactionPageParams struct {
	Size                   int32
	OffsetData             string
	FilterCoinIDList       []string
	FilterContractIDList   []string
	FilterTypeList         []string
	FilterStartCreatedTime int64
	FilterEndCreatedTime   int64
	FilterCloseOnly        *bool
	FilterOpenOnly         *bool
}

// PageDataPositionTransactionResponse represents paginated position transactions
type PageDataPositionTransactionResponse struct {
	Code       string                       `json:"code"`
	Data       *PageDataPositionTransaction `json:"data"`
	ErrorParam interface{}                  `json:"errorParam"`
	ErrorMsg   string                       `json:"msg"`
}

// GetCollateralTransactionPageParams represents the parameters for GetCollateralTransactionPage
type GetCollateralTransactionPageParams struct {
	Size                   int32
	OffsetData             string
	FilterCoinIDList       []string
	FilterTypeList         []string
	FilterStartCreatedTime int64
	FilterEndCreatedTime   int64
}

// PageDataCollateralTransactionResponse represents paginated collateral transactions
type PageDataCollateralTransactionResponse struct {
	Code       string                         `json:"code"`
	Data       *PageDataCollateralTransaction `json:"data"`
	ErrorParam interface{}                    `json:"errorParam"`
	ErrorMsg   string                         `json:"msg"`
}

// GetPositionTermPageParams represents the parameters for GetPositionTermPage
type GetPositionTermPageParams struct {
	Size                   int32
	OffsetData             string
	FilterCoinIDList       []string
	FilterContractIDList   []string
	FilterIsLongPosition   *bool
	FilterStartCreatedTime int64
	FilterEndCreatedTime   int64
}

// PageDataPositionTermResponse represents paginated position terms
type PageDataPositionTermResponse struct {
	Code       string                `json:"code"`
	Data       *PageDataPositionTerm `json:"data"`
	ErrorParam interface{}           `json:"errorParam"`
	ErrorMsg   string                `json:"msg"`
}

// ListCollateralResponse represents the response for GetCollateralByCoinID
type ListCollateralResponse struct {
	Code       string       `json:"code"`
	Data       []Collateral `json:"data"`
	ErrorParam interface{}  `json:"errorParam"`
	ErrorMsg   string       `json:"msg"`
}

// AccountResponse represents the response for GetAccountByID
type AccountResponse struct {
	Code       string      `json:"code"`
	Data       *Account    `json:"data"`
	ErrorParam interface{} `json:"errorParam"`
	ErrorMsg   string      `json:"msg"`
}

// Account represents account information
type Account struct {
	ID         string `json:"id"`
	UserID     string `json:"userId"`
	EthAddress string `json:"ethAddress"`
	L2Key      string `json:"l2Key"`
}

// GetAccountAssetSnapshotPageParams represents the parameters for GetAccountAssetSnapshotPage
type GetAccountAssetSnapshotPageParams struct {
	Size            int32
	OffsetData      string
	CoinID          string
	FilterTimeTag   *int32
	FilterStartTime int64
	FilterEndTime   int64
}

// PageDataAccountAssetSnapshotResponse represents paginated account asset snapshots
type PageDataAccountAssetSnapshotResponse struct {
	Code       string                        `json:"code"`
	Data       *PageDataAccountAssetSnapshot `json:"data"`
	ErrorParam interface{}                   `json:"errorParam"`
	ErrorMsg   string                        `json:"msg"`
}

// ListPositionTransactionResponse represents the response for GetPositionTransactionByID
type ListPositionTransactionResponse struct {
	Code       string        `json:"code"`
	Data       []interface{} `json:"data"`
	ErrorParam interface{}   `json:"errorParam"`
	ErrorMsg   string        `json:"msg"`
}

// ListCollateralTransactionResponse represents the response for GetCollateralTransactionByID
type ListCollateralTransactionResponse struct {
	Code       string        `json:"code"`
	Data       []interface{} `json:"data"`
	ErrorParam interface{}   `json:"errorParam"`
	ErrorMsg   string        `json:"msg"`
}

// GetAccountDeleverageLightResponse represents the response for GetAccountDeleverageLight
type GetAccountDeleverageLightResponse struct {
	Code       string                     `json:"code"`
	Data       *GetAccountDeleverageLight `json:"data"`
	ErrorParam interface{}                `json:"errorParam"`
	ErrorMsg   string                     `json:"msg"`
}

// UpdateLeverageSettingResponse represents the response for UpdateLeverageSetting
type UpdateLeverageSettingResponse struct {
	Code       string                 `json:"code"`
	Data       map[string]interface{} `json:"data"`
	ErrorParam interface{}            `json:"errorParam"`
	ErrorMsg   string                 `json:"msg"`
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
