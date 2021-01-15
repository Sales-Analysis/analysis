package abc

import (
	"testing"
)

var data = []float64{1.0, 2.0, 3.0}

func TestFindPositive(t *testing.T) {
	for _, v := range data {
		_, ok := find(data, v)
		if ok == false {
			t.Errorf("value %f exist in slice: %f", v, data)
		}
	}
}

var dataError = []float64{0.1, 1.1, 1.2}

func TestFindErrorCase(t *testing.T) {
	for _, v := range data {
		_, ok := find(dataError, v)
		if ok != false {
			t.Errorf("value %f not exist in slice: %f", v, data)
		}
	}
}

func TestSum(t *testing.T) {
	total := sum(data)
	if total != 6.0 {
		t.Errorf("sum slice %f not equal %f", data, 6.0)
	}
}
