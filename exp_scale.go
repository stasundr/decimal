package decimal

import "github.com/holiman/uint256"

var ten = uint256.NewInt(10)

// return 10**exp
func ExpScale(exp int16) *uint256.Int {
	expScale := uint256.NewInt(0).Set(ten)
	expInt := uint256.NewInt(uint64(exp))
	if exp < 0 {
		expInt.Neg(expInt)
	}

	return expScale.Exp(expScale, expInt)
}

// return 10**exp
func ExpScaleFast(exp int16) *uint256.Int {
	expScale := uint256.NewInt(0).Set(ten)
	expInt := uint256.NewInt(0).SetUint64(uint64(exp))
	if exp < 0 {
		expInt.Neg(expInt)
	}

	return expScale.Exp(expScale, expInt)
}
