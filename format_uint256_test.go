package decimal_test

import (
	"math/big"
	"testing"

	"github.com/holiman/uint256"
	"github.com/stasundr/decimal"
)

func BenchmarkFormatUint256(b *testing.B) {
	v := uint256.NewInt(0)
	v.SetFromBig(big.NewInt(9))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		decimal.FormatUint256(v, 5)
	}
}

func TestFormatUint256(t *testing.T) {
	v := uint256.NewInt(0)
	v.SetFromBig(big.NewInt(1))

	if expected, actual := "1", decimal.FormatUint256(v, 0); actual != expected {
		t.Fatalf("expected: %s actual: %s", expected, actual)
	}
	if expected, actual := "0.1", decimal.FormatUint256(v, 1); actual != expected {
		t.Fatalf("expected: %s actual: %s", expected, actual)
	}
	if expected, actual := "0.01", decimal.FormatUint256(v, 2); actual != expected {
		t.Fatalf("expected: %s actual: %s", expected, actual)
	}
	if expected, actual := "0.001", decimal.FormatUint256(v, 3); actual != expected {
		t.Fatalf("expected: %s actual: %s", expected, actual)
	}

	v.SetFromBig(big.NewInt(11))

	if expected, actual := "11", decimal.FormatUint256(v, 0); actual != expected {
		t.Fatalf("expected: %s actual: %s", expected, actual)
	}
	if expected, actual := "1.1", decimal.FormatUint256(v, 1); actual != expected {
		t.Fatalf("expected: %s actual: %s", expected, actual)
	}
	if expected, actual := "0.11", decimal.FormatUint256(v, 2); actual != expected {
		t.Fatalf("expected: %s actual: %s", expected, actual)
	}
	if expected, actual := "0.011", decimal.FormatUint256(v, 3); actual != expected {
		t.Fatalf("expected: %s actual: %s", expected, actual)
	}

	v.SetFromBig(big.NewInt(111))

	if expected, actual := "111", decimal.FormatUint256(v, 0); actual != expected {
		t.Fatalf("expected: %s actual: %s", expected, actual)
	}
	if expected, actual := "11.1", decimal.FormatUint256(v, 1); actual != expected {
		t.Fatalf("expected: %s actual: %s", expected, actual)
	}
	if expected, actual := "1.11", decimal.FormatUint256(v, 2); actual != expected {
		t.Fatalf("expected: %s actual: %s", expected, actual)
	}
	if expected, actual := "0.111", decimal.FormatUint256(v, 3); actual != expected {
		t.Fatalf("expected: %s actual: %s", expected, actual)
	}

	v.SetFromBig(big.NewInt(0))

	if expected, actual := "0", decimal.FormatUint256(v, 0); actual != expected {
		t.Fatalf("expected: %s actual: %s", expected, actual)
	}
	if expected, actual := "0", decimal.FormatUint256(v, 1); actual != expected {
		t.Fatalf("expected: %s actual: %s", expected, actual)
	}
	if expected, actual := "0", decimal.FormatUint256(v, 2); actual != expected {
		t.Fatalf("expected: %s actual: %s", expected, actual)
	}
	if expected, actual := "0", decimal.FormatUint256(v, 3); actual != expected {
		t.Fatalf("expected: %s actual: %s", expected, actual)
	}
}
