package ledstrips

import "image/color"

const APA102 = "APA102"
const MOCK = "MOCK"

type LedStrip interface {
	Fill(rgba color.RGBA)
	SetLed(i int, rgba color.RGBA)
	GetLedCount() int
	Paint()
}
