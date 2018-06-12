package main

import (
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/dji/tello"
	"gobot.io/x/gobot/platforms/joystick"

	"fmt"
	"os"
	"os/signal"
	"syscall"
)

const (
	dualshock4Ctl = "dualshock4"
)

// keyboard control mapping
const (
	takeOffCtrl    = joystick.TrianglePress
	landCtrl       = joystick.XPress
	stopCtrl       = joystick.CirclePress
	moveLRCtrl     = joystick.RightX
	moveFwdBkCtrl  = joystick.RightY
	moveUpDownCtrl = joystick.LeftY
	turnLRCtrl     = joystick.LeftX
	bounceCtrl     = joystick.L1Press
	palmLandCtrl   = joystick.L2Press
)

var (
	robot      *gobot.Robot
	flightMsg  = "Idle"
	flightData *tello.FlightData
	wifiData   *tello.WifiData
)

func main() {
	joystickAdaptor := joystick.NewAdaptor()
	stick := joystick.NewDriver(joystickAdaptor, "dualshock4")

	drone := tello.NewDriver("8890")

	sigChan := make(chan os.Signal, 2)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigChan
		exit()
	}()

	work := func() {
		stick.On(takeOffCtrl, func(data interface{}) {
			fmt.Println("TrianglePress + takeOffCtrl")
			fmt.Print(data)
		})
	}

	robot = gobot.NewRobot("tello",
		[]gobot.Connection{joystickAdaptor},
		[]gobot.Device{drone, stick},
		work,
	)

	robot.Start()
}

func exit() {
	robot.Stop()
	os.Exit(0)
}
