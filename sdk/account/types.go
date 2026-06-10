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
	Account                  *Account          `json:"account,omitempty"`
	PositionList             []Position        `json:"positionList"`
	CollateralList           []Collateral      `json:"collateralList"`
	Version                  string            `json:"version,omitempty"`
	PositionAssetList        []PositionAsset   `json:"positionAssetList,omitempty"`
	CollateralAssetModelList []CollateralAsset `json:"collateralAssetModelList,omitempty"`
	OraclePriceList          []IndexPrice      `json:"oraclePriceList,omitempty"`
	MarkPriceList            []IndexPrice      `json:"markPriceList,omitempty"`
}

// Position represents a position
type Position struct {
	UserID               string        `json:"userId,omitempty"`
	AccountID            string        `json:"accountId,omitempty"`
	CoinID               string        `json:"coinId,omitempty"`
	ContractID           string        `json:"contractId"`
	Size                 string        `json:"size,omitempty"`
	Price                string        `json:"price,omitempty"`
	OpenSize             string        `json:"openSize,omitempty"`
	OpenValue            string        `json:"openValue,omitempty"`
	OpenFee              string        `json:"openFee,omitempty"`
	FundingFee           string        `json:"fundingFee,omitempty"`
	LongTermCount        *int32        `json:"longTermCount,omitempty"`
	LongTermStat         *PositionStat `json:"longTermStat,omitempty"`
	LongTermCreatedTime  string        `json:"longTermCreatedTime,omitempty"`
	LongTermUpdatedTime  string        `json:"longTermUpdatedTime,omitempty"`
	ShortTermCount       *int32        `json:"shortTermCount,omitempty"`
	ShortTermStat        *PositionStat `json:"shortTermStat,omitempty"`
	ShortTermCreatedTime string        `json:"shortTermCreatedTime,omitempty"`
	ShortTermUpdatedTime string        `json:"shortTermUpdatedTime,omitempty"`
	LongTotalStat        *PositionStat `json:"longTotalStat,omitempty"`
	ShortTotalStat       *PositionStat `json:"shortTotalStat,omitempty"`
	CreatedTime          string        `json:"createdTime,omitempty"`
	UpdatedTime          string        `json:"updatedTime,omitempty"`
	MarginMode           string        `json:"marginMode,omitempty"`
	IsolatedMargin       string        `json:"isolatedMargin,omitempty"`
	AdjustedMargin       string        `json:"adjustedMargin,omitempty"`
	IsLiquidating        *bool         `json:"isLiquidating,omitempty"`
}

// Collateral represents collateral information
type Collateral struct {
	CoinID               string `json:"coinId"`
	Amount               string `json:"amount"`
	CumDepositAmount     string `json:"cumDepositAmount,omitempty"`
	CumWithdrawAmount    string `json:"cumWithdrawAmount,omitempty"`
	CumTransferInAmount  string `json:"cumTransferInAmount,omitempty"`
	CumTransferOutAmount string `json:"cumTransferOutAmount,omitempty"`
	CreatedTime          string `json:"createdTime,omitempty"`
	UpdatedTime          string `json:"updatedTime,omitempty"`
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
	ID                        string                  `json:"id"`
	UserID                    string                  `json:"userId"`
	EthAddress                string                  `json:"ethAddress"`
	ExternalAddress           string                  `json:"externalAddress,omitempty"`
	AccountName               string                  `json:"accountName,omitempty"`
	ClientAccountID           string                  `json:"clientAccountId,omitempty"`
	L2Key                     string                  `json:"l2Key,omitempty"`
	IsSystemAccount           *bool                   `json:"isSystemAccount,omitempty"`
	Signers                   []string                `json:"signers,omitempty"`
	SignerPermissions         map[string]string       `json:"signerPermissions,omitempty"`
	DefaultTradeSetting       *TradeSetting           `json:"defaultTradeSetting,omitempty"`
	ContractIDToTradeSetting  map[string]TradeSetting `json:"contractIdToTradeSetting,omitempty"`
	MaxLeverageLimit          string                  `json:"maxLeverageLimit,omitempty"`
	CreateOrderPerMinuteLimit *int32                  `json:"createOrderPerMinuteLimit,omitempty"`
	CreateOrderDelayMillis    *int32                  `json:"createOrderDelayMillis,omitempty"`
	ExtraType                 string                  `json:"extraType,omitempty"`
	ExtraDataJson             string                  `json:"extraDataJson,omitempty"`
	Status                    string                  `json:"status,omitempty"`
	IsLiquidating             *bool                   `json:"isLiquidating,omitempty"`
	ContractIDToMarginMode    map[string]string       `json:"contractIdToMarginMode,omitempty"`
	CreatedTime               string                  `json:"createdTime,omitempty"`
	UpdatedTime               string                  `json:"updatedTime,omitempty"`
}

// TradeSetting represents leverage and trading preferences.
type TradeSetting struct {
	ContractID   string `json:"contractId,omitempty"`
	Leverage     string `json:"leverage,omitempty"`
	MaxLeverage  string `json:"maxLeverage,omitempty"`
	MarginMode   string `json:"marginMode,omitempty"`
	PositionMode string `json:"positionMode,omitempty"`
	CreatedTime  string `json:"createdTime,omitempty"`
	UpdatedTime  string `json:"updatedTime,omitempty"`
}

// PositionStat represents cumulative position statistics.
type PositionStat struct {
	CumOpenSize       string `json:"cumOpenSize,omitempty"`
	CumOpenValue      string `json:"cumOpenValue,omitempty"`
	CumOpenFee        string `json:"cumOpenFee,omitempty"`
	CumCloseSize      string `json:"cumCloseSize,omitempty"`
	CumCloseValue     string `json:"cumCloseValue,omitempty"`
	CumCloseFee       string `json:"cumCloseFee,omitempty"`
	CumFundingFee     string `json:"cumFundingFee,omitempty"`
	CumLiquidateFee   string `json:"cumLiquidateFee,omitempty"`
	CumRealizePnl     string `json:"cumRealizePnl,omitempty"`
	CumDeleverageSize string `json:"cumDeleverageSize,omitempty"`
}

// PositionAsset represents position-level asset information.
type PositionAsset struct {
	CoinID                string `json:"coinId,omitempty"`
	ContractID            string `json:"contractId,omitempty"`
	PositionValue         string `json:"positionValue,omitempty"`
	OpenOrderValue        string `json:"openOrderValue,omitempty"`
	NotionalValue         string `json:"notionalValue,omitempty"`
	InitialMargin         string `json:"initialMargin,omitempty"`
	MaintenanceMargin     string `json:"maintenanceMargin,omitempty"`
	UnrealizedPnl         string `json:"unrealizedPnl,omitempty"`
	RealizedPnl           string `json:"realizedPnl,omitempty"`
	LiquidationPrice      string `json:"liquidationPrice,omitempty"`
	BankruptcyPrice       string `json:"bankruptcyPrice,omitempty"`
	MaxWithdrawableAmount string `json:"maxWithdrawableAmount,omitempty"`
}

// CollateralAsset represents account-level asset information.
type CollateralAsset struct {
	CoinID                string `json:"coinId,omitempty"`
	TotalEquity           string `json:"totalEquity,omitempty"`
	WalletBalance         string `json:"walletBalance,omitempty"`
	AvailableBalance      string `json:"availableBalance,omitempty"`
	InitialMargin         string `json:"initialMargin,omitempty"`
	MaintenanceMargin     string `json:"maintenanceMargin,omitempty"`
	OrderMargin           string `json:"orderMargin,omitempty"`
	PositionMargin        string `json:"positionMargin,omitempty"`
	UnrealizedPnl         string `json:"unrealizedPnl,omitempty"`
	CrossedUnrealizedPnl  string `json:"crossedUnrealizedPnl,omitempty"`
	IsolatedUnrealizedPnl string `json:"isolatedUnrealizedPnl,omitempty"`
	MaxWithdrawableAmount string `json:"maxWithdrawableAmount,omitempty"`
}

// IndexPrice represents oracle or mark price information.
type IndexPrice struct {
	ContractID           string `json:"contractId,omitempty"`
	PriceType            string `json:"priceType,omitempty"`
	PriceValue           string `json:"priceValue,omitempty"`
	CreatedTime          string `json:"createdTime,omitempty"`
	OraclePriceSignature string `json:"oraclePriceSignature,omitempty"`
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

// SetMarginModeParams represents the parameters for SetMarginMode.
type SetMarginModeParams struct {
	ContractID    string
	MarginMode    string
	ClientOrderID string
}

// SetMarginModeResponse represents the response for SetMarginMode.
type SetMarginModeResponse struct {
	Code       string                 `json:"code"`
	Data       map[string]interface{} `json:"data"`
	ErrorParam interface{}            `json:"errorParam"`
	ErrorMsg   string                 `json:"msg"`
}

// GetAccountPageParams represents params for v2 account page retrieval.
type GetAccountPageParams struct {
	Size       int32
	OffsetData string
}

// PageDataAccount represents paginated account list.
type PageDataAccount struct {
	DataList           []Account `json:"dataList,omitempty"`
	NextPageOffsetData *string   `json:"nextPageOffsetData,omitempty"`
}

// PageDataAccountResponse represents paginated account list response.
type PageDataAccountResponse struct {
	Code       string           `json:"code"`
	Data       *PageDataAccount `json:"data"`
	ErrorParam interface{}      `json:"errorParam"`
	ErrorMsg   string           `json:"msg"`
}

// GetPositionTransactionByIDParams represents the parameters for GetPositionTransactionByID
type GetPositionTransactionByIDParams struct {
	TransactionIDList []string
}

// GetCollateralTransactionByIDParams represents the parameters for GetCollateralTransactionByID
type GetCollateralTransactionByIDParams struct {
	TransactionIDList []string
}
