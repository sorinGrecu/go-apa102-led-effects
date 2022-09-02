package ledstrips

import (
	"go-led-strip/services/config"
	"image/color"
)

type mock struct {
	LedCount   int
	brightness uint8
}

func NewMockLedStripFromConfig(config *config.LedStripConfig) *mock {
	strip := &mock{
		LedCount:   config.LedCount,
		brightness: config.Brightness,
	}
	return strip
}

func (ledStrip *mock) GetLedCount() int {
	return ledStrip.LedCount
}

func (ledStrip *mock) Fill(rgba color.RGBA) {

}

func (ledStrip *mock) Paint() {
}

func (ledStrip *mock) SetLed(i int, rgba color.RGBA) {
}
