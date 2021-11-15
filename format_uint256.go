package decimal

import (
	"strings"

	"github.com/holiman/uint256"
)

// todo: faster uint256 format
func FormatUint256(number *uint256.Int, mantissa int) string {
	if number == nil {
		return "0"
	}
	if number.IsZero() {
		return "0"
	}

	stringNumber := number.ToBig().String()
	if mantissa == 0 {
		return stringNumber
	}

	if stringNumber == "" {
		return "0"
	}

	decimalsDiff := len(stringNumber) - mantissa

	if decimalsDiff > 0 {
		var str strings.Builder
		str.Grow(len(stringNumber) + 1)
		str.WriteString(stringNumber[0:decimalsDiff])

		zerosStart := len(stringNumber) - 1
		for zerosStart > decimalsDiff-1 && stringNumber[zerosStart] == '0' {
			zerosStart--
		}
		if zerosStart > decimalsDiff-1 {
			str.WriteString(".")
			str.WriteString(stringNumber[decimalsDiff : zerosStart+1])
		}

		return str.String()
	}
	if decimalsDiff < 0 {
		var zerosCount = mantissa - len(stringNumber)
		var str strings.Builder
		str.Grow(len(stringNumber) + 2 + zerosCount)
		str.WriteString("0.")
		for i := 0; i < zerosCount; i++ {
			str.WriteString("0")
		}

		zerosStart := len(stringNumber) - 1
		for zerosStart >= 0 && stringNumber[zerosStart] == '0' {
			zerosStart--
		}
		if zerosStart > -1 {
			str.WriteString(stringNumber[:zerosStart+1])
		}
		//str.WriteString(stringNumber[:zerosStart+])

		return str.String()
	}

	var str strings.Builder
	str.Grow(len(stringNumber) + 2)
	str.WriteString("0")
	zerosStart := len(stringNumber) - 1
	for stringNumber[zerosStart] == '0' && zerosStart >= 0 {
		zerosStart--
	}
	if zerosStart > -1 {
		str.WriteString(".")
		str.WriteString(stringNumber[:zerosStart+1])
	}
	return str.String()
}
