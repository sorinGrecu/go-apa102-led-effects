package config

import (
	"go-led-strip/effects/utils"
	"image/color"
)

type EffectsConfig struct {
	Effects []EffectConfig `yaml:"effects"`
}

type EffectConfig struct {
	Type        string  `yaml:"type"`
	Speed       float64 `yaml:"speed"`
	Span        int     `yaml:"span"`
	Pause       int     `yaml:"pause"`
	RefreshRate int     `yaml:"refreshRate"`
	MinDuration int     `yaml:"minDuration"`
	MaxDuration int     `yaml:"maxDuration"`
	Reverse     bool    `yaml:"reverse"`
	Colors      []RGBA  `yaml:"colors"`
}

type RGBA struct {
	R uint8 `yaml:"R"`
	G uint8 `yaml:"G"`
	B uint8 `yaml:"B"`
	A uint8 `yaml:"A"`
}

func (rgba *RGBA) GetColor() color.RGBA {
	return color.RGBA{
		R: rgba.R,
		G: rgba.G,
		B: rgba.B,
		A: rgba.A,
	}
}

func (config *EffectConfig) GetContinuum() utils.ColorContinuum {
	rgba := make([]color.RGBA, len(config.Colors))
	for i := range rgba {
		rgba[i] = config.Colors[i].GetColor()
	}
	return utils.ColorContinuum{
		Colors: rgba,
	}
}
