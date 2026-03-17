package internal

// L2Signature represents a Layer 2 signature
type L2Signature struct {
	R string `json:"r"`
	S string `json:"s"`
	V string `json:"v"`
}

// Order type constants
const (
	LimitOrderWithFeeType = int64(3)
	TransferType          = int64(4)
	CondTransferType      = int64(5)
	WithdrawalOrderType   = int64(6)
	WithdrawalToAddress   = int64(7)
)