package funding

// FundingRate represents funding rate data
type FundingRate struct {
	ContractId               *string `json:"contractId,omitempty"`
	FundingTime              *string `json:"fundingTime,omitempty"`
	FundingTimestamp         *string `json:"fundingTimestamp,omitempty"`
	OraclePrice              *string `json:"oraclePrice,omitempty"`
	IndexPrice               *string `json:"indexPrice,omitempty"`
	FundingRate              *string `json:"fundingRate,omitempty"`
	IsSettlement             *bool   `json:"isSettlement,omitempty"`
	ForecastFundingRate      *string `json:"forecastFundingRate,omitempty"`
	PreviousFundingRate      *string `json:"previousFundingRate,omitempty"`
	PreviousFundingTimestamp *string `json:"previousFundingTimestamp,omitempty"`
	PremiumIndex             *string `json:"premiumIndex,omitempty"`
	AvgPremiumIndex          *string `json:"avgPremiumIndex,omitempty"`
	PremiumIndexTimestamp    *string `json:"premiumIndexTimestamp,omitempty"`
	ImpactMarginNotional     *string `json:"impactMarginNotional,omitempty"`
	ImpactAskPrice           *string `json:"impactAskPrice,omitempty"`
	ImpactBidPrice           *string `json:"impactBidPrice,omitempty"`
	InterestRate             *string `json:"interestRate,omitempty"`
	PredictedFundingRate     *string `json:"predictedFundingRate,omitempty"`
	FundingRateIntervalMin   *string `json:"fundingRateIntervalMin,omitempty"`
}

// PageDataFundingRate represents paginated funding rate data
type PageDataFundingRate struct {
	DataList           []FundingRate `json:"dataList,omitempty"`
	NextPageOffsetData *string       `json:"nextPageOffsetData,omitempty"`
}

// Response types for funding API

// ResultPageDataFundingRate represents paginated funding rate data
type ResultPageDataFundingRate struct {
	Code       string               `json:"code"`
	Data       *PageDataFundingRate `json:"data"`
	ErrorParam interface{}          `json:"errorParam"`
	ErrorMsg   string               `json:"msg"`
}

// ResultListFundingRate represents list of funding rates
type ResultListFundingRate struct {
	Code       string        `json:"code"`
	Data       []FundingRate `json:"data"`
	ErrorParam interface{}   `json:"errorParam"`
	ErrorMsg   string        `json:"msg"`
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
