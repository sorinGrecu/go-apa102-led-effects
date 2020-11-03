package effects

import (
	"go-led-strip/effects/utils"
	"go-led-strip/services/config"
	"image/color"
)

type LengthBreathing struct {
	leftColor, rightColor      color.RGBA
	colorIndex, speedIncrement float64
	running, inhaling          bool
	EffectLedSpan, frameCount  int
}

func NewLengthBreathingFromConfig(config config.EffectConfig) *LengthBreathing {
	colorIndex := 0.0
	if config.Reverse {
		colorIndex = float64(config.Span - 1)
	}
	return &LengthBreathing{
		speedIncrement: utils.GetSpeedIncrement(config.Speed, config.RefreshRate, config.MinDuration, config.MaxDuration),
		colorIndex:     colorIndex,
		EffectLedSpan:  config.Span,
		rightColor:     config.GetContinuum().Colors[0],
		leftColor:      config.GetContinuum().Colors[1],
		inhaling:       !config.Reverse,
		running:        true,
	}
}

func (effect *LengthBreathing) computeFrame() []color.RGBA {
	returnedColor := effect.getForIndex(effect.colorIndex)
	if effect.inhaling {
		effect.colorIndex += effect.speedIncrement
	} else {
		effect.colorIndex -= effect.speedIncrement
	}

	switch {
	case effect.colorIndex >= float64(effect.EffectLedSpan):
		effect.colorIndex = float64(effect.EffectLedSpan - 1)
		effect.inhaling = false

	case effect.colorIndex < 0:
		effect.inhaling = true
		effect.colorIndex = 0

	}
	return returnedColor
}

func (effect *LengthBreathing) getForIndex(index float64) []color.RGBA {
	frame := make([]color.RGBA, effect.EffectLedSpan)
	for i := 0; i < int(index); i++ {
		frame[i] = effect.leftColor
	}
	frame[int(index)] = utils.GetGradient(effect.rightColor, effect.leftColor, index-float64(int(index)))
	for i := int(index) + 1; i < effect.EffectLedSpan; i++ {
		frame[i] = effect.rightColor
	}
	return frame
}

func (effect *LengthBreathing) GetFrame() []color.RGBA {
	if !effect.running {
		return make([]color.RGBA, effect.EffectLedSpan)
	}
	return effect.computeFrame()
}

func (effect *LengthBreathing) Start() {
	effect.running = true
}

func (effect *LengthBreathing) Stop() {
	effect.running = false
}
