package abc

func validate(measures []string, dimensions []float64) error{
	if err := checkSize(measures, dimensions); err != nil {
		return err
	}
	return nil
}

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