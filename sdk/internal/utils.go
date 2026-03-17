package internal

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

func GenerateUUID() string {
	return uuid.New().String()
}

// CalcNonce calculates a nonce from source string
func CalcNonce(src string) int64 {
	h := sha256.New()
	h.Write([]byte(src))
	hash := fmt.Sprintf("%x", h.Sum(nil))

	result, _ := big.NewInt(0).SetString(string(hash[:8]), 16)
	return result.Int64()
}

// JoinStrings joins a slice of strings with commas
func JoinStrings(strs []string) string {
	return strings.Join(strs, ",")
}

// GetValue converts a JSON value to a string representation
func GetValue(value interface{}) string {
	if value == nil {
		return ""
	}

	switch v := value.(type) {
	case string:
		return v
	case []interface{}:
		if len(v) == 0 {
			return ""
		}
		values := make([]string, len(v))
		for i, item := range v {
			values[i] = GetValue(item)
		}
		return strings.Join(values, "&")
	case map[string]interface{}:
		sortedMap := make(map[string]string)
		for key, val := range v {
			sortedMap[key] = GetValue(val)
		}

		// Get sorted keys
		keys := make([]string, 0, len(sortedMap))
		for k := range sortedMap {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		// Build key=value pairs
		pairs := make([]string, len(keys))
		for i, key := range keys {
			pairs[i] = key + "=" + sortedMap[key]
		}
		return strings.Join(pairs, "&")
	default:
		// Handle other primitive types
		return fmt.Sprint(v)
	}
}

func GetRandomClientId() string {
	nanoTimestamp := time.Now().UnixNano()
	return strconv.FormatInt(nanoTimestamp, 10)
}

func ToBigInt(number string) *big.Int {
	if number == "" {
		return big.NewInt(0)
	}
	if strings.HasPrefix(number, "0x") {
		val, _ := new(big.Int).SetString(number[2:], 16)
		return val
	}
	val, _ := new(big.Int).SetString(number, 10)
	return val
}

func HexToBigInteger(hex string) (*big.Int, error) {
	if len(hex) > 2 && hex[:2] == "0x" {
		hex = hex[2:]
	}
	result := new(big.Int)
	result, ok := result.SetString(hex, 16)
	if !ok {
		return nil, fmt.Errorf("invalid hex string: %s", hex)
	}
	return result, nil
}
