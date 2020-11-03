package utils

import (
	"image/color"
	"math"
)

func GetGradient(start color.RGBA, end color.RGBA, offset float64) color.RGBA {
	offset = math.Min(offset, 1)
	red := extractGradientOffset(start.R, end.R, offset)
	green := extractGradientOffset(start.G, end.G, offset)
	blue := extractGradientOffset(start.B, end.B, offset)
	return color.RGBA{R: red, G: green, B: blue, A: 255}
}

func extractGradientOffset(start uint8, end uint8, offset float64) uint8 {
	return uint8((offset * float64(end)) + (1-offset)*float64(start))
}

// Min and Max Duration is basically how much we want it to take to get from 0 to 1 in increments.
// For example, if the speed percentage is equal to 1, then it will be close to the min duration
// Given a speed percentage, refresh rate, the minimum and maximum duration of the animation,
// this will return by how much the speed should be incremented at each update cycle
func GetSpeedIncrement(speedPercentage float64, refreshRate, minDuration, maxDuration int) float64 {
	//There will be a range of 100 steps. Computing how many ms one step means based on
	//the range between max and min duration
	oneStepInMs := math.Abs(float64(maxDuration-minDuration)) / 100.00
	//The higher the speed is set, the smaller this speed divider will be
	speedModifierInMs := float64(maxDuration) - (speedPercentage * oneStepInMs)
	speedModifierInS := speedModifierInMs / 1000
	speed := 1 / (float64(refreshRate) * speedModifierInS)
	return speed
}

// Given a length percentage, and the extent of each color, this will return the size of the pause between colors
func GetLengthIncrement(lengthPercentage, colorLength int) int {
	return int(math.Round(GetValueFromPercentageAndValue(ValidateStep(lengthPercentage), float64(colorLength))))
}
