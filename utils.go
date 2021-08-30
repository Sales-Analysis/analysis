package analysis

import (
	"encoding/csv"
	"github.com/xuri/excelize/v2"
	"os"
)

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