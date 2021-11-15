package decimal

import (
	"fmt"
	"testing"

	"github.com/holiman/uint256"
	"github.com/stretchr/testify/assert"
)

func TestDecimal_RescaleAndTobBigString(t *testing.T) {
	x, ok := NewDecimalFromString("1")
	assert.True(t, ok)
	assert.Equal(t, "1", x.ToBig().String())
	assert.Equal(t, "100000", x.Rescale(5).ToBig().String())

	x, ok = NewDecimalFromString("1.0000000")
	assert.True(t, ok)
	assert.Equal(t, "1", x.ToBig().String())
	assert.Equal(t, "100000", x.Rescale(5).ToBig().String())

	x, ok = NewDecimalFromString("1.00000001")
	assert.True(t, ok)
	assert.Equal(t, uint8(8), x.GetMantissa())
	assert.Equal(t, "1.00000001", x.String())
	assert.Equal(t, "1.00000001", x.Rescale(12).String())
	assert.Equal(t, "1", x.Rescale(7).String())

	assert.Equal(t, "10000000", x.ToBig().String())
	assert.Equal(t, "1", x.Rescale(5).String())
	assert.Equal(t, "100000", x.Rescale(5).ToBig().String())
}

func TestDecimal_TobBig(t *testing.T) {
	x, ok := NewDecimalFromString("1.0000010000")
	assert.True(t, ok)
	assert.Equal(t, "1000001", x.ToBig().String())
	assert.Equal(t, uint8(6), x.GetMantissa())

	x, ok = NewDecimalFromString("1.000001")
	assert.True(t, ok)
	assert.Equal(t, "1000001", x.ToBig().String())
	assert.Equal(t, uint8(6), x.GetMantissa())

	x, ok = NewDecimalFromString("1.1")
	assert.True(t, ok)
	assert.Equal(t, "11", x.ToBig().String())
	assert.Equal(t, uint8(1), x.GetMantissa())

	x, ok = NewDecimalFromString("0")
	assert.True(t, ok)
	assert.Equal(t, "0", x.ToBig().String())
	assert.Equal(t, uint8(0), x.GetMantissa())

	x, ok = NewDecimalFromString("0.0000000100000")
	assert.True(t, ok)
	assert.Equal(t, "1", x.ToBig().String())
	assert.Equal(t, uint8(8), x.GetMantissa())

	x, ok = NewDecimalFromString("0.0000000100000")
	assert.True(t, ok)
	assert.Equal(t, "1", x.ToBig().String())
	assert.Equal(t, uint8(8), x.GetMantissa())

	x, ok = NewDecimalFromString("10")
	assert.True(t, ok)
	assert.Equal(t, "10", x.ToBig().String())
	assert.Equal(t, uint8(0), x.GetMantissa())

	x, ok = NewDecimalFromString("10.0")
	assert.True(t, ok)
	assert.Equal(t, "10", x.ToBig().String())
	assert.Equal(t, uint8(0), x.GetMantissa())

	x, ok = NewDecimalFromString("10.00000")
	assert.True(t, ok)
	assert.Equal(t, "10", x.ToBig().String())
	assert.Equal(t, uint8(0), x.GetMantissa())

	x, ok = NewDecimalFromString("10.000001")
	assert.True(t, ok)
	assert.Equal(t, "10000001", x.ToBig().String())
	assert.Equal(t, uint8(6), x.GetMantissa())

	x, ok = NewDecimalFromString("10.")
	assert.True(t, ok)
	assert.Equal(t, "10", x.ToBig().String())
	assert.Equal(t, uint8(0), x.GetMantissa())

	x, ok = NewDecimalFromString("0.00000000")
	assert.True(t, ok)
	assert.Equal(t, "0", x.ToBig().String())
	assert.Equal(t, uint8(0), x.GetMantissa())
}

func TestDecimal_Format(t *testing.T) {
	x, ok := NewDecimalFromString("1.0000010000")
	assert.True(t, ok)
	assert.Equal(t, "1.000001", x.String())

	x, ok = NewDecimalFromString("0.0000010000")
	assert.True(t, ok)
	assert.Equal(t, "0.000001", x.String())

	x, ok = NewDecimalFromString("12222.0000010000")
	assert.True(t, ok)
	assert.Equal(t, "12222.000001", x.String())

	x, ok = NewDecimalFromString("12222.000000000000")
	assert.True(t, ok)
	assert.Equal(t, "12222", x.String())

	x, ok = NewDecimalFromString("1.000000000000")
	assert.True(t, ok)
	assert.Equal(t, "1", x.String())

	x, ok = NewDecimalFromString("1.0")
	assert.True(t, ok)
	assert.Equal(t, "1", x.String())

	x, ok = NewDecimalFromString("1.00")
	assert.True(t, ok)
	assert.Equal(t, "1", x.String())

	x, ok = NewDecimalFromString("10.00")
	assert.True(t, ok)
	assert.Equal(t, "10", x.String())

	x, ok = NewDecimalFromString("0.00")
	assert.True(t, ok)
	assert.Equal(t, "0", x.String())

	x, ok = NewDecimalFromString("0.00001")
	assert.True(t, ok)
	assert.Equal(t, "0.00001", x.String())

	x, ok = NewDecimalFromString("10000000000")
	assert.True(t, ok)
	assert.Equal(t, "10000000000", x.String())

	x, ok = NewDecimalFromString("1221")
	assert.True(t, ok)
	assert.Equal(t, "1221", x.String())

	x, ok = NewDecimalFromString("101010.00000000000")
	assert.True(t, ok)
	assert.Equal(t, "101010", x.String())

	x, ok = NewDecimalFromString("0")
	assert.True(t, ok)
	assert.Equal(t, "0", x.String())

	x, ok = NewDecimalFromString("0.")
	assert.True(t, ok)
	assert.Equal(t, "0", x.String())

	x, ok = NewDecimalFromString("110000000000.0000000")
	assert.True(t, ok)
	assert.Equal(t, "110000000000", x.String())

}

func TestDecimal_Add(t *testing.T) {
	x := NewDecimalFromUint256(uint256.NewInt(0).SetUint64(10), 0)
	y := NewDecimalFromUint256(uint256.NewInt(0).SetUint64(10), 0)
	z := NewDecimalFromUint256(uint256.NewInt(0).SetUint64(20), 0)
	if !x.Add(y).Eq(z) {
		t.Fatalf("x add y is not equal z")
	}

	x = NewDecimalFromUint256(uint256.NewInt(0).SetUint64(10), 1)
	y = NewDecimalFromUint256(uint256.NewInt(0).SetUint64(10), 1)
	z = NewDecimalFromUint256(uint256.NewInt(0).SetUint64(20), 1)
	if !x.Add(y).Eq(z) {
		t.Fatalf("x add y is not equal z")
	}

	x = NewDecimalFromUint256(uint256.NewInt(0).SetUint64(10), 2) // 0.10
	y = NewDecimalFromUint256(uint256.NewInt(0).SetUint64(10), 1) // 1
	z = NewDecimalFromUint256(uint256.NewInt(0).SetUint64(11), 1) // 1.1
	if !x.Add(y).Eq(z) {
		t.Fatalf("x add y is not equal z")
	}

	// from strings
	var ok bool
	x, ok = NewDecimalFromString("0.000000001")
	if !ok {
		t.Fatal("error convert from string")
	}
	y, ok = NewDecimalFromString("0.000000001")
	if !ok {
		t.Fatal("error convert from string")
	}
	z, ok = NewDecimalFromString("0.000000002")
	if !ok {
		t.Fatal("error convert from string")
	}
	if !x.Add(y).Eq(z) {
		t.Fatalf("x is not equal y")
	}

	// from strings
	x, ok = NewDecimalFromString("0.0000000001")
	if !ok {
		t.Fatal("error convert from string")
	}
	y, ok = NewDecimalFromString("0.000000001")
	if !ok {
		t.Fatal("error convert from string")
	}
	z, ok = NewDecimalFromString("0.0000000011")
	if !ok {
		t.Fatal("error convert from string")
	}
	if !x.Add(y).Eq(z) {
		t.Fatalf("x is not equal y")
	}

	// from strings
	x, ok = NewDecimalFromString("10000000000000")
	if !ok {
		t.Fatal("error convert from string")
	}
	y, ok = NewDecimalFromString("10000000000001.00001")
	if !ok {
		t.Fatal("error convert from string")
	}
	z, ok = NewDecimalFromString("20000000000001.00001")
	if !ok {
		t.Fatal("error convert from string")
	}
	if !x.Add(y).Eq(z) {
		t.Fatalf("x is not equal y")
	}

	// from strings
	x, ok = NewDecimalFromString("10000000000000")
	if !ok {
		t.Fatal("error convert from string")
	}
	y, ok = NewDecimalFromString("0.0000000000000001")
	if !ok {
		t.Fatal("error convert from string")
	}
	z, ok = NewDecimalFromString("10000000000000.0000000000000001")
	if !ok {
		t.Fatal("error convert from string")
	}
	if !x.Add(y).Eq(z) {
		t.Fatalf("x is not equal y")
	}
}

func TestDecimal_Sub(t *testing.T) {
	x := NewDecimalFromUint256(uint256.NewInt(0).SetUint64(30), 0)
	y := NewDecimalFromUint256(uint256.NewInt(0).SetUint64(10), 0)
	z := NewDecimalFromUint256(uint256.NewInt(0).SetUint64(20), 0)
	if !x.Sub(y).Eq(z) {
		t.Fatalf("x add y is not equal z")
	}

	var ok bool

	// from strings
	x, ok = NewDecimalFromString("30")
	if !ok {
		t.Fatal("error convert from string")
	}
	y, ok = NewDecimalFromString("20.5")
	if !ok {
		t.Fatal("error convert from string")
	}
	z, ok = NewDecimalFromString("9.5")
	if !ok {
		t.Fatal("error convert from string")
	}
	if !x.Sub(y).Eq(z) {
		t.Fatalf("x is not equal y")
	}
}

func TestDecimal_Mul(t *testing.T) {
	x := NewDecimalFromUint256(uint256.NewInt(0).SetUint64(10), 0)
	y := NewDecimalFromUint256(uint256.NewInt(0).SetUint64(10), 0)
	z := NewDecimalFromUint256(uint256.NewInt(0).SetUint64(100), 0)
	if !x.Mul(y).Eq(z) {
		t.Fatalf("x add y is not equal z")
	}

	x = NewDecimalFromUint256(uint256.NewInt(0).SetUint64(10), 1)
	y = NewDecimalFromUint256(uint256.NewInt(0).SetUint64(10), 1)
	z = NewDecimalFromUint256(uint256.NewInt(0).SetUint64(100), 2)
	if !x.Mul(y).Eq(z) {
		t.Fatalf("x add y is not equal z")
	}

	// from string
	var ok bool
	x, ok = NewDecimalFromString("10001")
	if !ok {
		t.Fatal("error convert from string")
	}
	y, ok = NewDecimalFromString("0.1")
	if !ok {
		t.Fatal("error convert from string")
	}
	z, ok = NewDecimalFromString("1000.1")
	if !ok {
		t.Fatal("error convert from string")
	}
	if !x.Mul(y).Eq(z) {
		t.Fatalf("x is not equal y")
	}

	// from string
	x, ok = NewDecimalFromString("0.00001")
	if !ok {
		t.Fatal("error convert from string")
	}
	y, ok = NewDecimalFromString("0.00001")
	if !ok {
		t.Fatal("error convert from string")
	}
	z, ok = NewDecimalFromString("0.0000000001")
	if !ok {
		t.Fatal("error convert from string")
	}
	if !x.Mul(y).Eq(z) {
		t.Fatalf("x is not equal y")
	}

	// from string
	x, ok = NewDecimalFromString("0.000000000000000001")
	if !ok {
		t.Fatal("error convert from string")
	}
	y, ok = NewDecimalFromString("0.000000000000000001")
	if !ok {
		t.Fatal("error convert from string")
	}
	z, ok = NewDecimalFromString("0.000000000000000000000000000000000001")
	if !ok {
		t.Fatal("error convert from string")
	}
	if !x.Mul(y).Eq(z) {
		t.Fatalf("x is not equal y")
	}
}

func TestDecimal_Div(t *testing.T) {
	x := NewDecimalFromUint256(uint256.NewInt(0).SetUint64(10), 0)
	y := NewDecimalFromUint256(uint256.NewInt(0).SetUint64(10), 0)
	z := NewDecimalFromUint256(uint256.NewInt(0).SetUint64(1), 0)
	if !x.Div(y).Eq(z) {
		t.Fatalf("x is not equal y: %s: %s", x.String(), z.String())
	}

	// from string
	var ok bool
	x, ok = NewDecimalFromString("100")
	if !ok {
		t.Fatal("error convert from string")
	}
	y, ok = NewDecimalFromString("10")
	if !ok {
		t.Fatal("error convert from string")
	}
	z, ok = NewDecimalFromString("10")
	if !ok {
		t.Fatal("error convert from string")
	}
	if !x.Div(y).Eq(z) {
		t.Fatalf("x is not equal y")
	}

	//
	x, ok = NewDecimalFromString("10")
	if !ok {
		t.Fatal("error convert from string")
	}
	y, ok = NewDecimalFromString("100")
	if !ok {
		t.Fatal("error convert from string")
	}
	z, ok = NewDecimalFromString("0.1")
	if !ok {
		t.Fatal("error convert from string")
	}
	if !x.Div(y).Eq(z) {
		t.Fatalf("x is not equal y")
	}

	// from string
	x, ok = NewDecimalFromString("1001")
	if !ok {
		t.Fatal("error convert from string")
	}
	y, ok = NewDecimalFromString("10.4")
	if !ok {
		t.Fatal("error convert from string")
	}
	z, ok = NewDecimalFromString("96.25")
	if !ok {
		t.Fatal("error convert from string")
	}
	a := x.Div(y)
	if !a.Eq(z) {
		t.Fatalf("x is not equal y: %s: %s", a.String(), z.String())
	}

	// from string
	x, ok = NewDecimalFromString("0.5")
	if !ok {
		t.Fatal("error convert from string")
	}
	y, ok = NewDecimalFromString("0.8")
	if !ok {
		t.Fatal("error convert from string")
	}
	z, ok = NewDecimalFromString("0.625")
	if !ok {
		t.Fatal("error convert from string")
	}
	a = x.Div(y)
	if !a.Eq(z) {
		t.Fatalf("x is not equal y: %s: %s", a.String(), z.String())
	}

	// from string
	x, ok = NewDecimalFromString("0.005")
	if !ok {
		t.Fatal("error convert from string")
	}
	y, ok = NewDecimalFromString("0.8")
	if !ok {
		t.Fatal("error convert from string")
	}
	z, ok = NewDecimalFromString("0.00625")
	if !ok {
		t.Fatal("error convert from string")
	}
	a = x.Div(y)
	if !a.Eq(z) {
		t.Fatalf("x is not equal y: %s: %s", a.String(), z.String())
	}

	// from string
	x, ok = NewDecimalFromString("100000000")
	if !ok {
		t.Fatal("error convert from string")
	}
	y, ok = NewDecimalFromString("1000")
	if !ok {
		t.Fatal("error convert from string")
	}
	z, ok = NewDecimalFromString("100000")
	if !ok {
		t.Fatal("error convert from string")
	}
	a = x.Div(y)
	if !a.Eq(z) {
		t.Fatalf("x is not equal y: %s: %s", a.String(), z.String())
	}

	// from string
	x, ok = NewDecimalFromString("1000")
	if !ok {
		t.Fatal("error convert from string")
	}
	y, ok = NewDecimalFromString("100000000")
	if !ok {
		t.Fatal("error convert from string")
	}
	z, ok = NewDecimalFromString("0.00001")
	if !ok {
		t.Fatal("error convert from string")
	}
	a = x.Div(y)
	if !a.Eq(z) {
		t.Fatalf("x is not equal y: %s: %s", a.String(), z.String())
	}

	// from string
	x, ok = NewDecimalFromString("1000.5")
	if !ok {
		t.Fatal("error convert from string")
	}
	y, ok = NewDecimalFromString("100000000")
	if !ok {
		t.Fatal("error convert from string")
	}
	z, ok = NewDecimalFromString("0.000010005")
	if !ok {
		t.Fatal("error convert from string")
	}
	a = x.Div(y)
	if !a.Eq(z) {
		t.Fatalf("x is not equal y: %s: %s", a.String(), z.String())
	}

	// from string
	x, ok = NewDecimalFromString("2")
	if !ok {
		t.Fatal("error convert from string")
	}
	y, ok = NewDecimalFromString("3")
	if !ok {
		t.Fatal("error convert from string")
	}
	z, ok = NewDecimalFromString("0.66666666")
	if !ok {
		t.Fatal("error convert from string")
	}
	a = x.Div(y)
	a.Rescale(8)
	if !a.Eq(z) {
		t.Fatalf("x is not equal y: %s: %s", a.String(), z.String())
	}

	// from string
	x, ok = NewDecimalFromString("2.1")
	if !ok {
		t.Fatal("error convert from string")
	}
	y, ok = NewDecimalFromString("3.1112321312312123")
	if !ok {
		t.Fatal("error convert from string")
	}
	z, ok = NewDecimalFromString("0.67497374")
	if !ok {
		t.Fatal("error convert from string")
	}
	a = x.Div(y)
	a.Rescale(8)
	if !a.Eq(z) {
		t.Fatalf("x is not equal y: %s: %s", a.String(), z.String())
	}

	// from string
	x, ok = NewDecimalFromString("3.1112321312312123")
	if !ok {
		t.Fatal("error convert from string")
	}
	y, ok = NewDecimalFromString("2.11")
	if !ok {
		t.Fatal("error convert from string")
	}
	z, ok = NewDecimalFromString("1.47451759")
	if !ok {
		t.Fatal("error convert from string")
	}
	a = x.Div(y)
	a.Rescale(8)
	if !a.Eq(z) {
		t.Fatalf("x is not equal y: %s: %s", a.String(), z.String())
	}
}

func TestDecimal_Eq(t *testing.T) {
	x := NewDecimalFromUint256(uint256.NewInt(0).SetUint64(10), 0)
	y := NewDecimalFromUint256(uint256.NewInt(0).SetUint64(10), 0)
	if !x.Eq(y) {
		t.Fatalf("x is not equal y")
	}

	x = NewDecimalFromUint256(uint256.NewInt(0).SetUint64(10), 0)
	y = NewDecimalFromUint256(uint256.NewInt(0).SetUint64(20), 0)
	if x.Eq(y) {
		t.Fatalf("x is equal y")
	}

	x = NewDecimalFromUint256(uint256.NewInt(0).SetUint64(10), 2)
	y = NewDecimalFromUint256(uint256.NewInt(0).SetUint64(100), 3)
	if !x.Eq(y) {
		t.Fatalf("x is not equal y")
	}

	// from strings
	var ok bool
	x, ok = NewDecimalFromString("10")
	if !ok {
		t.Fatal("error convert from string")
	}
	y, ok = NewDecimalFromString("10")
	if !ok {
		t.Fatal("error convert from string")
	}
	if !x.Eq(y) {
		t.Fatalf("x is not equal y")
	}

	// from strings
	x, ok = NewDecimalFromString("10.1")
	if !ok {
		t.Fatal("error convert from string")
	}
	y, ok = NewDecimalFromString("10.1")
	if !ok {
		t.Fatal("error convert from string")
	}
	if !x.Eq(y) {
		t.Fatalf("x is not equal y")
	}

	// from strings
	x, ok = NewDecimalFromString("10.1")
	if !ok {
		t.Fatal("error convert from string")
	}
	y = NewDecimalFromUint256(uint256.NewInt(0).SetUint64(101), 1)
	if !ok {
		t.Fatal("error convert from string")
	}
	if !x.Eq(y) {
		t.Fatalf("x is not equal y")
	}

	// from strings
	x, ok = NewDecimalFromString("10.1111")
	if !ok {
		t.Fatal("error convert from string")
	}
	y = NewDecimalFromUint256(uint256.NewInt(0).SetUint64(1011110), 5)
	if !ok {
		t.Fatal("error convert from string")
	}
	if !x.Eq(y) {
		t.Fatalf("x is not equal y")
	}

	// from strings
	x, ok = NewDecimalFromString("0.000001")
	if !ok {
		t.Fatal("error convert from string")
	}
	y, ok = NewDecimalFromString("0.000001")
	if !ok {
		t.Fatal("error convert from string")
	}
	if !x.Eq(y) {
		t.Fatalf("x is not equal y")
	}
}

func TestExpScaleFast(t *testing.T) {
	expScale := uint256.NewInt(0).Set(ten)
	fmt.Println(expScale.Bytes(), expScale.Uint64())
	for _, n := range expScale.Bytes() {
		fmt.Printf("% 08b", n) // prints 00000000 11111101
	}
	fmt.Println("")

	expScale2 := uint256.NewInt(0).Exp(expScale, uint256.NewInt(0).SetUint64(2))
	fmt.Println(expScale2.Bytes(), expScale2.Uint64())
	for _, n := range expScale2.Bytes() {
		fmt.Printf("% 08b", n) // prints 00000000 11111101
	}
	fmt.Println("")

	expScale3 := uint256.NewInt(0).Exp(expScale, uint256.NewInt(0).SetUint64(3))
	fmt.Println(expScale3.Bytes(), expScale3.Uint64())
	for _, n := range expScale3.Bytes() {
		fmt.Printf("% 08b", n) // prints 00000000 11111101
	}
	fmt.Println("")
}
