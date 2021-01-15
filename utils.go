package abc

// sum return total sum slice
func sum(arr []float64) float64 {
	total := 0.0
	for _, v := range arr {
		total += v
	}
	return total

}

// Find takes a slice and looks for an element in it. If found it will
// return it's key, otherwise it will return -1 and a bool of false.
func find(arr []float64, value float64) (int, bool) {
	for i, v := range arr {
		if v == value {
			return i, true
		}
	}
	return -1, false
}
