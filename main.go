package main

import (
	"Driver-go/elevio"
	"fmt"
)

func main() {

	numFloors := 4

	elevio.Init("localhost:15657", numFloors)

	var d elevio.MotorDirection = elevio.MD_Up
	//elevio.SetMotorDirection(d)

	drv_buttons := make(chan elevio.ButtonEvent)
	drv_floors := make(chan int)
	drv_obstr := make(chan bool)
	drv_stop := make(chan bool)

	go elevio.PollButtons(drv_buttons)
	go elevio.PollFloorSensor(drv_floors)
	go elevio.PollObstructionSwitch(drv_obstr)
	go elevio.PollStopButton(drv_stop)

	fmt.Println("Test")

	var Ele_State State

	for {
		select {
		case a := <-drv_buttons:

			switch Ele_State {

			case Idle:

			case MovingBetweenFloors:

			case MovingPassingFloor:

			case DoorsOpen:

			}

			fmt.Printf("%+v\n", a)
			elevio.SetButtonLamp(a.Button, a.Floor, true)

		case a := <-drv_floors:

			switch Ele_State {

			case Idle:

			case MovingBetweenFloors:

			case MovingPassingFloor:

			case DoorsOpen:

			}

			fmt.Printf("%+v\n", a)
			if a == numFloors-1 {
				d = elevio.MD_Down
			} else if a == 0 {
				d = elevio.MD_Up
			}
			elevio.SetMotorDirection(d)

		case a := <-drv_obstr:

			switch Ele_State {

			case Idle:

			case MovingBetweenFloors:

			case MovingPassingFloor:

			case DoorsOpen:

			}

			fmt.Printf("%+v\n", a)
			if a {
				elevio.SetMotorDirection(elevio.MD_Stop)
			} else {
				elevio.SetMotorDirection(d)
			}

		case a := <-drv_stop:
			switch Ele_State {

			case Idle:

			case MovingBetweenFloors:

			case MovingPassingFloor:

			case DoorsOpen:

			}

			fmt.Printf("%+v\n", a)
			for f := 0; f < numFloors; f++ {
				for b := elevio.ButtonType(0); b < 3; b++ {
					elevio.SetButtonLamp(b, f, false)
				}
			}
		}
	}
}

type State int

const (
	Idle                State = 0
	MovingBetweenFloors       = 1
	MovingPassingFloor        = 2
	DoorsOpen                 = 3
)
