package main

import "fmt"

// Methods: Accelerate, Brake, Steer

type Car string

func (c Car) Accelerate() {
	fmt.Println("Speeding up")
}
func (c Car) Brake() {
	fmt.Println("Slowing down")
}
func (c Car) Steer(direction string) {
	fmt.Println("Turning")
}

type Truck string

func (t Truck) Accelerate() {
	fmt.Println("Speeding up")
}
func (t Truck) Brake() {
	fmt.Println("Slowing down")
}
func (t Truck) Steer(direction string) {
	fmt.Println("Turning")
}

type Vehicle interface {
	Accelerate()
	Brake()
	Steer(string)
}

func main() {
	var vehicle Vehicle = Car("Toyota Yaris")
	vehicle.Accelerate()
	vehicle.Steer("left")

	vehicle = Truck("Toyota Tacoma")
	vehicle.Brake()
	vehicle.Steer("right")
}
