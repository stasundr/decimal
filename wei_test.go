package decimal

import (
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestConvertToWei(t *testing.T) {
	actual, ok := NewDecimalFromString("1")
	assert.True(t, ok)
	assert.Equal(t, new(big.Int).SetUint64(1), ConvertToWei(actual, 0))

	actual, ok = NewDecimalFromString("1")
	assert.True(t, ok)
	assert.Equal(t, new(big.Int).SetUint64(10000000000), ConvertToWei(actual, 10))

	actual, ok = NewDecimalFromString("999999999999999999")
	assert.True(t, ok)
	expected, ok := new(big.Int).SetString("9999999999999999990000000000", 10)
	assert.True(t, ok)
	assert.Equal(t, expected, ConvertToWei(actual, 10))

	actual, ok = NewDecimalFromString("115792089237316195423570985008687907853269984665640564039457584007913129639935")
	assert.True(t, ok)
	expected, ok = new(big.Int).SetString("115792089237316195423570985008687907853269984665640564039457584007913129639935", 10)
	assert.True(t, ok)
	assert.Equal(t, expected, ConvertToWei(actual, 0))

	actual, ok = NewDecimalFromString("115792089237316195423570985008687907853269984665640564039457584007913129639935")
	assert.True(t, ok)
	expected, ok = new(big.Int).SetString("ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff", 16)
	assert.True(t, ok)
	assert.Equal(t, expected, ConvertToWei(actual, 0))

	actual, ok = NewDecimalFromString("10.10001")
	assert.True(t, ok)
	expected, ok = new(big.Int).SetString("1010001", 10)
	assert.True(t, ok)
	assert.Equal(t, expected, ConvertToWei(actual, 5))

	actual, ok = NewDecimalFromString("10.10001111111999999999901")
	assert.True(t, ok)
	expected, ok = new(big.Int).SetString("1010001111111999999999901", 10)
	assert.True(t, ok)
	assert.Equal(t, expected, ConvertToWei(actual, 23))

	actual, ok = NewDecimalFromString("10.1000111111199999999990112312313123123123")
	assert.True(t, ok)
	expected, ok = new(big.Int).SetString("101000111111199999999990112312313123123123", 10)
	assert.True(t, ok)
	assert.Equal(t, expected, ConvertToWei(actual, 40))

	actual, ok = NewDecimalFromString("115792089237316195423570985008687907853269984665640564039457.584007913129639935")
	assert.True(t, ok)
	expected, ok = new(big.Int).SetString("ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff", 16)
	assert.True(t, ok)
	assert.Equal(t, expected, ConvertToWei(actual, 18))

	actual, ok = NewDecimalFromString("115792089237316195423570985008687907853269984665640564039457584007913129639935")
	assert.True(t, ok)
	expected, ok = new(big.Int).SetString("ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff", 16)
	assert.True(t, ok)
	actualWei, overflow := ConvertToWeiOverflow(actual, 18)
	assert.True(t, overflow)
	assert.Nil(t, actualWei)

	actual, ok = NewDecimalFromString("115792089237316195423570985008687907853269984665640564039457.584007913129639936")
	assert.False(t, ok)
	assert.Nil(t, actual)

	actual, ok = NewDecimalFromString("115792089237316195423570985008687907853269984665640564039457.584007913129639935")
	assert.True(t, ok)
	expected, ok = new(big.Int).SetString("ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff", 16)
	assert.True(t, ok)
	actualWei, overflow = ConvertToWeiOverflow(actual, 18)
	assert.False(t, overflow)
	assert.Equal(t, expected, actualWei)
}
