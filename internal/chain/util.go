package chain

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

func IsValidAddress(address string, checksummed bool) bool {
	if !common.IsHexAddress(address) {
		return false
	}
	return !checksummed || common.HexToAddress(address).Hex() == address
}

func EtherToWei(amount float64) *big.Int {
	base := new(big.Float).SetFloat64(amount)
	mul := new(big.Float).Mul(base, new(big.Float).SetFloat64(1000000000000000000))
	wei, _ := mul.Int(nil)
	return wei
}
