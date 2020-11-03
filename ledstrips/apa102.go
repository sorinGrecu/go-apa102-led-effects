package ledstrips

import (
	"go-led-strip/services/config"
	"gobot.io/x/gobot/drivers/spi"
	"gobot.io/x/gobot/platforms/raspi"
	"image/color"
)

type apa102 struct {
	LedCount   int
	brightness uint8
	driver     spi.APA102Driver
}

func NewApa102LedStripFromConfig(config *config.LedStripConfig) *apa102 {
	raspiAdaptor := raspi.NewAdaptor()
	driver := *spi.NewAPA102Driver(raspiAdaptor, config.LedCount, config.Brightness)
	strip := &apa102{
		LedCount:   config.LedCount,
		brightness: config.Brightness,
		driver:     driver,
	}
	strip.driver.Start()
	return strip
}

func (ledStrip *apa102) GetLedCount() int {
	return ledStrip.LedCount
}

func (ledStrip *apa102) Fill(rgba color.RGBA) {
	defer ledStrip.driver.Draw()
	for i := 0; i < ledStrip.LedCount; i++ {
		ledStrip.driver.SetRGBA(i, rgba)
	}
}

func (ledStrip *apa102) Paint() {
	ledStrip.driver.Draw()
}

func (ledStrip *apa102) SetLed(i int, rgba color.RGBA) {
	ledStrip.driver.SetRGBA(i, rgba)
}

/*func (ledStrip *apa102) Test() {
	ledStrip.Fill(colornames.Red)
	time.Sleep(100 * time.Millisecond)
	ledStrip.Fill(colornames.Green)
	time.Sleep(100 * time.Millisecond)
	ledStrip.Fill(colornames.Blue)
	time.Sleep(100 * time.Millisecond)
	ledStrip.Fill(colornames.Black)
}*/
