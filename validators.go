package abc

func validate(measure []string, dimension []float64) ([]string, []float64, error) {
	if err := checkSize(measure, dimension); err != nil {
		return measure, dimension, err
	}
	if err := negativeValue(dimension); err != nil {
		return measure, dimension, err
	}
	measure, dimension = uniqueMeasures(measure, dimension)
	return measure, dimension, nil
}

// checkSize check length measure/dimension
func checkSize(measure []string, dimension []float64) error {
	if len(measure) == 0 {
		return sizeErr(measureNotEqualZero)
	} else if len(dimension) == 0 {
		return sizeErr(dimensionNotEqualZero)
	} else if len(measure) != len(dimension) {
		return sizeErr(measureNotEqualDimension)
	}
	return nil
}

// negativeValue return error if there is one negative value
func negativeValue(dimension []float64) error {

	for _, v := range dimension {
		if v < 0.0 {
			return negativeValueErr(negativeValueDimension)
		}
	}

	return nil
}

// uniqueMeasures return new array without duplucate
func uniqueMeasures(measure []string, dimension []float64) ([]string, []float64) {
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
