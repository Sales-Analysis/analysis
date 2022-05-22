package abc

import (
	"testing"
)

func TestSortParameters(t *testing.T) {
	var measures = []string{"Товар 3", "Товар 1", "Товар 2", "Товар 4"}
	var measuresExpected = []string{"Товар 3", "Товар 2", "Товар 4", "Товар 1"}
	var dimensions = []float64{100, 10, 50, 20}
	var dimensionsExpected = []float64{100, 50, 20, 10}
	m, d := sortParameters(measures, dimensions)
	for i, v := range m {
		if v != measuresExpected[i] {
			t.Errorf("result array %v not equal test data %v", m, measuresExpected)	
		}
		if d[i] != dimensionsExpected[i] {
			t.Errorf("result array %v not equal test data %v", m, measuresExpected)	
		}
	}
}

func TestSortDuplicateValues(t *testing.T) {
	var measures = []string{"Товар 1", "Товар 2", "Товар 3", "Товар 4", "Товар 5"}
	var measuresExpected = []string{"Товар 2", "Товар 1", "Товар 3", "Товар 4", "Товар 5"} 
	var dimensions = []float64{100, 200, 100, 35, 5}
	var dimensionsExpected = []float64{200, 100, 100, 35, 5}
	m, d := sortParameters(measures, dimensions)
	for i, v := range m {
		if v != measuresExpected[i] {
			t.Errorf("result array %v not equal test data %v", m, measuresExpected)	
		}
		if d[i] != dimensionsExpected[i] {
			t.Errorf("result array %v not equal test data %v", m, measuresExpected)	
		}
	}
}