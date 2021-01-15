package abc

func validate(measure []string, dimension []float64) ([]string, []float64) {
	return dropDuplicate(measure, dimension)
}

// dropDuplicate return new array without duplucate
func dropDuplicate(measure []string, dimension []float64) ([]string, []float64) {
	d := make(map[string]float64)

	for i, v := range measure {
		if _, ok := d[v]; ok {
			d[v] += dimension[i]
		} else {
			d[v] = dimension[i]
		}
	}
	return mapToSlice(d)
}

func mapToSlice(data map[string]float64) ([]string, []float64) {
	var keys []string
	var values []float64

	for key, value := range data {
		keys = append(keys, key)
		values = append(values, value)
	}
	return keys, values
}
