package analysis

import (
	"testing"
)

var measuresTestData = []string{
	"Товар 1", "Товар 2", "Товар 3", "Товар 4", "Товар 5",
}

var dimensionsTestData = []float64{
	100, 50, 35, 20, 5,
}

var depositTestData = []float64{
	47.61904761904762, 23.80952380952381, 16.666666666666668, 9.523809523809524, 2.380952380952381,
}

var cumulativeShareTestData = []float64{
	47.61904761904762, 71.42857142857143, 88.0952380952381, 97.61904761904762, 100,
}

var groupTestData = []string{
	"A", "A", "B", "C", "C",
}

func TestABCCsv(t *testing.T) {
	records, _ := readCsv("./data/abc_test.csv")
	pluID, measures, dimensions := preData(records)
	report, _ := ABC(pluID, measures, dimensions)
	// fmt.Println(report.Measures, measuresTestData)
	for i, v := range report.Measures {
		if v != measuresTestData[i] {
			t.Errorf("result array %v not equal test data %v", report.Measures, measuresTestData)
		}
		if report.Dimensions[i] != dimensionsTestData[i] {
			t.Errorf("result array %v not equal test data %v", report.Dimensions[i], dimensionsTestData[i])
		}
		if report.Deposit[i] != depositTestData[i] {
			t.Errorf("result array %v not equal test data %v", report.Deposit[i], depositTestData[i])
		}
		if report.CumulativeShare[i] != cumulativeShareTestData[i] {
			t.Errorf("result array %v not equal test data %v", report.CumulativeShare[i], cumulativeShareTestData[i])
		}
		if report.Group[i] != groupTestData[i] {
			t.Errorf("result array %v not equal test data %v", report.Group[i], groupTestData[i])
		}
	}
}

func TestABCExcel(t *testing.T) {
	records, _ := readExcel("./data/abc_test.xlsx")
	pluID, measures, dimensions := preData(records)
	report, _ := ABC(pluID, measures, dimensions)
	for i, v := range report.Measures {
		if v != measuresTestData[i] {
			t.Errorf("result array %v not equal test data %v", report.Measures, measuresTestData)
		}
		if report.Dimensions[i] != dimensionsTestData[i] {
			t.Errorf("result array %v not equal test data %v", report.Dimensions[i], dimensionsTestData[i])
		}
		if report.Deposit[i] != depositTestData[i] {
			t.Errorf("result array %v not equal test data %v", report.Deposit[i], depositTestData[i])
		}
		if report.CumulativeShare[i] != cumulativeShareTestData[i] {
			t.Errorf("result array %v not equal test data %v", report.CumulativeShare[i], cumulativeShareTestData[i])
		}
		if report.Group[i] != groupTestData[i] {
			t.Errorf("result array %v not equal test data %v", report.Group[i], groupTestData[i])
		}
	}
}

func TestABCDuplicates(t *testing.T) {
	records, _ := readCsv("./data/abc_simple.csv")
	pluID, measures, dimensions := preData(records)
	report, _ := ABC(pluID, measures, dimensions)
	if len(report.Duplicates) != 1 {
		t.Errorf("length result array %v not equal test data %v", len(report.Duplicates), 1)
	}
	if (report.Duplicates[0].PluID != 1) &&
		(report.Duplicates[0].Measure != "Товар 1") &&
		(report.Duplicates[0].Dimension != 100) {
		t.Errorf("The values in the structure differ from the test values")
	}
}

func TestACBDuplicateValues(t *testing.T) {
	var pluID = []int64{
		1, 2, 3, 4, 5,
	}
	var measures = []string{
		"Товар 1", "Товар 2", "Товар 3", "Товар 4", "Товар 5",
	}

	var dimensions = []float64{
		100, 200, 100, 35, 5,
	}

	ABC(pluID, measures, dimensions)

	// fmt.Println(report)
}
