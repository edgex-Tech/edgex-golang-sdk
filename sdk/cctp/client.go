package cctp

import (
	"encoding/hex"
	"fmt"
	"math"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
)

const (
	ZeroBytes32                  = "0x" + "0000000000000000000000000000000000000000000000000000000000000000"
	DefaultIrisBaseURL           = "https://iris-api.circle.com"
	DefaultSourceCCTPDomain      = 28
	DefaultDestCCTPDomain        = 0
	DefaultEdgeUSDC              = "0x98d2919b9A214E6Fa5384AC81E6864bA686Ad74c"
	DefaultEdgeTokenMessenger    = "0x98706A006bc632Df31CAdFCBD43F38887ce2ca5c"
	DefaultETHMessageTransmitter = "0x81D40F21F12A8F0E3252Bccb954D722d4c464B64"
	DefaultStandardFinality      = 2000
	DefaultFastFinality          = 1000
	DefaultClaimGasLimit         = 500000
)

func Bytes32Address(address string) (string, error) {
	if !common.IsHexAddress(address) {
		return "", fmt.Errorf("invalid EVM address: %s", address)
	}
	checksummed := common.HexToAddress(address).Hex()
	return "0x" + strings.Repeat("00", 12) + strings.ToLower(checksummed[2:]), nil
}

func ExtractMinimumFeeBps(payload interface{}, finalityThreshold int) (decimal.Decimal, error) {
	rows := irisRows(payload)
	for _, row := range rows {
		if intValue(row["finalityThreshold"]) != finalityThreshold {
			continue
		}
		minimumFee := strings.TrimSpace(fmt.Sprint(row["minimumFee"]))
		if minimumFee == "" {
			break
		}
		return decimal.NewFromString(minimumFee)
	}
	available := make([]string, 0, len(rows))
	for _, row := range rows {
		if value, ok := row["finalityThreshold"]; ok {
			available = append(available, fmt.Sprint(value))
		}
	}
	return decimal.Decimal{}, fmt.Errorf("Circle fee response missing minimumFee for finalityThreshold=%d; available=%s", finalityThreshold, strings.Join(available, ","))
}

func CalculateFeeFromBps(amount int64, feeBps decimal.Decimal) int64 {
	fee := decimal.NewFromInt(amount).Mul(feeBps).Div(decimal.NewFromInt(10000)).Ceil()
	return fee.IntPart()
}

func IrisMessages(payload interface{}) []map[string]interface{} {
	return irisRows(payload)
}

func IsCompleteIrisMessage(message map[string]interface{}) bool {
	attestation := strings.TrimSpace(fmt.Sprint(message["attestation"]))
	return strings.EqualFold(strings.TrimSpace(fmt.Sprint(message["status"])), "complete") &&
		strings.HasPrefix(attestation, "0x") &&
		!strings.EqualFold(attestation, "pending_confirmations")
}

func SelectIrisMessage(payload interface{}, messageIndex *int, eventNonce string) (map[string]interface{}, error) {
	messages := IrisMessages(payload)
	if len(messages) == 0 {
		return nil, nil
	}
	if messageIndex != nil {
		if *messageIndex < 0 || *messageIndex >= len(messages) {
			return nil, fmt.Errorf("Iris message index out of range: %d; count=%d", *messageIndex, len(messages))
		}
		return messages[*messageIndex], nil
	}
	if strings.TrimSpace(eventNonce) != "" {
		target := strings.ToLower(strings.TrimSpace(eventNonce))
		for _, message := range messages {
			if strings.ToLower(strings.TrimSpace(fmt.Sprint(message["eventNonce"]))) == target {
				return message, nil
			}
		}
		return nil, fmt.Errorf("Iris response missing eventNonce=%s", eventNonce)
	}
	for _, message := range messages {
		if IsCompleteIrisMessage(message) {
			return message, nil
		}
	}
	return messages[0], nil
}

func irisRows(payload interface{}) []map[string]interface{} {
	rows := make([]map[string]interface{}, 0)
	switch value := payload.(type) {
	case []map[string]interface{}:
		return value
	case []interface{}:
		for _, item := range value {
			if row, ok := item.(map[string]interface{}); ok {
				rows = append(rows, row)
			}
		}
	case map[string]interface{}:
		if directMessages, ok := value["messages"].([]map[string]interface{}); ok {
			return directMessages
		}
		if messages, ok := value["messages"].([]interface{}); ok {
			for _, item := range messages {
				if row, ok := item.(map[string]interface{}); ok {
					rows = append(rows, row)
				}
			}
			return rows
		}
		if dataRows, ok := value["data"].([]interface{}); ok {
			for _, item := range dataRows {
				if row, ok := item.(map[string]interface{}); ok {
					rows = append(rows, row)
				}
			}
		}
		if feeRows, ok := value["fees"].([]interface{}); ok {
			for _, item := range feeRows {
				if row, ok := item.(map[string]interface{}); ok {
					rows = append(rows, row)
				}
			}
		}
		if nested, ok := value["data"].(map[string]interface{}); ok {
			rows = append(rows, nested)
		} else if len(value) > 0 {
			rows = append(rows, value)
		}
	}
	return rows
}

func intValue(value interface{}) int {
	switch v := value.(type) {
	case int:
		return v
	case int32:
		return int(v)
	case int64:
		return int(v)
	case float64:
		return int(math.Round(v))
	case string:
		var parsed int
		fmt.Sscanf(v, "%d", &parsed)
		return parsed
	default:
		return 0
	}
}

func hexBytes(input string) ([]byte, error) {
	return hex.DecodeString(strings.TrimPrefix(input, "0x"))
}
