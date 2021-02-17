package abc

// sum return total sum slice
func sum(arr []float64) float64 {
	total := 0.0
	for _, v := range arr {
		total += v
	}
	return total

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

func findInt(arr []int, value int) (int, bool) {
	return find(len(arr), func(i int) bool { return arr[i] == value })
}

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
