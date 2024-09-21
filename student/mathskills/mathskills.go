package mathskills

import "math"




func Average(data []float64)  float64 {
	sum := 0.0
	count := len(data)
	// Calculate the sum of the numbers
	for _, num := range data {
		sum += num
	}

	// Calculate the average
	average := float64(sum) / float64(count)
	roundedAverage := math.Round(average)
	return math.Round(roundedAverage)
}

func StandardDeviation(data []float64) float64 {
	return math.Sqrt(Variance(data))
}

func Variance(data []float64)float64 {
	var sqDiff float64
	mean := Average(data)
	for _, n := range data {
		diff := n - mean
		sqDiff += diff * diff
	}
	return sqDiff / float64(len(data))
}

func Range(data []float64) (int, int) {
	start := len(data) - 4  // Determine the starting index for the last four elements
	if start < 0 {          // If the starting index is negative, set it to 0
		start = 0
	}

	window := data[start:]   // Create a slice containing the last four (or fewer) elements
	mean := Average(window)  // Calculate the mean of the window slice
	stdDev := StandardDeviation(window)  // Calculate the standard deviation of the window slice

	// Calculate the lower and upper limits based on the mean and standard deviation
	lowerLimit := mean - (2 * stdDev)
	upperLimit := mean + (2 * stdDev)

	// Round the limits to the nearest integer and return them
	return int(math.Round(lowerLimit)), int(math.Round(upperLimit))
}

