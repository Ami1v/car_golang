package main

import (
	"fmt"
	"golang-task/car-main/car"
	"golang-task/car-main/headlights"
	"golang-task/car-main/stereo"
	"golang-task/car-main/wheels"
)

func main() {
	fmt.Println(car.OpenDoor() + headlights.TurnOn() + stereo.TurnOn() + stereo.BoostBass() + wheels.Accelerate() + wheels.Steer())
}