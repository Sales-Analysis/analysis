/*
In materials management, ABC analysis is an inventory categorization technique.
ABC analysis divides an inventory into three categories:
	"A items" with very tight control and accurate records,
	"B items" with less tightly controlled and good records,
	"C items" with the simplest controls possible and minimal records.
*/

package abc

import (
	"sort"
)

// ABC represents set of methods for the analysis
type ABC struct {
	Measures            []string
	Dimensions          []float64
	Deposit             []float64
	AccumulativeDeposit []float64
}

// GetABC return link to the struct ABC
func GetABC(measures []string, dimensions []float64) *ABC {
	measures, dimensions, err := validate(measures, dimensions)
	if err != nil {
		return &ABC{}
	}
	return &ABC{
		Measures:   measures,
		Dimensions: dimensions,
	}
}

// sum of all dimensions
func (abc *ABC) sumDimensions() float64 {
	return sum(abc.Dimensions)
}

// sum of all deposit
func (abc *ABC) sumPercent() float64 {
	return sum(abc.Deposit)
}

// sorting positions in descending order
func (abc *ABC) sort() {
	d := make([]float64, len(abc.Dimensions))

	copy(d, abc.Dimensions)
	sort.Sort(sort.Reverse(sort.Float64Slice(d)))

	var m []string
	for _, v := range d {
		i, _ := find(abc.Dimensions, v)
		m = append(m, abc.Measures[i])
	}

	abc.Measures = m
	abc.Dimensions = d
}

// determine deposit each position
func (abc *ABC) deposit() {
	var d []float64

	s := abc.sumDimensions()

	for _, v := range abc.Dimensions {
		percent := v * 100.0 / s
		d = append(d, percent)
	}
	abc.Deposit = d
}

// accumulative percentage
func (abc *ABC) accumDeposit() {
	var a []float64
	accum := 0.0
	for _, v := range abc.Deposit {
		accum += v
		a = append(a, accum)
	}
	abc.AccumulativeDeposit = a
}

// add group A or B or C for each position
func (abc *ABC) addGroup() {
	// TODO границы групп сделать изменяемыми
	var g []string

	for _, v := range abc.AccumulativeDeposit {
		if v <= 80.0 {
			g = append(g, "A")
		} else if v >= 80.0 && v <= 95.0 {
			g = append(g, "B")
		} else {
			g = append(g, "C")
		}
	}
}
