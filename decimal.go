package decimal

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/holiman/uint256"
)

var Zero = NewDecimalZero()

// Support only unsigned operations
type Decimal struct {
	value    *uint256.Int
	mantissa uint8
}

func (d *Decimal) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

func (d *Decimal) UnmarshalJSON(dataJson []byte) error {
	var data string
	if err := json.Unmarshal(dataJson, &data); err != nil {
		return fmt.Errorf("error unmarshal decimal: %s: %w", string(dataJson), err)
	}

	if !d.FromString(data) {
		return fmt.Errorf("error unmarshal decimal: %s", data)
	}

	return nil
}

func (d *Decimal) MarshalYAML() (interface{}, error) {
	return d.String(), nil
}

func (d *Decimal) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var data string
	if err := unmarshal(&data); err != nil {
		return nil
	}

	if !d.FromString(data) {
		return fmt.Errorf("unmarshal decimal: %s", data)
	}

	return nil
}

// Eq return d == y
func (d *Decimal) Eq(y *Decimal) bool {
	xx := NewDecimal(d)
	yy := NewDecimal(y)

	if yy.mantissa > xx.mantissa {
		xx.Rescale(yy.mantissa)
	} else if yy.mantissa < xx.mantissa {
		yy.Rescale(xx.mantissa)
	}

	return xx.value.Eq(yy.value)
}

// Gt return d > y
func (d *Decimal) Gt(y *Decimal) bool {
	xx := NewDecimal(d)
	yy := NewDecimal(y)

	if yy.mantissa > xx.mantissa {
		xx.Rescale(yy.mantissa)
	} else if yy.mantissa < xx.mantissa {
		yy.Rescale(xx.mantissa)
	}

	return xx.value.Gt(yy.value)
}

// Lt return d < y
func (d *Decimal) Lt(y *Decimal) bool {
	xx := NewDecimal(d)
	yy := NewDecimal(y)

	if yy.mantissa > xx.mantissa {
		xx.Rescale(yy.mantissa)
	} else if yy.mantissa < xx.mantissa {
		yy.Rescale(xx.mantissa)
	}

	return xx.value.Lt(yy.value)
}

// Add d = d + y and return d
func (d *Decimal) Add(y *Decimal) *Decimal {
	xx := NewDecimal(d)
	yy := NewDecimal(y)

	if yy.mantissa > xx.mantissa {
		xx.Rescale(yy.mantissa)
	} else if yy.mantissa < xx.mantissa {
		yy.Rescale(xx.mantissa)
	}

	d.value.Add(xx.value, yy.value)
	d.mantissa = xx.mantissa

	return d
}

// AddOverflow d = d + y and return d
func (d *Decimal) AddOverflow(y *Decimal) (*Decimal, bool) {
	xx := NewDecimal(d)
	yy := NewDecimal(y)

	if yy.mantissa > xx.mantissa {
		_, overflow := xx.RescaleOverflow(yy.mantissa)
		if overflow {
			return nil, true
		}
	} else if yy.mantissa < xx.mantissa {
		_, overflow := yy.RescaleOverflow(xx.mantissa)
		if overflow {
			return nil, true
		}
	}

	_, overflow := d.value.AddOverflow(xx.value, yy.value)
	if overflow {
		return nil, true
	}
	d.mantissa = xx.mantissa

	return d, false
}

// Sub d = d - y and return d
func (d *Decimal) Sub(y *Decimal) *Decimal {
	xx := NewDecimal(d)
	yy := NewDecimal(y)

	if yy.mantissa > xx.mantissa {
		xx.Rescale(yy.mantissa)
	} else if yy.mantissa < xx.mantissa {
		yy.Rescale(xx.mantissa)
	}

	d.value.Sub(xx.value, yy.value)
	d.mantissa = xx.mantissa

	return d
}

// SubOverflow d = d - y and return d
func (d *Decimal) SubOverflow(y *Decimal) (*Decimal, bool) {
	xx := NewDecimal(d)
	yy := NewDecimal(y)

	if yy.mantissa > xx.mantissa {
		_, overflow := xx.RescaleOverflow(yy.mantissa)
		if overflow {
			return nil, true
		}
	} else if yy.mantissa < xx.mantissa {
		_, overflow := yy.RescaleOverflow(xx.mantissa)
		if overflow {
			return nil, true
		}
	}

	if xx.Lt(yy) {
		return nil, true
	}

	_, overflow := d.value.SubOverflow(xx.value, yy.value)
	if overflow {
		return nil, true
	}
	d.mantissa = xx.mantissa

	return d, false
}

// Mul d = d * y and return d
func (d *Decimal) Mul(y *Decimal) *Decimal {
	xx := NewDecimal(d)
	yy := NewDecimal(y)

	if yy.mantissa > xx.mantissa {
		xx.Rescale(yy.mantissa)
	} else if yy.mantissa < xx.mantissa {
		yy.Rescale(xx.mantissa)
	}

	d.value.Mul(xx.value, yy.value)
	d.mantissa = xx.mantissa + yy.mantissa

	d.Optimize()

	return d
}

// MulOverflow d = d * y and return d
func (d *Decimal) MulOverflow(y *Decimal) (*Decimal, bool) {
	xx := NewDecimal(d)
	yy := NewDecimal(y)

	if yy.mantissa > xx.mantissa {
		_, overflow := xx.RescaleOverflow(yy.mantissa)
		if overflow {
			return nil, true
		}
	} else if yy.mantissa < xx.mantissa {
		_, overflow := yy.RescaleOverflow(xx.mantissa)
		if overflow {
			return nil, true
		}
	}

	_, overflow := d.value.MulOverflow(xx.value, yy.value)
	if overflow {
		return nil, true
	}

	d.mantissa = xx.mantissa + yy.mantissa

	d.Optimize()

	return d, false
}

const defaultDivScale = 20

// Div d = d / y and return d
func (d *Decimal) Div(y *Decimal) *Decimal {
	if y.Eq(Zero) {
		return NewDecimalZero()
	}

	xx := NewDecimal(d)
	yy := NewDecimal(y)

	var scalerest uint8
	e := int64(xx.mantissa) - int64(yy.mantissa) - int64(defaultDivScale)
	// todo: check overflow uint8

	if e < 0 {
		xx.value.Mul(xx.value, ExpScale(int16(-e)))
		scalerest = defaultDivScale
	} else {
		yy.value.Mul(yy.value, ExpScale(int16(e)))
		scalerest = xx.mantissa
	}

	d.value.Div(xx.value, yy.value)
	d.mantissa = scalerest

	d.Optimize()

	return d
}

// DivOverflow d = d / y and return d
func (d *Decimal) DivOverflow(y *Decimal) (*Decimal, bool) {
	if y.Eq(Zero) {
		return NewDecimalZero(), false
	}

	xx := NewDecimal(d)
	yy := NewDecimal(y)

	var scalerest uint8
	e := int64(xx.mantissa) - int64(yy.mantissa) - int64(defaultDivScale)
	if e > MaxUint8 {
		return nil, true
	}

	if e < 0 {
		_, overflow := xx.value.MulOverflow(xx.value, ExpScale(int16(-e)))
		if overflow {
			return nil, true
		}
		scalerest = defaultDivScale
	} else {
		_, overflow := yy.value.MulOverflow(yy.value, ExpScale(int16(e)))
		if overflow {
			return nil, true
		}
		scalerest = xx.mantissa
	}

	d.value.Div(xx.value, yy.value)
	d.mantissa = scalerest

	d.Optimize()

	return d, false
}

func (d *Decimal) SetFromBig(value *big.Int, mantissa uint8) (v *Decimal, overflow bool) {
	overflow = d.value.SetFromBig(value)
	d.SetMantissa(mantissa)
	return d, overflow
}

func (d *Decimal) SetValue(value *uint256.Int) *Decimal {
	d.value = value
	return d
}

func (d *Decimal) GetValue() *uint256.Int {
	return d.value
}

func (d *Decimal) SetMantissa(mantissa uint8) *Decimal {
	d.mantissa = mantissa
	return d
}

func (d *Decimal) GetMantissa() uint8 {
	return d.mantissa
}

func (d *Decimal) FromString(value string) (ok bool) {
	if d == nil {
		*d = *NewDecimalZero()
	}

	d.value = uint256.NewInt(0)
	if value == "" || value == "0" {
		return true
	}

	var mantissa uint8 = 0
	var valBig = new(big.Int)
	var parts = strings.Split(value, ".")

	if len(parts) > 2 {
		return false
	} else if len(parts) == 1 {
		valBig, ok = valBig.SetString(value, 10)
	} else {
		if len(parts[1]) > MaxUint8 {
			return false
		}

		// drop suffix zeros
		zerosStart := len(parts[1]) - 1
		for zerosStart >= 0 && parts[1][zerosStart] == '0' {
			zerosStart--
		}
		parts[1] = parts[1][:zerosStart+1]

		valBig, ok = valBig.SetString(strings.Join(parts, ""), 10)
		mantissa = uint8(len(parts[1]))
	}

	if !ok {
		return false
	}

	if overflow := d.value.SetFromBig(valBig); overflow {
		return false
	}
	d.mantissa = mantissa

	return true
}

func (d *Decimal) RescaleOverflow(mantissa uint8) (value *Decimal, overflow bool) {
	if d == nil {
		return nil, false
	}

	if mantissa == d.mantissa {
		return d, false
	}

	if mantissa > d.mantissa {
		_, isOverflow := d.value.MulOverflow(d.value, ExpScale(int16(mantissa-d.mantissa)))
		if isOverflow {
			return nil, isOverflow
		}

		d.mantissa = mantissa
		return d, false
	}

	if mantissa < d.mantissa {
		d.value.Div(d.value, ExpScale(int16(d.mantissa-mantissa)))
		d.mantissa = mantissa
		return d, false
	}

	return d, false
}

func (d *Decimal) Rescale(mantissa uint8) *Decimal {
	if d == nil {
		return nil
	}

	if mantissa == d.mantissa {
		return d
	}

	if mantissa > d.mantissa {
		// todo: check the << scaling method
		d.value.Mul(d.value, ExpScale(int16(mantissa-d.mantissa)))
		d.mantissa = mantissa
		return d
	}

	if mantissa < d.mantissa {
		d.value.Div(d.value, ExpScale(int16(d.mantissa-mantissa)))
		d.mantissa = mantissa
		return d
	}

	return d
}

func (d *Decimal) ToBig() *big.Int {
	if d.value == nil {
		d.value = uint256.NewInt(0)
	}

	return d.value.ToBig()
}

func (d *Decimal) Optimize() {
	rawStringValue := d.ToBig().String()
	if rawStringValue == "0" {
		return
	}

	var decimalsForOptimize uint8
	for i := len(rawStringValue) - 1; i >= 0; i-- {
		if rawStringValue[i] != '0' {
			break
		}
		decimalsForOptimize++
	}

	if decimalsForOptimize > 0 {
		if decimalsForOptimize <= d.mantissa {
			d.Rescale(d.mantissa - decimalsForOptimize)
		} else {
			d.Rescale(0)
		}
	}
}

func (d *Decimal) String() string {
	if d == nil || d.value == nil {
		return "0"
	}

	return FormatUint256(d.value, int(d.mantissa))
}

func (d *Decimal) IsZero() bool {
	return d.value.IsZero()
}

func NewDecimalFromUint256(value *uint256.Int, mantissa uint8) *Decimal {
	valueCopy := uint256.NewInt(0)
	copy(valueCopy[:], value[:4])

	return &Decimal{
		value:    valueCopy,
		mantissa: mantissa,
	}
}

func NewDecimalFromBig(value *big.Int, mantissa uint8) *Decimal {
	if value == nil {
		value = new(big.Int)
	}
	valueUint256, overflow := uint256.FromBig(value)
	if overflow {
		return NewDecimalZero()
	}

	return &Decimal{
		value:    valueUint256,
		mantissa: mantissa,
	}
}

func NewDecimalFromUint64(value uint64) *Decimal {
	d, _ := NewDecimalZero().SetFromBig(big.NewInt(int64(value)), 0)
	return d
}

func NewDecimalFromFloat64(value float64) *Decimal {
	val, _ := NewDecimalFromString(strconv.FormatFloat(value, 'f', -1, 64))
	return val
}

func NewDecimalZero() *Decimal {
	return &Decimal{
		value:    uint256.NewInt(0),
		mantissa: 0,
	}
}

func NewDecimalOne() *Decimal {
	return &Decimal{
		value:    uint256.NewInt(0).SetUint64(1),
		mantissa: 0,
	}
}

func NewDecimalFromString(val string) (*Decimal, bool) {
	d := NewDecimalZero()
	if !d.FromString(val) {
		return nil, false
	}

	return d, true
}

func NewDecimalFromStringOrDefault(val string, defaultValue *Decimal) *Decimal {
	d := NewDecimalZero()
	if !d.FromString(val) {
		return defaultValue
	}

	return d
}

func NewDecimal(decimal *Decimal) *Decimal {
	if decimal == nil {
		decimal = NewDecimalZero()
	}
	return NewDecimalFromUint256(decimal.value, decimal.mantissa)
}
