package abc

import (
	"fmt"
	"testing"
)

/*
 	N	Товар		Стоимость
	1	Товар1		100
	2	Товар2		200
	3	Товар3		300
	4	Товар4		400
	5	Товар5		500
	6	Товар6		600

*/

var goods = []string{
	"Товар1",
	"Товар2",
	"Товар3",
	"Товар4",
	"Товар5",
	"Товар6",
}

var cost = []float64{
	100,
	200,
	300,
	400,
	500,
	600,
}

func initABC() *ABC {
	return &ABC{
		Measure:   goods,
		Dimension: cost,
	}
}

func status(t *testing.T, a float64, b float64) {
	if a != b {
		t.Errorf("got %f want %f", a, b)
	}
}

func TestSumTotal(t *testing.T) {
	checkSum := func(t *testing.T, got float64, want float64) {
		t.Helper()
		status(t, got, want)
	}

	t.Run("SumDimension", func(t *testing.T) {
		abc := initABC()
		checkSum(t, abc.sumDimension(), 2100.0)
	})

	t.Run("TestCalcABC", func(t *testing.T) {
		abc := initABC()
		abc.deposit()
		checkSum(t, abc.sumPercent(), 100.0)
	})
}

/*
var goods = []string{
	"Товар1",
	"Товар2",
	"Товар3",
	"Товар1",
	"Товар5",
	"Товар6",
}

var cost = []float64{
	100,
	200,
	300,
	400,
	500,
	600,
}
*/

func TestCalcABC(t *testing.T) {
	abc := GetABC(goods, cost)
	abc.sort()
	abc.deposit()

	d := abc.Deposit

	totalDeposit := 0.0
	for _, v := range d {
		totalDeposit += v
	}
	want := 100.0
	status(t, totalDeposit, want)

	abc.accumDeposit()
	v := abc.AccumulativeDeposit[len(abc.AccumulativeDeposit)-1]
	status(t, v, want)

	abc.addGroup()
}

var goods1 = []string{
	"Товар1",
	"Товар2",
	"Товар3",
	"Товар1",
	"Товар5",
	"Товар6",
}

var cost1 = []float64{
	100, // "Товар1"
	200,
	300,
	400, // "Товар1"
	500,
	600,
}

func TestCalcABC1(t *testing.T) {
	abc := GetABC(goods1, cost1)
	abc.sort()

	abc.deposit()
	d := abc.Deposit

	totalDeposit := 0.0
	for _, v := range d {
		totalDeposit += v
	}
	want := 100.0
	status(t, totalDeposit, want)
	abc.accumDeposit()

	v := abc.AccumulativeDeposit[len(abc.AccumulativeDeposit)-1]
	status(t, v, want)

	abc.addGroup()

}

func TestCheckSize(t *testing.T) {
	measure1 := []string{}
	dimension1 := []float64{1, 2, 3}

	if err := checkSize(measure1, dimension1); err != nil {
		if fmt.Sprintf("%s", err) != measureNotEqualZero {
			t.Errorf("'%s' error codes must be '%s'", err, measureNotEqualZero)
		}
	}

	measure2 := []string{"a", "b"}
	dimension2 := []float64{}
	if err := checkSize(measure2, dimension2); err != nil {
		if fmt.Sprintf("%s", err) != dimensionNotEqualZero {
			t.Errorf("'%s' error codes must be '%s'", err, dimensionNotEqualZero)
		}
	}

	measure3 := []string{"a", "b", "c"}
	dimension3 := []float64{0, 1}
	if err := checkSize(measure3, dimension3); err != nil {
		if fmt.Sprintf("%s", err) != measureNotEqualDimension {
			t.Errorf("'%s' error codes must be '%s'", err, measureNotEqualDimension)
		}
	}

}
