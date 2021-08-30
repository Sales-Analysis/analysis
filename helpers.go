package analysis

import (
	"strconv"
)

func preData(records map[int]interface{}) ([]int64, []string, []float64){
	var pluID []int64
	var measures []string
	var dimensions []float64
	for i, record := range records {
		if i == 0 {
			continue
		}
		r := record.(map[int]interface{})

		d, _ := strconv.ParseFloat(r[2].(string), 64)
		p, _ := strconv.ParseInt(r[0].(string), 10, 64)

		pluID = append(pluID, p)
		measures = append(measures, r[1].(string))
		dimensions = append(dimensions, d)
	}
	return pluID, measures, dimensions
}