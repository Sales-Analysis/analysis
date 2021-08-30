package analysis

import "github.com/Sales-Analysis/analysis/abc"

func ABC(pluID []int64, measures []string, dimensions []float64) (*abc.ABC, error) {
	return abc.Report(pluID, measures, dimensions)
}
