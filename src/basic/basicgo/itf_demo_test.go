package basicgo

import (
	"fmt"
	"testing"
)

type CarTypeGetter interface {
	GetCarType() string
}

type Car struct {
	Name string
}

type SportCar Car

func NewSportCar() *SportCar {
	return &SportCar{}
}

func (s *SportCar) GetCarType() string {
	return "sport car"
}

type PassengerCar Car

func NewPassengerCar() *PassengerCar {
	return &PassengerCar{}
}

func (s *PassengerCar) GetCarType() string {
	return "passenger car"
}

// do something with car type
func procCarType(car CarTypeGetter)  {
	fmt.Println(car.GetCarType())
}

func TestItfCarTypeDemo(t *testing.T)  {
	sportCar := NewSportCar()
	passengerCar := NewPassengerCar()
	procCarType(sportCar)
	procCarType(passengerCar)
}