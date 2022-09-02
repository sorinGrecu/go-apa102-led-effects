package services

import (
	"fmt"
	"go-led-strip/ledstrips"
	"golang.org/x/image/colornames"
	"time"
)

type LedStripService struct {
	ticker time.Ticker
	StripFrameService
	ledstrips.LedStrip
	running bool
	state   chan bool
}

func NewLedStripService(frameService *StripFrameService, strip ledstrips.LedStrip) *LedStripService {
	l := LedStripService{
		StripFrameService: *frameService,
		LedStrip:          strip,
	}
	return &l
}

func (stripService *LedStripService) InitTicker() {
	stripService.ticker = *time.NewTicker(40 * time.Millisecond)
	stripService.state = make(chan bool)
	go func() {
		for {
			select {
			case <-stripService.ticker.C:
				if stripService.running {
					frame := stripService.getFullFrame()
					for i, f := range frame {
						stripService.SetLed(i, f)
					}
					stripService.LedStrip.Paint()
				}
			case stripService.running = <-stripService.state:
				fmt.Println("state:", stripService.running)
			}
		}
	}()
}

func (stripService *LedStripService) Start() {
	stripService.state <- true
}

func (stripService *LedStripService) Stop() {
	stripService.state <- false
}

func (stripService *LedStripService) Toggle() {
	stripService.state <- !<-stripService.state
}

func (stripService *LedStripService) Test() {
	stripService.LedStrip.Fill(colornames.Red)
	time.Sleep(200 * time.Millisecond)
	stripService.LedStrip.Fill(colornames.Green)
	time.Sleep(200 * time.Millisecond)
	stripService.LedStrip.Fill(colornames.Blue)
	time.Sleep(200 * time.Millisecond)
	stripService.LedStrip.Fill(colornames.Black)
}
