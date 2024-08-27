package mathskills

import (
	"math"
	"sort"
)

// Calculating The sum of the data set
func Sum(data []float64) float64 {
	var sum float64
	for i := 0; i < len(data); i++ {
		sum += data[i]
	}
	return sum
}

// Finding the average of the data set
func Average(data []float64) float64 {
	return Sum(data) / float64(len(data))
}

// Calculating the median
func Median(data []float64) float64 {
	// Sort the data
	sort.Float64s(data)

	n := len(data)
	var median float64

	if n%2 == 1 {
		median = data[n/2]
	} else {
		median = (data[(n/2)-1] + data[n/2]) / 2
	}

	return median
}

// Finding the variance
func Variance(data []float64) float64 {
	sumSquaredDifferences := 0.0
	for _, num := range data {
		difference := num - Average(data)
		sumSquaredDifferences += difference * difference
	}
	count := float64(len(data))
	return sumSquaredDifferences / count
}

func Std(data []float64) float64 {
	varianceValue := Variance(data)
	return math.Sqrt(varianceValue)
}
