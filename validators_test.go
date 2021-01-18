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
	measure1 := []string{}
	dimension1 := []float64{1, 2, 3}

	if err := checkSize(measure1, dimension1); err != nil {
		checkStatusError(t, err, measureNotEqualZero)
	}

	measure2 := []string{"a", "b"}
	dimension2 := []float64{}
	if err := checkSize(measure2, dimension2); err != nil {
		checkStatusError(t, err, dimensionNotEqualZero)
	}

	measure3 := []string{"a", "b", "c"}
	dimension3 := []float64{0, 1}
	if err := checkSize(measure3, dimension3); err != nil {
		checkStatusError(t, err, measureNotEqualDimension)
	}

}

func TestNegativeValue(t *testing.T) {
	dimension := []float64{1, 2, 3, 4, -5, 6}
	if err := negativeValue(dimension); err != nil {
		checkStatusError(t, err, negativeValueDimension)
	}
}

func TestUniqueMeasure(t *testing.T) {
	measure := []string{"a", "b", "c", "a", "b", "b1"}
	dimension := []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0}

	measure, dimension = uniqueMeasures(measure, dimension)
	if len(measure) != len(dimension) {
		t.Errorf("size measure '%d' not equal size dimension '%d'", len(measure), len(dimension))
	}
	want := []float64{5.0, 7.0, 3.0, 6.0}
	for i, v := range want {
		if v != dimension[i] {
			t.Errorf("value dimension not equal test data: %f != %f", v, dimension[i])
		}
	}

	fmt.Println(measure, dimension)
}
