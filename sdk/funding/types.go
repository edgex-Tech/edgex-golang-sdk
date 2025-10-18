package funding

// Response types for funding API

// ResultPageDataFundingRate represents paginated funding rate data
type ResultPageDataFundingRate struct {
	Code       string        `json:"code"`
	Data       *PageData     `json:"data"`
	ErrorParam []interface{} `json:"errorParam"`
}

// PageData represents pagination data
type PageData struct {
	List       []interface{} `json:"list"`
	OffsetData string        `json:"offsetData"`
}

// ResultListFundingRate represents list of funding rates
type ResultListFundingRate struct {
	Code       string        `json:"code"`
	Data       []interface{} `json:"data"`
	ErrorParam []interface{} `json:"errorParam"`
}

// Request parameter types

// GetFundingRateParams represents parameters for GetFundingRate
type GetFundingRateParams struct {
	ContractID string
	Size       *int32
	Offset     *string
	From       *int64
	To         *int64
}

// GetLatestFundingRateParams represents parameters for GetLatestFundingRate
type GetLatestFundingRateParams struct {
	ContractID string
}
