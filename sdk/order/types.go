package order

import "time"

// TimeInForce constants
type TimeInForce string

const (
	TimeInForce_UNKNOWN_TIME_IN_FORCE TimeInForce = "UNKNOWN_TIME_IN_FORCE"
	TimeInForce_GOOD_TIL_CANCEL       TimeInForce = "GOOD_TIL_CANCEL"
	TimeInForce_FILL_OR_KILL          TimeInForce = "FILL_OR_KILL"
	TimeInForce_IMMEDIATE_OR_CANCEL   TimeInForce = "IMMEDIATE_OR_CANCEL"
	TimeInForce_POST_ONLY             TimeInForce = "POST_ONLY"
)

// Order side constants
const (
	OrderSideBuy  = "BUY"
	OrderSideSell = "SELL"
)

// Response code constants
const (
	ResponseCodeSuccess = "SUCCESS"
)

// OrderType represents the type of order
type OrderType string

const (
	OrderTypeUnknown          OrderType = "UNKNOWN_ORDER_TYPE"
	OrderTypeLimit            OrderType = "LIMIT"
	OrderTypeMarket           OrderType = "MARKET"
	OrderTypeStopLimit        OrderType = "STOP_LIMIT"
	OrderTypeStopMarket       OrderType = "STOP_MARKET"
	OrderTypeTakeProfitLimit  OrderType = "TAKE_PROFIT_LIMIT"
	OrderTypeTakeProfitMarket OrderType = "TAKE_PROFIT_MARKET"
)

// L2Signature represents a Layer 2 signature
type L2Signature struct {
	R *string `json:"r,omitempty"`
	S *string `json:"s,omitempty"`
	V *string `json:"v,omitempty"`
}

// OpenTpSl represents open take-profit/stop-loss configuration
type OpenTpSl struct {
	TriggerPrice     *string `json:"triggerPrice,omitempty"`
	TriggerPriceType *string `json:"triggerPriceType,omitempty"`
	Price            *string `json:"price,omitempty"`
	Size             *string `json:"size,omitempty"`
}

// CreateOrder represents the result of creating an order
type CreateOrder struct {
	OrderId *string `json:"orderId,omitempty"`
}

// GetMaxCreateOrderSize represents max order size information
type GetMaxCreateOrderSize struct {
	MaxBuySize *string `json:"maxBuySize,omitempty"`
	MaxSellSize *string `json:"maxSellSize,omitempty"`
	Ask1Price *string `json:"ask1Price,omitempty"`
	Bid1Price *string `json:"bid1Price,omitempty"`
}

// OrderFillTransaction represents an order fill transaction
type OrderFillTransaction struct {
	Id              *string `json:"id,omitempty"`
	OrderId         *string `json:"orderId,omitempty"`
	UserId          *string `json:"userId,omitempty"`
	AccountId       *string `json:"accountId,omitempty"`
	CoinId          *string `json:"coinId,omitempty"`
	ContractId      *string `json:"contractId,omitempty"`
	Side            *string `json:"side,omitempty"`
	FillPrice       *string `json:"fillPrice,omitempty"`
	FillSize        *string `json:"fillSize,omitempty"`
	FillValue       *string `json:"fillValue,omitempty"`
	FillFee         *string `json:"fillFee,omitempty"`
	FillType        *string `json:"fillType,omitempty"`
	MatchSequenceId *string `json:"matchSequenceId,omitempty"`
	CreatedTime     *string `json:"createdTime,omitempty"`
}

// Order represents an order
type Order struct {
	Id                        *string      `json:"id,omitempty"`
	UserId                    *string      `json:"userId,omitempty"`
	AccountId                 *string      `json:"accountId,omitempty"`
	CoinId                    *string      `json:"coinId,omitempty"`
	ContractId                *string      `json:"contractId,omitempty"`
	Side                      *string      `json:"side,omitempty"`
	Price                     *string      `json:"price,omitempty"`
	Size                      *string      `json:"size,omitempty"`
	ClientOrderId             *string      `json:"clientOrderId,omitempty"`
	Type                      *string      `json:"type,omitempty"`
	TimeInForce               *string      `json:"timeInForce,omitempty"`
	ReduceOnly                *bool        `json:"reduceOnly,omitempty"`
	TriggerPrice              *string      `json:"triggerPrice,omitempty"`
	TriggerPriceType          *string      `json:"triggerPriceType,omitempty"`
	ExpireTime                *string      `json:"expireTime,omitempty"`
	SourceKey                 *string      `json:"sourceKey,omitempty"`
	IsPositionTpsl            *bool        `json:"isPositionTpsl,omitempty"`
	IsLiquidate               *bool        `json:"isLiquidate,omitempty"`
	IsDeleverage              *bool        `json:"isDeleverage,omitempty"`
	OpenTpslParentOrderId     *string      `json:"openTpslParentOrderId,omitempty"`
	IsSetOpenTp               *bool        `json:"isSetOpenTp,omitempty"`
	OpenTp                    *OpenTpSl    `json:"openTp,omitempty"`
	IsSetOpenSl               *bool        `json:"isSetOpenSl,omitempty"`
	OpenSl                    *OpenTpSl    `json:"openSl,omitempty"`
	IsWithoutMatch            *bool        `json:"isWithoutMatch,omitempty"`
	WithoutMatchFillSize      *string      `json:"withoutMatchFillSize,omitempty"`
	WithoutMatchFillValue     *string      `json:"withoutMatchFillValue,omitempty"`
	WithoutMatchPeerAccountId *string      `json:"withoutMatchPeerAccountId,omitempty"`
	WithoutMatchPeerOrderId   *string      `json:"withoutMatchPeerOrderId,omitempty"`
	MaxLeverage               *string      `json:"maxLeverage,omitempty"`
	TakerFeeRate              *string      `json:"takerFeeRate,omitempty"`
	MakerFeeRate              *string      `json:"makerFeeRate,omitempty"`
	LiquidateFeeRate          *string      `json:"liquidateFeeRate,omitempty"`
	MarketLimitPrice          *string      `json:"marketLimitPrice,omitempty"`
	MarketLimitValue          *string      `json:"marketLimitValue,omitempty"`
	L2Nonce                   *string      `json:"l2Nonce,omitempty"`
	L2Value                   *string      `json:"l2Value,omitempty"`
	L2Size                    *string      `json:"l2Size,omitempty"`
	L2LimitFee                *string      `json:"l2LimitFee,omitempty"`
	L2ExpireTime              *string      `json:"l2ExpireTime,omitempty"`
	L2Signature               *L2Signature `json:"l2Signature,omitempty"`
	ExtraType                 *string      `json:"extraType,omitempty"`
	ExtraDataJson             *string      `json:"extraDataJson,omitempty"`
	Status                    *string      `json:"status,omitempty"`
	MatchSequenceId           *string      `json:"matchSequenceId,omitempty"`
	TriggerTime               *string      `json:"triggerTime,omitempty"`
	TriggerPriceTime          *string      `json:"triggerPriceTime,omitempty"`
	TriggerPriceValue         *string      `json:"triggerPriceValue,omitempty"`
	CancelReason              *string      `json:"cancelReason,omitempty"`
	CumFillSize               *string      `json:"cumFillSize,omitempty"`
	CumFillValue              *string      `json:"cumFillValue,omitempty"`
	CumFillFee                *string      `json:"cumFillFee,omitempty"`
	MaxFillPrice              *string      `json:"maxFillPrice,omitempty"`
	MinFillPrice              *string      `json:"minFillPrice,omitempty"`
	CumLiquidateFee           *string      `json:"cumLiquidateFee,omitempty"`
	CumRealizePnl             *string      `json:"cumRealizePnl,omitempty"`
	CumMatchSize              *string      `json:"cumMatchSize,omitempty"`
	CumMatchValue             *string      `json:"cumMatchValue,omitempty"`
	CumMatchFee               *string      `json:"cumMatchFee,omitempty"`
	CumFailSize               *string      `json:"cumFailSize,omitempty"`
	CumFailValue              *string      `json:"cumFailValue,omitempty"`
	CumFailFee                *string      `json:"cumFailFee,omitempty"`
	CumApprovedSize           *string      `json:"cumApprovedSize,omitempty"`
	CumApprovedValue          *string      `json:"cumApprovedValue,omitempty"`
	CumApprovedFee            *string      `json:"cumApprovedFee,omitempty"`
	CreatedTime               *string      `json:"createdTime,omitempty"`
	UpdatedTime               *string      `json:"updatedTime,omitempty"`
}

// PageDataOrder represents paginated order data
type PageDataOrder struct {
	DataList           []Order `json:"dataList,omitempty"`
	NextPageOffsetData *string `json:"nextPageOffsetData,omitempty"`
}

// PageDataOrderFillTransaction represents paginated order fill transaction data
type PageDataOrderFillTransaction struct {
	DataList           []OrderFillTransaction `json:"dataList,omitempty"`
	NextPageOffsetData *string                `json:"nextPageOffsetData,omitempty"`
}

// Common filter types used across different order APIs
type OrderFilterParams struct {
	FilterCoinIdList     []string // Filter by coin IDs, empty means all coins
	FilterContractIdList []string // Filter by contract IDs, empty means all contracts
	FilterTypeList       []string // Filter by order types
	FilterStatusList     []string // Filter by order statuses
	FilterIsLiquidate    *bool    // Filter by liquidation status
	FilterIsDeleverage   *bool    // Filter by deleverage status
	FilterIsPositionTpsl *bool    // Filter by position take-profit/stop-loss status
}

// Common pagination parameters
type PaginationParams struct {
	Size       string // Size of the page, must be greater than 0 and less than or equal to 100/200
	OffsetData string // Offset data for pagination. Empty string gets the first page
}

// OrderFillTransactionParams represents parameters for getting order fill transactions
type OrderFillTransactionParams struct {
	PaginationParams
	OrderFilterParams
	FilterOrderIdList []string // Filter by order IDs, empty means all orders

	// Time filters
	FilterStartCreatedTimeInclusive uint64 // Filter start time (inclusive), 0 means from earliest
	FilterEndCreatedTimeExclusive   uint64 // Filter end time (exclusive), 0 means until latest
}

// GetActiveOrderParams represents parameters for getting active orders
type GetActiveOrderParams struct {
	PaginationParams
	OrderFilterParams

	// Time filters
	FilterStartCreatedTimeInclusive uint64 // Filter start time (inclusive), 0 means from earliest
	FilterEndCreatedTimeExclusive   uint64 // Filter end time (exclusive), 0 means until latest
}

// GetHistoryOrderParams represents parameters for getting historical orders
type GetHistoryOrderParams struct {
	PaginationParams
	OrderFilterParams

	// Time filters
	FilterStartCreatedTimeInclusive uint64 // Filter start time (inclusive), 0 means from earliest
	FilterEndCreatedTimeExclusive   uint64 // Filter end time (exclusive), 0 means until latest
}

// CreateOrderParams represents parameters for creating an order
type CreateOrderParams struct {
	ContractId    string    `json:"contractId"`
	Price         string    `json:"price"`
	Size          string    `json:"size"`
	Type          OrderType `json:"type"`
	Side          string    `json:"side"`
	ExpireTime    time.Time `json:"expireTime,omitempty"`
	ClientOrderId *string   `json:"clientOrderId"`
	TimeInForce   string    `json:"timeInForce,omitempty"`
	ReduceOnly    bool      `json:"reduceOnly,omitempty"`
}

// CancelOrderParams represents parameters for canceling orders
type CancelOrderParams struct {
	OrderId    string // Order ID to cancel
	ClientId   string // Client order ID to cancel
	ContractId string // Contract ID for canceling all orders
}

// ResultCreateOrder represents the result of creating an order
type ResultCreateOrder struct {
	Code       string       `json:"code"`
	Data       *CreateOrder `json:"data"`
	ErrorParam interface{}  `json:"errorParam"`
	ErrorMsg   string       `json:"msg"`
}

// ResultPageDataOrder represents paginated order data
type ResultPageDataOrder struct {
	Code       string         `json:"code"`
	Data       *PageDataOrder `json:"data"`
	ErrorParam interface{}    `json:"errorParam"`
	ErrorMsg   string         `json:"msg"`
}

// ResultPageDataOrderFillTransaction represents paginated order fill transaction data
type ResultPageDataOrderFillTransaction struct {
	Code       string                        `json:"code"`
	Data       *PageDataOrderFillTransaction `json:"data"`
	ErrorParam interface{}                   `json:"errorParam"`
	ErrorMsg   string                        `json:"msg"`
}

// ResultListOrder represents list of orders
type ResultListOrder struct {
	Code       string      `json:"code"`
	Data       []Order     `json:"data"`
	ErrorParam interface{} `json:"errorParam"`
	ErrorMsg   string      `json:"msg"`
}

// ResultGetMaxCreateOrderSize represents the result of getting max order size
type ResultGetMaxCreateOrderSize struct {
	Code       string                 `json:"code"`
	Data       *GetMaxCreateOrderSize `json:"data"`
	ErrorParam interface{}            `json:"errorParam"`
	ErrorMsg   string                 `json:"msg"`
}
