package decimal

import (
	"math/big"
)

// ConvertToWei deprecated. Instead of ConvertToWei just use decimal.ToBig()
func ConvertToWei(dec *Decimal, mantissa uint8) *big.Int {
	if dec == nil || dec.value == nil {
		return new(big.Int).SetUint64(0)
	}

	return NewDecimal(dec).ToBig()
}
