package main

import "fmt"

type Engine interface {
	Start()
	Stop()
}

type Car struct {
	Engine
	wheelCount int
}

func (car *Car) Start() {
	fmt.Println("Car Is Starting...")
}

func (car *Car) Stop() {
	fmt.Println("Car Is Stoped")
}

func (car *Car) GoToWorkIn() {
	car.Start()
	fmt.Println("Car Now is In Work...")
	car.Stop()
}

func (car *Car) NumberOfWheels() int {
	return car.wheelCount
}

type Benz struct {
	Car
}

func (b *Benz) SayHiToMerkel() {
	fmt.Println("Hi, Merkel")
}

func main() {
	benz := &Benz{Car{wheelCount: 4}}
	benz.GoToWorkIn()
	benz.SayHiToMerkel()
	fmt.Println(benz.NumberOfWheels())
}
