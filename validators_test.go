package abc

import (
	"fmt"
	"sort"
	"testing"
)

func checkStatusError(t *testing.T, err error, status string) {
	if fmt.Sprintf("%s", err) != status {
		t.Errorf("'%s' error codes must be '%s'", err, status)
	}
}

func TestCheckSize(t *testing.T) {
	measures1 := []string{}
	dimensions1 := []float64{1, 2, 3}

	if err := checkSize(measures1, dimensions1); err != nil {
		checkStatusError(t, err, measuresNotEqualZero)
	}

	measures2 := []string{"a", "b"}
	dimensions2 := []float64{}
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
	dimensions := []float64{1, 2, 3, 4, 5, 6}
	if err := negativeValue(dimensions); err != nil {
		checkStatusError(t, err, negativeValueDimensions)
	}
}

func TestUniqueMeasure(t *testing.T) {
	measures := []string{"a", "b", "c", "a", "b", "b1"}
	dimensions := []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0}

	measures, dimensions = uniqueMeasures(measures, dimensions)
	if len(measures) != len(dimensions) {
		t.Errorf("size measure '%d' not equal size dimension '%d'", len(measures), len(dimensions))
	}
	sort.Float64s(dimensions)
	want := []float64{3.0, 5.0, 6.0, 7.0}

	for i, v := range want {
		if v != dimensions[i] {
			t.Errorf("value dimension not equal test data: %f != %f", v, dimensions[i])
		}
	}
}
