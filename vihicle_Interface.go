package main

import "fmt"

type Car struct {
	name string
}

func (c Car) Accelerate() {
	fmt.Println("Speeding up ", c.name)
}
func (c Car) Brake() {
	fmt.Println("Stopping ", c.name)
}
func (c Car) Steer(direction string) {
	fmt.Println("Turnning ", direction, " ", c.name)
}

type Truck struct {
	name string
}

func (t Truck) Accelerate() {
	fmt.Println("Speeding up ", t.name)
}
func (t Truck) Brake() {
	fmt.Println("Stopping ", t.name)
}
func (t Truck) Steer(direction string) {
	fmt.Println("Turnning ", direction, " ", t.name)
}

type Vehicle interface {
	Accelerate()
	Brake()
	Steer(string)
}

func main() {
	var vehicle Vehicle = Car{name: "Opel"}
	vehicle.Accelerate()
	vehicle.Steer("Left")
	vehicle.Brake()
	vehicle = Truck{name: "ВАЗ-3110"}
	vehicle.Accelerate()
	vehicle.Steer("Right")
	vehicle.Brake()

}
