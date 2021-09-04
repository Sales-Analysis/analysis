package analysis

import "github.com/Sales-Analysis/analysis/abc"

func ABC(pluID []int64, measures []string, dimensions []float64) (*abc.ABC, error) {
	return abc.Abc(pluID, measures, dimensions)
}

func ABCTwoParameters(pluID []int64, measures []string, dimensionsSales []float64, dimensionsMargin []float64) {
	abc.AbcTwoParameters(pluID, measures, dimensionsSales, dimensionsMargin)
}