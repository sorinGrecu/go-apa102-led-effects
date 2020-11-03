package utils

import (
	"github.com/adam-lavrik/go-imath/ix"
	"image/color"
	"math/rand"
	"time"
)

type ColorContinuum struct {
	CurrentIndex int
	Colors       []color.RGBA
}

func (continuum *ColorContinuum) IncrementAndGetColor() color.RGBA {
	continuum.CurrentIndex = (continuum.CurrentIndex + 1) % len(continuum.Colors)
	return continuum.GetCurrentColor()
}

func (continuum *ColorContinuum) GetCurrentColor() color.RGBA {
	return continuum.Colors[continuum.CurrentIndex]
}

func (continuum *ColorContinuum) peekNextColor() color.RGBA {
	return continuum.Colors[continuum.CurrentIndex+1%len(continuum.Colors)]
}

func (continuum *ColorContinuum) PeekPreviousColor() color.RGBA {
	if continuum.CurrentIndex == 0 {
		return continuum.Colors[len(continuum.Colors)-1]
	}
	return continuum.Colors[continuum.CurrentIndex-1]
}

// Generates two random indexes to be picked from the length of the Colors array.
// The second index will be picked from
func (continuum *ColorContinuum) getTwoRandom() (color.RGBA, color.RGBA) {
	rand.Seed(time.Now().UnixNano())
	firstIndex := rand.Intn(len(continuum.Colors))
	secondIndex := ix.Abs(firstIndex - rand.Intn(len(continuum.Colors)-1))
	return continuum.Colors[firstIndex], continuum.Colors[secondIndex]
}
