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

func ConvertToWeiOverflow(dec *Decimal, mantissa uint8) (wei *big.Int, overflow bool) {
	if dec == nil || dec.value == nil {
		return new(big.Int).SetUint64(0), false
	}

	value, isOverflow := NewDecimal(dec).RescaleOverflow(mantissa)
	if isOverflow {
		return nil, isOverflow
	}

	return value.ToBig(), false
}
