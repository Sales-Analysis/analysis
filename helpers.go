package analysis

import (
	"strconv"
)


func preData(records map[int]interface{}) ([]int64, []string, []float64, []float64){
	var pluID []int64
	var measures []string
	var dimensionsSales []float64
	var dimensionsMargin []float64

	for i, record := range records {
		if i == 0 {
			continue
		}
		r := record.(map[int]interface{})

		ds, _ := strconv.ParseFloat(r[2].(string), 64)

		var dm float64
		if r[3] != nil {
			dm, _ = strconv.ParseFloat(r[3].(string), 64)
		}

		p, _ := strconv.ParseInt(r[0].(string), 10, 64)

		pluID = append(pluID, p)
		measures = append(measures, r[1].(string))
		dimensionsSales = append(dimensionsSales, ds)
		dimensionsMargin = append(dimensionsMargin, dm)
	}
	return pluID, measures, dimensionsSales, dimensionsMargin
}