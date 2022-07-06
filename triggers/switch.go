package triggers

import (
        "fmt"
    	"go-led-strip/services"
        "gobot.io/x/gobot"
        "gobot.io/x/gobot/drivers/gpio"
        "golang.org/x/image/colornames"
        "gobot.io/x/gobot/platforms/raspi"
)

type SwitchTrigger struct {
	ledStripService services.LedStripService
}

func NewSwitchTrigger(ledStripService *services.LedStripService) *SwitchTrigger {
	l := SwitchTrigger{ledStripService: *ledStripService}
    l.initialize()
	return &l
}

func (switchTrigger *SwitchTrigger) initialize() {
        fmt.Println("Reed Switch initialized")
        r := raspi.NewAdaptor()

        // 31 for the Raspberry Pi 3 B+ as it h
        button := gpio.NewButtonDriver(r, "31")

        work := func() {
                button.On(gpio.ButtonPush, func(data interface{}) {
                        switchTrigger.ledStripService.Start()
                        fmt.Println("Drawer opened")
                })

                button.On(gpio.ButtonRelease, func(data interface{}) {
                        switchTrigger.ledStripService.Stop()
                	    switchTrigger.ledStripService.LedStrip.Fill(colornames.Black)
                        fmt.Println("Drawer closed")
                })
        }

        robot := gobot.NewRobot("buttonBot",
                []gobot.Connection{r},
                []gobot.Device{button},
                work,
        )

        robot.Start()
}
