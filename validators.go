package abc

func validate(measures []string, dimensions []float64) error{
	if err := checkSize(measures, dimensions); err != nil {
		return err
	}
	if err := negativeValue(dimensions); err != nil {
		return err
	}
	return nil
}

// checkSize check length measure/dimension
func checkSize(measures []string, dimensions []float64)  error{
	if len(measures) == 0 {
		return sizeErr(measuresNotEqualZero)
	} else if len(dimensions) == 0 {
		return sizeErr(dimensionsNotEqualZero)
	} else if len(measures) != len(dimensions) {
		return sizeErr(measuresNotEqualDimensions)
	}
	return nil
}

// negativeValue return error if there is one negative value
func negativeValue(dimensions []float64) error {
	for _, v := range dimensions {
		if v < 0.0 {
			return negativeValueErr(negativeValueDimensions)
		}
	}
	return nil
}