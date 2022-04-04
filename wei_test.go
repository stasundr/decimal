package decimal

import (
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestConvertToWei(t *testing.T) {
	d, ok := NewDecimalFromString("1")
	assert.True(t, ok)
	assert.Equal(t, ConvertToWei(d, 0), new(big.Int).SetUint64(1))

	d, ok = NewDecimalFromString("1")
	assert.True(t, ok)
	assert.Equal(t, ConvertToWei(d, 10), new(big.Int).SetUint64(1))

	d, ok = NewDecimalFromString("999999999999999999")
	assert.True(t, ok)
	assert.Equal(t, ConvertToWei(d, 10), new(big.Int).SetUint64(999999999999999999))

	d, ok = NewDecimalFromString("115792089237316195423570985008687907853269984665640564039457584007913129639935")
	assert.True(t, ok)
	expected, ok := new(big.Int).SetString("115792089237316195423570985008687907853269984665640564039457584007913129639935", 10)
	assert.Equal(t, ConvertToWei(d, 0), expected)

	d, ok = NewDecimalFromString("115792089237316195423570985008687907853269984665640564039457584007913129639935")
	assert.True(t, ok)
	expected, ok = new(big.Int).SetString("ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff", 16)
	assert.Equal(t, ConvertToWei(d, 0), expected)

	d, ok = NewDecimalFromString("10.10001")
	assert.True(t, ok)
	expected, ok = new(big.Int).SetString("1010001", 10)
	assert.Equal(t, ConvertToWei(d, 0), expected)

	d, ok = NewDecimalFromString("10.10001111111999999999901")
	assert.True(t, ok)
	expected, ok = new(big.Int).SetString("1010001111111999999999901", 10)
	assert.Equal(t, ConvertToWei(d, 0), expected)

	d, ok = NewDecimalFromString("10.1000111111199999999990112312313123123123")
	assert.True(t, ok)
	expected, ok = new(big.Int).SetString("101000111111199999999990112312313123123123", 10)
	assert.Equal(t, ConvertToWei(d, 0), expected)
}
