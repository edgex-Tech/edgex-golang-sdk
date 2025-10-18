package transfer

// Response types for transfer API

// ResultListTransferOut represents list of transfer out records
type ResultListTransferOut struct {
	Code       string        `json:"code"`
	Data       []interface{} `json:"data"`
	ErrorParam []interface{} `json:"errorParam"`
}

// ResultListTransferIn represents list of transfer in records
type ResultListTransferIn struct {
	Code       string        `json:"code"`
	Data       []interface{} `json:"data"`
	ErrorParam []interface{} `json:"errorParam"`
}

// ResultGetTransferOutAvailableAmount represents available transfer out amount
type ResultGetTransferOutAvailableAmount struct {
	Code       string        `json:"code"`
	Data       interface{}   `json:"data"`
	ErrorParam []interface{} `json:"errorParam"`
}

// ResultCreateTransferOut represents the result of creating a transfer out
type ResultCreateTransferOut struct {
	Code       string        `json:"code"`
	Data       interface{}   `json:"data"`
	ErrorParam []interface{} `json:"errorParam"`
}

// GetCode returns the Code field value
func (r *ResultListTransferOut) GetCode() string {
	if r == nil {
		return ""
	}
	return r.Code
}

// GetData returns the Data field value
func (r *ResultListTransferOut) GetData() []interface{} {
	if r == nil {
		return nil
	}
	return r.Data
}

// GetCode returns the Code field value
func (r *ResultListTransferIn) GetCode() string {
	if r == nil {
		return ""
	}
	return r.Code
}

// GetData returns the Data field value
func (r *ResultListTransferIn) GetData() []interface{} {
	if r == nil {
		return nil
	}
	return r.Data
}

// GetCode returns the Code field value
func (r *ResultGetTransferOutAvailableAmount) GetCode() string {
	if r == nil {
		return ""
	}
	return r.Code
}

// GetData returns the Data field value
func (r *ResultGetTransferOutAvailableAmount) GetData() interface{} {
	if r == nil {
		return nil
	}
	return r.Data
}

// GetCode returns the Code field value
func (r *ResultCreateTransferOut) GetCode() string {
	if r == nil {
		return ""
	}
	return r.Code
}

// GetData returns the Data field value
func (r *ResultCreateTransferOut) GetData() interface{} {
	if r == nil {
		return nil
	}
	return r.Data
}

// Request parameter types

// GetTransferOutByIdParams represents parameters for GetTransferOutById
type GetTransferOutByIdParams struct {
	TransferId string
}

// GetTransferInByIdParams represents parameters for GetTransferInById
type GetTransferInByIdParams struct {
	TransferId string
}

// GetWithdrawAvailableAmountParams represents parameters for GetWithdrawAvailableAmount
type GetWithdrawAvailableAmountParams struct {
	CoinId string
}

// CreateTransferOutParams represents parameters for CreateTransferOut
type CreateTransferOutParams struct {
	CoinId            string
	Amount            string
	ReceiverAccountId string
	ReceiverL2Key     string
	ClientTransferId  string
	TransferReason    string
}
