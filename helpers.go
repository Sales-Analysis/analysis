package abc

import "strconv"

func preData(records map[int]interface{}) ([]string, []float64){
	var measures []string
	var dimensions []float64
	for i, record := range records {
		if i == 0 {
			continue
		}
		r := record.(map[int]interface{})

		d, _ := strconv.ParseFloat(r[2].(string), 64)

		measures = append(measures, r[1].(string))
		dimensions = append(dimensions, d)
	}
	return measures, dimensions
}