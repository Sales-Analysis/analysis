package analysis

import (
	"testing"
)

func TestReadCsv(t *testing.T) {
	records, err := readCsv("./data/abc_test.csv")
	if err != nil {
		t.Error(err)
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