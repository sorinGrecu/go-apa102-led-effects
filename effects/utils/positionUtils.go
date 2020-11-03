package utils

import "github.com/adam-lavrik/go-imath/ix"

// Validates that the step given for various settings is between 0 and 100
func ValidateStep(step int) int {
	return ix.Min(ix.Max(0, step), 100)
}

func GetValueFromPercentageAndValue(percentage int, value float64) float64 {
	return float64(percentage) * (value / 100)
}
