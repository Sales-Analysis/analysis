package abc


type ABCSM struct {
	ABCSales *ABC
	ABCMargin *ABC
}

type ABC struct {
	Measures        []string
	Dimensions      []float64
	Deposit         []float64
	CumulativeShare []float64
	Group           []string
	Duplicates      []duplicate
}

type duplicate struct {
	PluID     int64
	Measure   string
	Dimension float64
}
