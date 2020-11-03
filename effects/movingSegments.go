package effects

import (
	"go-led-strip/effects/utils"
	"go-led-strip/services/config"
	"golang.org/x/image/colornames"
	"image/color"
)

type MovingSegments struct {
	EffectLedSpan, startingPosition, colorPauseSize int
	running                                         bool
	colorTransitionDecimal, speedIncrement          float64
	ColorContinuum                                  utils.ColorContinuum
	config                                          config.EffectConfig
}

func NewMovingSegmentsFromConfig(config config.EffectConfig) *MovingSegments {
	continuum := config.GetContinuum()
	return &MovingSegments{
		speedIncrement: utils.GetSpeedIncrement(config.Speed, config.RefreshRate, config.MinDuration, config.MaxDuration),
		colorPauseSize: utils.GetLengthIncrement(config.Pause, config.Span/len(continuum.Colors)),
		ColorContinuum: continuum,
		EffectLedSpan:  config.Span,
		config:         config,
	}
}

func (effect *MovingSegments) computeFrame() []color.RGBA {
	frame := make([]color.RGBA, effect.EffectLedSpan)
	if effect.colorTransitionDecimal > 1 {
		effect.colorTransitionDecimal = 0
		effect.moveColorsToLeft()
	}
	ledCount := effect.EffectLedSpan
	colorLength := ledCount / len(effect.ColorContinuum.Colors)

	currentPoint := effect.startingPosition
	blackCount := 0

	for i := 1; i <= ledCount; i++ {
		color := effect.ColorContinuum.GetCurrentColor()
		transitionColor := colornames.Black
		if effect.colorPauseSize == 0 {
			transitionColor = effect.ColorContinuum.GetCurrentColor()
		}
		switch {
		case blackCount == 0:
			color = utils.GetGradient(effect.ColorContinuum.PeekPreviousColor(), transitionColor, effect.colorTransitionDecimal)
		case blackCount < effect.colorPauseSize:
			color = transitionColor
		case blackCount == effect.colorPauseSize:
			color = utils.GetGradient(transitionColor, effect.ColorContinuum.GetCurrentColor(), effect.colorTransitionDecimal)
		}
		frame[currentPoint] = color
		blackCount++
		if i%colorLength == 0 {
			effect.ColorContinuum.IncrementAndGetColor()
			blackCount = 0
		}

		currentPoint++
		currentPoint = currentPoint % ledCount
	}
	effect.colorTransitionDecimal += effect.speedIncrement
	if effect.config.Reverse {
		for i, j := 0, len(frame)-1; i < j; i, j = i+1, j-1 {
			frame[i], frame[j] = frame[j], frame[i]
		}
	}
	return frame
}

func (effect *MovingSegments) GetFrame() []color.RGBA {
	if !effect.running {
		return make([]color.RGBA, effect.EffectLedSpan)
	}
	return effect.computeFrame()
}

func (effect *MovingSegments) moveColorsToLeft() {
	effect.startingPosition--
	if effect.startingPosition < 0 {
		effect.startingPosition = effect.EffectLedSpan - 1
	}
}

func (effect *MovingSegments) Start() {
	effect.running = true
}

func (effect *MovingSegments) Stop() {
	effect.running = false
}
