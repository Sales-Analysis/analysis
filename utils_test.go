package abc

import (
	"testing"
)

var dataFloat = []float64{1.0, 2.0, 3.0}

func TestFindFloatPositive(t *testing.T) {
	for _, v := range dataFloat {
		_, ok := findFloat64(dataFloat, v)
		if ok == false {
			t.Errorf("value %f exist in slice: %f", v, dataFloat)
		}
	}
}

var dataFloatError = []float64{0.1, 1.1, 1.2}

func TestFindFloatErrorCase(t *testing.T) {
	for _, v := range dataFloat {
		_, ok := findFloat64(dataFloatError, v)
		if ok != false {
			t.Errorf("value %f not exist in slice: %f", v, dataFloat)
		}
	}
}

var dataInt = []int{1, 2, 3}

func TestFindIntPositive(t *testing.T) {
	for _, v := range dataInt {
		_, ok := findInt(dataInt, v)
		if ok == false {
			t.Errorf("value %d exist in slice: %d", v, dataInt)
		}
	}
}

var dataIntError = []int{4, 5, 6}

func TestFindIntErrorCase(t *testing.T) {
	for _, v := range dataInt {
		_, ok := findInt(dataIntError, v)
		if ok != false {
			t.Errorf("value %d not exist in slice: %d", v, dataInt)
		}
	}
}

func TestSum(t *testing.T) {
	total := sum(dataFloat)
	if total != 6.0 {
		t.Errorf("sum slice %f not equal %f", dataFloat, 6.0)
	}
}

func TestReadCsv(t *testing.T) {
	records, err := readCsv("./data/abc_test.csv")
	if err != nil {
		t.Errorf("Error, while file reading")
	}
	if len(records) != 6 {
		t.Errorf("Data loss")
	}
}

func TestReadExcel(t *testing.T)  {
	records, err := readExcel("./data/abc_test.xlsx")
	if err != nil {
		t.Errorf("Error, while file reading")
	}
	if len(records) != 6 {
		t.Errorf("Data loss")
	}
}