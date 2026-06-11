package internal

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	gethmath "github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

// TypedData aliases go-ethereum EIP-712 typed data structure.
type TypedData = apitypes.TypedData

// TypedDataType aliases go-ethereum EIP-712 type structure.
type TypedDataType = apitypes.Type

// TypedDataTypes aliases go-ethereum EIP-712 type dictionary.
type TypedDataTypes = apitypes.Types

// TypedDataDomain aliases go-ethereum EIP-712 domain structure.
type TypedDataDomain = apitypes.TypedDataDomain

// TypedDataMessage aliases go-ethereum EIP-712 message structure.
type TypedDataMessage = apitypes.TypedDataMessage

// NewTypedDataDomain builds an EIP-712 domain from basic inputs.
func NewTypedDataDomain(name, version, chainID, verifyingContract string) (TypedDataDomain, error) {
	domain := TypedDataDomain{
		Name:    strings.TrimSpace(name),
		Version: strings.TrimSpace(version),
	}

	chainID = strings.TrimSpace(chainID)
	if chainID != "" {
		chainIDInt := big.NewInt(0)
		base := 10
		if strings.HasPrefix(strings.ToLower(chainID), "0x") {
			base = 0
		}
		if _, ok := chainIDInt.SetString(chainID, base); !ok {
			return TypedDataDomain{}, fmt.Errorf("invalid chain id: %s", chainID)
		}
		if chainIDInt.Sign() < 0 {
			return TypedDataDomain{}, fmt.Errorf("chain id must be non-negative: %s", chainID)
		}
		domain.ChainId = new(gethmath.HexOrDecimal256)
		if chainIDInt.BitLen() > 256 {
			return TypedDataDomain{}, fmt.Errorf("chain id too large")
		}
		(*big.Int)(domain.ChainId).Set(chainIDInt)
	}

	verifyingContract = strings.TrimSpace(verifyingContract)
	if verifyingContract != "" {
		if !common.IsHexAddress(verifyingContract) {
			return TypedDataDomain{}, fmt.Errorf("invalid verifying contract address: %s", verifyingContract)
		}
		domain.VerifyingContract = common.HexToAddress(verifyingContract).Hex()
	}

	return domain, nil
}

// DeriveAddressFromPrivateKey derives EVM address from a secp256k1 private key.
func DeriveAddressFromPrivateKey(privateKeyHex string) (string, error) {
	privateKeyHex = strings.TrimSpace(strings.TrimPrefix(privateKeyHex, "0x"))
	if privateKeyHex == "" {
		return "", fmt.Errorf("private key is empty")
	}

	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return "", fmt.Errorf("invalid private key: %w", err)
	}

	return crypto.PubkeyToAddress(privateKey.PublicKey).Hex(), nil
}

// SignTypedDataWithPrivateKey signs EIP-712 typed data and returns a 0x-prefixed signature.
func SignTypedDataWithPrivateKey(privateKeyHex string, typedData TypedData) (string, error) {
	privateKeyHex = strings.TrimSpace(strings.TrimPrefix(privateKeyHex, "0x"))
	if privateKeyHex == "" {
		return "", fmt.Errorf("private key is empty")
	}

	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return "", fmt.Errorf("invalid private key: %w", err)
	}

	digest, _, err := apitypes.TypedDataAndHash(typedData)
	if err != nil {
		return "", fmt.Errorf("failed to hash typed data: %w", err)
	}

	signature, err := crypto.Sign(digest, privateKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign typed data: %w", err)
	}
	// Ethereum personal signatures generally encode v as 27/28.
	signature[64] += 27

	return "0x" + hex.EncodeToString(signature), nil
}
