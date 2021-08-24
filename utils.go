package abc

import (
	"encoding/csv"
	"github.com/xuri/excelize/v2"
	"os"
)

// sum return total sum slice
func sum(arr []float64) float64 {
	total := 0.0
	for _, v := range arr {
		total += v
	}
	return total
}

// cumulativeSum return cumulative sum slice
func cumulativeSum(arr []float64) []float64 {
	var a []float64
	accum := 0.0
	for _, v := range arr {
		accum += v
		a = append(a, accum)
	}
	return a
}

// find takes a slice and looks for an element in it. If found it will
// return it's key, otherwise it will return -1 and a bool of false.
func find(n int, f func(int) bool) (int, bool) {
	for i := 0; i < n; i++ {
		if f(i) {
			return i, true
		}
	}
	return -1, false
}

// findInt searches for value in slice type int.
// If value exist return index and true, otherwise -1 and false.
func findInt(arr []int, value int) (int, bool) {
	return find(len(arr), func(i int) bool { return arr[i] == value })
}

// findFloat searches for value in slice type float.
// If value exist return index and true, otherwise -1 and false.
func findFloat64(arr []float64, value float64) (int, bool) {
	return find(len(arr), func(i int) bool { return arr[i] == value })
}

// mapToSlice return two slice keys/values
func mapToSlice(data map[string]float64) ([]string, []float64) {
	var keys []string
	var values []float64

	for key, value := range data {
		keys = append(keys, key)
		values = append(values, value)
	}
	return keys, values
}


// read csv file.
// return map[int]interface{}, or error
func readCsv(path string) (map[int]interface{}, error){
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records := make(map[int]interface{})
	row := 0

	for {
		record, e := reader.Read()
		if e != nil {
			break
		}

		r := make(map[int]interface{})
		for i, v := range record {
			r[i] = v
		}
		records[row] = r
		row += 1
	}
	return records, nil
}

// read excel file
// return map[int]interface{}, or error
func readExcel(path string) (map[int]interface{}, error){
	file, err := excelize.OpenFile(path)
	if err != nil {
		return nil, err
	}
	// TODO: зависит от языка
	// rows, err := file.GetRows("шаблон")
	rows, err := file.GetRows("шаблон")

	if err != nil {
		return nil, err
	}
	records := make(map[int]interface{})
	for i, row := range rows {
		r := make(map[int]interface{})
		for j, v := range row {
			r[j] = v
		}
		records[i] = r
	}
	return records, nil
}