package decimal

import (
	"math/big"
)

func ConvertToWei(dec *Decimal, mantissa uint8) *big.Int {
	if dec == nil || dec.value == nil {
		return new(big.Int).SetUint64(0)
	}

	return NewDecimal(dec).Rescale(mantissa).ToBig()
}
