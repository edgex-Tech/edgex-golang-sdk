package quote

// Response types for quote API

// ResultGetTickerSummaryModel represents ticker summary
type ResultGetTickerSummaryModel struct {
	Code       string        `json:"code"`
	Data       interface{}   `json:"data"`
	ErrorParam []interface{} `json:"errorParam"`
}

// ResultListTicker represents list of tickers
type ResultListTicker struct {
	Code       string        `json:"code"`
	Data       []interface{} `json:"data"`
	ErrorParam []interface{} `json:"errorParam"`
}

// ResultPageDataKline represents paginated K-line data
type ResultPageDataKline struct {
	Code       string        `json:"code"`
	Data       *PageData     `json:"data"`
	ErrorParam []interface{} `json:"errorParam"`
}

// PageData represents pagination data
type PageData struct {
	List       []interface{} `json:"list"`
	OffsetData string        `json:"offsetData"`
}

// ResultListDepth represents list of depth data
type ResultListDepth struct {
	Code       string        `json:"code"`
	Data       []interface{} `json:"data"`
	ErrorParam []interface{} `json:"errorParam"`
}

// ResultListContractKline represents list of contract K-line data
type ResultListContractKline struct {
	Code       string        `json:"code"`
	Data       []interface{} `json:"data"`
	ErrorParam []interface{} `json:"errorParam"`
}

// GetCode returns the Code field value
func (r *ResultGetTickerSummaryModel) GetCode() string {
	if r == nil {
		return ""
	}
	return r.Code
}

// GetData returns the Data field value
func (r *ResultGetTickerSummaryModel) GetData() interface{} {
	if r == nil {
		return nil
	}
	return r.Data
}

// GetCode returns the Code field value
func (r *ResultListTicker) GetCode() string {
	if r == nil {
		return ""
	}
	return r.Code
}

// GetData returns the Data field value
func (r *ResultListTicker) GetData() []interface{} {
	if r == nil {
		return nil
	}
	return r.Data
}

// GetCode returns the Code field value
func (r *ResultPageDataKline) GetCode() string {
	if r == nil {
		return ""
	}
	return r.Code
}

// GetData returns the Data field value
func (r *ResultPageDataKline) GetData() *PageData {
	if r == nil {
		return nil
	}
	return r.Data
}

// GetDataList returns the list of data
func (p *PageData) GetDataList() []interface{} {
	if p == nil {
		return nil
	}
	return p.List
}

// GetCode returns the Code field value
func (r *ResultListDepth) GetCode() string {
	if r == nil {
		return ""
	}
	return r.Code
}

// GetData returns the Data field value
func (r *ResultListDepth) GetData() []interface{} {
	if r == nil {
		return nil
	}
	return r.Data
}

// GetCode returns the Code field value
func (r *ResultListContractKline) GetCode() string {
	if r == nil {
		return ""
	}
	return r.Code
}

// GetData returns the Data field value
func (r *ResultListContractKline) GetData() []interface{} {
	if r == nil {
		return nil
	}
	return r.Data
}

// Request parameter types

// GetKLineParams represents parameters for GetKLine
type GetKLineParams struct {
	ContractID string
	Interval   string
	Size       int64
	PriceType  string
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
	Interval    string
	Size        int64
	PriceType   string
	From        *int64
	To          *int64
}
