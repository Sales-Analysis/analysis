package abc

import (
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

func check(t *testing.T, a float64, b float64) {
	if a != b {
		t.Errorf("got %f want %f", a, b)
	}
}

func TestSumDimension(t *testing.T) {
	abc := &ABC{
		Measure:   goods,
		Dimension: cost,
	}
	want := 2100.0
	check(t, abc.sumDimension(), want)
}

func TestSumPercent(t *testing.T) {
	abc := &ABC{
		Measure:   goods,
		Dimension: cost,
	}
	abc.deposit()
	want := 100.0
	check(t, abc.sumPercent(), want)
}

func TestCalcABC(t *testing.T) {
	abc := &ABC{
		Measure:   goods,
		Dimension: cost,
	}
	abc.sort()
	abc.deposit()

	d := abc.Deposit

	totalDeposit := 0.0
	for _, v := range d {
		totalDeposit += v
	}
	want := 100.0
	check(t, totalDeposit, want)

	abc.accumDeposit()
	v := abc.AccumulativeDeposit[len(abc.AccumulativeDeposit)-1]
	check(t, v, want)

	abc.addGroup()

}
