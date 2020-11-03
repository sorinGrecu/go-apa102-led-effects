package ledstrips

/*
import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"golang.org/x/image/colornames"
	"image/color"
	"time"
)

type onscreenLedStrip struct {
	LedCount   int
	canvasLeds []fyne.CanvasObject
	*fyne.Container
}

func (ledStrip *onscreenLedStrip) Test() {
	ledStrip.Fill(colornames.Red)
	time.Sleep(1000 * time.Millisecond)
	ledStrip.Fill(colornames.Green)
	time.Sleep(1000 * time.Millisecond)
	ledStrip.Fill(colornames.Blue)
	time.Sleep(1000 * time.Millisecond)
	ledStrip.Fill(colornames.Black)
}

func (ledStrip *onscreenLedStrip) GetLedCount() int {
	return ledStrip.LedCount
}

func (ledStrip *onscreenLedStrip) Paint() {
	ledStrip.Refresh()
}

func (ledStrip *onscreenLedStrip) Fill(rgba color.RGBA) {
	defer ledStrip.Refresh()
	for i := 0; i < ledStrip.LedCount; i++ {
		ledStrip.SetRGBA(i, rgba)
	}
}

func (ledStrip *onscreenLedStrip) SetRGBA(i int, rgba color.RGBA) {
	ledStrip.canvasLeds[i] = canvas.NewRectangle(rgba)
}

func NewOnscreenLedStrip(ledCount int) *onscreenLedStrip {
	objects := make([]fyne.CanvasObject, ledCount)
	for i := range objects {
		objects[i] = canvas.NewRectangle(color.Black)
	}
	boxLayout := layout.NewGridLayout(ledCount)
	container := fyne.NewContainerWithLayout(boxLayout, objects...)

	strip := &onscreenLedStrip{
		LedCount:   ledCount,
		canvasLeds: objects,
		Container:  container,
	}
	return strip
}

func (ledStrip *onscreenLedStrip) Show() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Canvas")
	myCanvas := myWindow.Canvas()
	myCanvas.SetContent(ledStrip.Container)
	myWindow.Resize(fyne.NewSize(400, 100))
	myWindow.ShowAndRun()
}
*/
