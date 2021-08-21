// Package abc
/*
	ABC sales analysis identifies the share of each of the products (or categories) in the total turnover,
	ranks the assortment positions (or categories) according to the degree of significance of their contribution
	to the total turnover, assigning conditional classes: A, B, C
*/
package abc

import (
	"sort"
)

type ABC struct {
	Measures []string
	Dimensions []float64
	Deposit []float64
	CumulativeShare []float64
	Group []string
}


func abc(measures []string, dimensions []float64) (*ABC, error){

	if err := validate(measures, dimensions); err != nil {
		return &ABC{}, err
	}

	m, d := sortParameters(measures, dimensions)
	total := totalSum(d)
	deposit := shareOfSales(d, total)
	cs := cumulativeShare(deposit)
	g := group(cs)
	return &ABC{
		m,
		d,
		deposit,
		cs,
		g,
	}, nil
}

// sortParameters Sorts the list in descending order of the sales value.
func sortParameters(measures []string, dimensions []float64) ([]string, []float64){
	d := make([]float64, len(dimensions))
	copy(d, dimensions)

	sort.Sort(sort.Reverse(sort.Float64Slice(d)))

	var m []string
	for _, v := range d {
		index, _ := findFloat64(dimensions, v)
		m = append(m, measures[index])
	}
	return m, d
}

// totalSum Determines the total amount of sales.
func totalSum(dimensions []float64) float64{
	return sum(dimensions)
}

// shareOfSales Determines the share of sales for each analyzed position
// (sales for each are divided by the total amount, the value in % is output).
//
// The values are taken out in a separate column "share".
func shareOfSales(dimensions []float64, total float64) []float64{
	var deposit []float64
	for _, v := range dimensions {
		d := v * 100.0 / total
		deposit = append(deposit, d)
	}
	return deposit
}

// cumulativeShare Considers the accumulated share as a cumulative total.
//
// The values are taken out in a separate next column "accumulated share".
func cumulativeShare(deposit []float64)  []float64{
	return cumulativeSum(deposit)
}

// group Assign categories A/B/C
func group(cumulativeShare []float64) []string{
	var g []string

	var position string
	for _, v := range cumulativeShare {
		if v <= 80.0 {
			position = "A"
		} else if v >= 80.0 && v <= 95.0 {
			position = "B"
		} else {
			position = "C"
		}
		g = append(g, position)
	}
	return g
}