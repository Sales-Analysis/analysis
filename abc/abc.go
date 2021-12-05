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
	Measures        []string
	Dimensions      []float64
	Deposit         []float64
	CumulativeShare []float64
	Group           []string
	Duplicates      []duplicate
}

// Report return struct analysis
func Report(pluID []int64, measures []string, dimensions []float64) (*ABC, error) {
	if err := validate(measures, dimensions); err != nil {
		return &ABC{}, err
	}
	m, d, duplicates := removeDuplicate(pluID, measures, dimensions)
	m, d = sortParameters(m, d)
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
		duplicates,
	}, nil
}

type duplicate struct {
	PluID     int64
	Measure   string
	Dimension float64
}

// removeDuplicate return new slices without duplicate
func removeDuplicate(pluID []int64, measures []string, dimensions []float64) ([]string, []float64, []duplicate) {
	var m []string
	var d []float64

	pluIDKey := make(map[int64]bool)
	measuresKey := make(map[string]bool)
	dimensionsKey := make(map[float64]bool)

	var duplicates []duplicate
	for i, item := range measures {
		_, vm := measuresKey[item]
		_, vd := dimensionsKey[dimensions[i]]
		_, vp := pluIDKey[pluID[i]]
		if !vm || !vd || !vp {
			measuresKey[item] = true
			dimensionsKey[dimensions[i]] = true
			pluIDKey[pluID[i]] = true
			m = append(m, item)
			d = append(d, dimensions[i])
		} else {
			duplicates = append(duplicates, duplicate{
				PluID:     pluID[i],
				Measure:   item,
				Dimension: dimensions[i],
			})
		}
	}
	return m, d, duplicates
}

// sortParameters Sorts the list in descending order of the sales value.
func sortParameters(measures []string, dimensions []float64) ([]string, []float64) {
	// TODO: problem if duplicate dimensions
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
func totalSum(dimensions []float64) float64 {
	return sum(dimensions)
}

// shareOfSales Determines the share of sales for each analyzed position
// (sales for each are divided by the total amount, the value in % is output).
//
// The values are taken out in a separate column "share".
func shareOfSales(dimensions []float64, total float64) []float64 {
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
func cumulativeShare(deposit []float64) []float64 {
	return cumulativeSum(deposit)
}

// group Assign categories A/B/C
func group(cumulativeShare []float64) []string {
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
