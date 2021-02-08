package abc

import (
	"testing"
)

var measures = [][]string{
	[]string{"Товар11", "Товар12", "Товар13", "Товар14", "Товар15", "Товар16"},
	[]string{"Товар21", "Товар21", "Товар23", "Товар24", "Товар25", "Товар26"},
	[]string{"Товар31", "Товар31", "Товар31", "Товар31", "Товар31", "Товар31"},
	[]string{"Товар41", "Товар42"},
	[]string{"Товар51"},
	[]string{},
	[]string{"Товар7"},
	[]string{"Товар81", "Товар82"},
}

var dimensions = [][]float64{
	[]float64{1, 2, 3, 4, 5, 6},
	[]float64{1, 2, 3, 4, 5, 6},
	[]float64{1, 2, 3, 4, 5, 6},
	[]float64{1, 2},
	[]float64{1},
	[]float64{1, 2, 3, 4, 5, 6},
	[]float64{},
	[]float64{1, -2},
}

// output dataset
var groups = [][]string{
	{"A", "A", "A", "B", "C", "C"},
	{"A", "A", "A", "B", "C"},
	{"C"},
	{"A", "C"},
	{"C"},
}
var sumDeposit = 100.0
var sumDimensions = []float64{21.0, 21.0, 21.0, 3.0, 1.0}
var errorSet = []string{
	"measure size is 0",
	"dimension size is 0",
	"in the dimensions is a negative value",
}

func findStr(array []string, str string) bool {
	for _, v := range array {
		if v == str {
			return true
		}
	}
	return false
}

func TestCalcABC1(t *testing.T) {
	for i, v := range measures {
		abc, err := GetABC(v, dimensions[i])
		if err != nil {
			if findStr(errorSet, err.Error()) == false {
				t.Errorf("error code is not exist: %s", err.Error())
			}
		} else {
			if abc.SumDimensions != sumDimensions[i] {
				t.Errorf("result array %v not equal test dataset %v", abc.SumDimensions, sumDimensions[i])
			}
			if int(abc.SumDeposit) != int(sumDeposit) {
				t.Errorf("get sum deposit %f want sum deposit %f", abc.SumDeposit, sumDeposit)
			}
			if len(abc.Group) != len(groups[i]) {
				t.Errorf("result array %v not equal test dataset %v", abc.Group, groups[i])
			}
			for j, v := range abc.Group {
				if v != groups[i][j] {
					t.Errorf("result array %v not equal test dataset %v", abc.Group, groups[i])
				}
			}
		}
	}
}
