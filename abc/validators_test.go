package abc

import (
	"fmt"
	"testing"
)

func checkStatusError(t *testing.T, err error, status string) {
	if fmt.Sprintf("%s", err) != status {
		t.Errorf("'%s' error codes must be '%s'", err, status)
	}
}

func TestCheckSize(t *testing.T) {
	var measures1 []string
	dimensions1 := []float64{1, 2, 3}

	if err := checkSize(measures1, dimensions1); err != nil {
		checkStatusError(t, err, measuresNotEqualZero)
	}

	measures2 := []string{"a", "b"}
	var dimensions2 []float64
	if err := checkSize(measures2, dimensions2); err != nil {
		checkStatusError(t, err, dimensionsNotEqualZero)
	}

	measures3 := []string{"a", "b", "c"}
	dimensions3 := []float64{0, 1}
	if err := checkSize(measures3, dimensions3); err != nil {
		checkStatusError(t, err, measuresNotEqualDimensions)
	}

}

func TestNegativeValue(t *testing.T) {
	dimensions := []float64{1, 2, 3, 4, -5, 6}
	if err := negativeValue(dimensions); err != nil {
		checkStatusError(t, err, negativeValueDimensions)
	}
}
