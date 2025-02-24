package order

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
	ClientOrderId *string   `json:"clientOrderId,omitempty"`
	L2ExpireTime  *int64    `json:"l2ExpireTime,omitempty"`
	TimeInForce   string    `json:"timeInForce,omitempty"`
	ReduceOnly    bool      `json:"reduceOnly,omitempty"`
}

// CancelOrderParams represents parameters for canceling orders
type CancelOrderParams struct {
	OrderId    string // Order ID to cancel
	ClientId   string // Client order ID to cancel
	ContractId string // Contract ID for canceling all orders
}

// OrderResponse represents the response from creating an order
type OrderResponse struct {
	Code         string      `json:"code"`
	Data         interface{} `json:"data"`
	ErrorParam   interface{} `json:"errorParam,omitempty"`
	RequestTime  string      `json:"requestTime"`
	ResponseTime string      `json:"responseTime"`
	TraceId      string      `json:"traceId"`
}

// MaxOrderSizeResponse represents the response from getting max order size
type MaxOrderSizeResponse struct {
	Code         string      `json:"code"`
	Data         interface{} `json:"data"`
	ErrorParam   interface{} `json:"errorParam,omitempty"`
	RequestTime  string      `json:"requestTime"`
	ResponseTime string      `json:"responseTime"`
	TraceId      string      `json:"traceId"`
}

// OrderListResponse represents the response from getting a list of orders
type OrderListResponse struct {
	Code         string      `json:"code"`
	Data         interface{} `json:"data"`
	ErrorParam   interface{} `json:"errorParam,omitempty"`
	RequestTime  string      `json:"requestTime"`
	ResponseTime string      `json:"responseTime"`
	TraceId      string      `json:"traceId"`
}

// OrderPageResponse represents the response from getting paginated orders
type OrderPageResponse struct {
	Code         string      `json:"code"`
	Data         interface{} `json:"data"`
	ErrorParam   interface{} `json:"errorParam,omitempty"`
	RequestTime  string      `json:"requestTime"`
	ResponseTime string      `json:"responseTime"`
	TraceId      string      `json:"traceId"`
}

// OrderFillTransactionResponse represents the response from getting order fill transactions
type OrderFillTransactionResponse struct {
	Code         string      `json:"code"`
	Data         interface{} `json:"data"`
	ErrorParam   interface{} `json:"errorParam,omitempty"`
	RequestTime  string      `json:"requestTime"`
	ResponseTime string      `json:"responseTime"`
	TraceId      string      `json:"traceId"`
}

// OrderFillFilterParams represents parameters for filtering order fill transactions
type OrderFillFilterParams struct {
	OrderFilterParams
	FilterOrderIdList []string // Filter by order IDs, empty means all orders
}
