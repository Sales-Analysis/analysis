package abc

import (
	"fmt"
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

func TestCalcABC1(t *testing.T) {
	for i, v := range measures {
		abc, err := GetABC(v, dimensions[i])
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(abc)
		}
	}
}
