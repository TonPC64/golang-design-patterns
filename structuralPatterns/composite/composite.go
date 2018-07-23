package main

import (
	"fmt"
)

type Athlete struct{}

func (a *Athlete) Train() {
	fmt.Println("Training")
}

type CompositeSwimmerA struct {
	MyAthlete Athlete
	MySwim    func()
}

func Swim() {
	fmt.Println("Swiming!")
}

// embed
type Animal struct{}

func (a *Animal) Eat() {
	println("Eating!")
}

type Shark struct {
	Animal
	Swim func()
}

//interface all in
type Swimmer interface {
	Swim()
}

type Trainer interface {
	Train()
}

type SwimmerImpl struct{}

func (s *SwimmerImpl) Swim() {
	println("Swimming!")
}

type CompositeSwimmerB struct {
	Trainer
	Swimmer
}

func main() {

	////
	localSwim := Swim

	swimmer := CompositeSwimmerA{
		MySwim: localSwim,
	}
	swimmer.MyAthlete.Train()
	swimmer.MySwim()

	////
	fish := Shark{
		Swim: Swim,
	}

	fish.Eat()
	fish.Swim()

	////
	swimmer2 := CompositeSwimmerB{
		&Athlete{},
		&SwimmerImpl{},
	}

	swimmer2.Train()
	swimmer2.Swim()
}
