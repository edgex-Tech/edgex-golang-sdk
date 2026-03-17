package quote

// KlineType represents the K-line interval type
type KlineType string

const (
	KlineTypeUnknown  KlineType = "UNKNOWN_KLINE_TYPE"
	KlineType1Minute  KlineType = "MINUTE_1"
	KlineType5Minute  KlineType = "MINUTE_5"
	KlineType15Minute KlineType = "MINUTE_15"
	KlineType30Minute KlineType = "MINUTE_30"
	KlineType1Hour    KlineType = "HOUR_1"
	KlineType2Hour    KlineType = "HOUR_2"
	KlineType4Hour    KlineType = "HOUR_4"
	KlineType6Hour    KlineType = "HOUR_6"
	KlineType8Hour    KlineType = "HOUR_8"
	KlineType12Hour   KlineType = "HOUR_12"
	KlineType1Day     KlineType = "DAY_1"
	KlineType1Week    KlineType = "WEEK_1"
	KlineType1Month   KlineType = "MONTH_1"
)

// PriceType represents the price type for K-line data
type PriceType string

const (
	PriceTypeUnknown      PriceType = "UNKNOWN_PRICE_TYPE"
	PriceTypeOraclePrice  PriceType = "ORACLE_PRICE"
	PriceTypeIndexPrice   PriceType = "INDEX_PRICE"
	PriceTypeLastPrice    PriceType = "LAST_PRICE"
	PriceTypeAsk1Price    PriceType = "ASK1_PRICE"
	PriceTypeBid1Price    PriceType = "BID1_PRICE"
	PriceTypeOpenInterest PriceType = "OPEN_INTEREST"
)

// TickerSummary represents ticker summary data
type TickerSummary struct {
	Period       *string `json:"period,omitempty"`
	Trades       *string `json:"trades,omitempty"`
	Value        *string `json:"value,omitempty"`
	OpenInterest *string `json:"openInterest,omitempty"`
}

// Ticker represents 24-hour ticker data
type Ticker struct {
	ContractId         *string `json:"contractId,omitempty"`
	ContractName       *string `json:"contractName,omitempty"`
	PriceChange        *string `json:"priceChange,omitempty"`
	PriceChangePercent *string `json:"priceChangePercent,omitempty"`
	Trades             *string `json:"trades,omitempty"`
	Size               *string `json:"size,omitempty"`
	Value              *string `json:"value,omitempty"`
	High               *string `json:"high,omitempty"`
	Low                *string `json:"low,omitempty"`
	Open               *string `json:"open,omitempty"`
	Close              *string `json:"close,omitempty"`
	HighTime           *string `json:"highTime,omitempty"`
	LowTime            *string `json:"lowTime,omitempty"`
	StartTime          *string `json:"startTime,omitempty"`
	EndTime            *string `json:"endTime,omitempty"`
	LastPrice          *string `json:"lastPrice,omitempty"`
	IndexPrice         *string `json:"indexPrice,omitempty"`
	OraclePrice        *string `json:"oraclePrice,omitempty"`
	OpenInterest       *string `json:"openInterest,omitempty"`
	FundingRate        *string `json:"fundingRate,omitempty"`
	FundingTime        *string `json:"fundingTime,omitempty"`
	NextFundingTime    *string `json:"nextFundingTime,omitempty"`
}

// Kline represents K-line data
type Kline struct {
	KlineId       *string `json:"klineId,omitempty"`
	ContractId    *string `json:"contractId,omitempty"`
	ContractName  *string `json:"contractName,omitempty"`
	KlineType     *string `json:"klineType,omitempty"`
	KlineTime     *string `json:"klineTime,omitempty"`
	PriceType     *string `json:"priceType,omitempty"`
	Trades        *string `json:"trades,omitempty"`
	Size          *string `json:"size,omitempty"`
	Value         *string `json:"value,omitempty"`
	High          *string `json:"high,omitempty"`
	Low           *string `json:"low,omitempty"`
	Open          *string `json:"open,omitempty"`
	Close         *string `json:"close,omitempty"`
	MakerBuySize  *string `json:"makerBuySize,omitempty"`
	MakerBuyValue *string `json:"makerBuyValue,omitempty"`
}

// BookOrder represents an order book entry
type BookOrder struct {
	Price *string `json:"price,omitempty"`
	Size  *string `json:"size,omitempty"`
}

// Depth represents order book depth data
type Depth struct {
	StartVersion *string     `json:"startVersion,omitempty"`
	EndVersion   *string     `json:"endVersion,omitempty"`
	Level        *int32      `json:"level,omitempty"`
	ContractId   *string     `json:"contractId,omitempty"`
	ContractName *string     `json:"contractName,omitempty"`
	Asks         []BookOrder `json:"asks,omitempty"`
	Bids         []BookOrder `json:"bids,omitempty"`
	DepthType    *string     `json:"depthType,omitempty"`
}

// PageDataKline represents paginated K-line data
type PageDataKline struct {
	DataList           []Kline `json:"dataList,omitempty"`
	NextPageOffsetData *string `json:"nextPageOffsetData,omitempty"`
}

// ContractMultiKline represents contract multi K-line data
type ContractMultiKline struct {
	ContractId *string `json:"contractId,omitempty"`
	DataList   []Kline `json:"dataList,omitempty"`
}

// Response types for quote API

// ResultGetTickerSummaryModel represents ticker summary
type ResultGetTickerSummaryModel struct {
	Code       string         `json:"code"`
	Data       *TickerSummary `json:"data"`
	ErrorParam interface{}    `json:"errorParam"`
	ErrorMsg   string         `json:"msg"`
}

// ResultListTicker represents list of tickers
type ResultListTicker struct {
	Code       string      `json:"code"`
	Data       []Ticker    `json:"data"`
	ErrorParam interface{} `json:"errorParam"`
	ErrorMsg   string      `json:"msg"`
}

// ResultPageDataKline represents paginated K-line data
type ResultPageDataKline struct {
	Code       string         `json:"code"`
	Data       *PageDataKline `json:"data"`
	ErrorParam interface{}    `json:"errorParam"`
	ErrorMsg   string       `json:"msg"`
}

// ResultListDepth represents list of depth data
type ResultListDepth struct {
	Code       string      `json:"code"`
	Data       []Depth     `json:"data"`
	ErrorParam interface{} `json:"errorParam"`
	ErrorMsg   string       `json:"msg"`
}

// ResultListContractKline represents list of contract K-line data
type ResultListContractKline struct {
	Code       string               `json:"code"`
	Data       []ContractMultiKline `json:"data"`
	ErrorParam interface{}          `json:"errorParam"`
	ErrorMsg   string       `json:"msg"`
}

// Request parameter types

// GetKLineParams represents parameters for GetKLine
type GetKLineParams struct {
	ContractID string
	Interval   KlineType
	PriceType  PriceType
	Size       int64
	OffsetData string
	From       *int64
	To         *int64
}

// GetOrderBookDepthParams represents parameters for GetOrderBookDepth
type GetOrderBookDepthParams struct {
	ContractID string
	Size       int64
	Precision  *string
}

// GetMultiContractKLineParams represents parameters for GetMultiContractKLine
type GetMultiContractKLineParams struct {
	ContractIDs []string
	Interval    KlineType
	Size        int64
	PriceType   PriceType
	From        *int64
	To          *int64
}
