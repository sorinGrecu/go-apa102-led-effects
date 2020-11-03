package services

import (
	"go-led-strip/effects"
	"image/color"
)

type StripFrameService struct {
	ledEffects []effects.LedEffect
	ledCount   int
}

func NewStripFrameService(ledCount int, ledEffects []effects.LedEffect) *StripFrameService {
	return &StripFrameService{
		ledCount:   ledCount,
		ledEffects: ledEffects,
	}
}

func (service *StripFrameService) getFullFrame() []color.RGBA {
	result := make([]color.RGBA, 0)
	for _, effect := range service.ledEffects {
		result = append(result, effect.GetFrame()...)
	}
	return result
}
