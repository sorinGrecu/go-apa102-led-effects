package effects

import "image/color"

const MOVING_SEGMENTS = "MOVING_SEGMENTS"
const LENGTH_BREATHING = "LENGTH_BREATHING"

type LedEffect interface {
	GetFrame() []color.RGBA
	Start()
	Stop()
}
