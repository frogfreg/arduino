package main

import (
	"fmt"
	"time"

	"gobot.io/x/gobot/v2"
	"gobot.io/x/gobot/v2/drivers/gpio"
	"gobot.io/x/gobot/v2/platforms/firmata"
)

func main() {
	servo()
}

func servo() {

	firmataAdaptor := firmata.NewAdaptor("/dev/ttyUSB0")
	sd := gpio.NewServoDriver(firmataAdaptor, "13")

	work := func() {
		for {
			fmt.Println("setting to min...")
			if err := sd.ToMin(); err != nil {
				panic(err)
			}
			time.Sleep(1 * time.Second)
			fmt.Println("setting to center...")
			if err := sd.ToCenter(); err != nil {
				panic(err)
			}
			time.Sleep(1 * time.Second)
			fmt.Println("setting to max...")
			if err := sd.ToMax(); err != nil {
				panic(err)
			}
			time.Sleep(1 * time.Second)
		}
	}

	robot := gobot.NewRobot("bot",
		[]gobot.Connection{firmataAdaptor},
		[]gobot.Device{sd},
		work,
	)

	if err := robot.Start(); err != nil {
		panic(err)
	}

}

func led() {

	firmataAdaptor := firmata.NewAdaptor("/dev/ttyUSB0")
	led := gpio.NewLedDriver(firmataAdaptor, "13")

	work := func() {
		gobot.Every(100*time.Millisecond, func() {
			if err := led.Toggle(); err != nil {
				fmt.Println(err)
			}
		})
	}

	robot := gobot.NewRobot("bot",
		[]gobot.Connection{firmataAdaptor},
		[]gobot.Device{led},
		work,
	)

	if err := robot.Start(); err != nil {
		panic(err)
	}
}
