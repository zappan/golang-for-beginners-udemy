package main

import (
	"fmt"
)

// composition vs. inheritance
// COMPOSIION - EMBEDING TYPES IN OTHER TYPES

type Vehicle struct {
	NumberOfWheels int
	NumberOfPassengers int
}

type Car struct {
	Make string
	Model string
	Year int
	IsElectric bool
	IsHybrid bool
	Vehicle Vehicle  // composition, embedding the vehicle type in Car
}


func (v Vehicle) showDetails() {
	fmt.Println("NumberOfPassengers:", v.NumberOfPassengers)
	fmt.Println("NumberOfWheels:", v.NumberOfWheels)
}

func (c Car) show() {
	fmt.Println("Make:", c.Make)
	fmt.Println("Model:", c.Model)
	fmt.Println("Year:", c.Year)
	fmt.Println("IsElectric:", c.IsElectric)
	fmt.Println("IsHybrid:", c.IsHybrid)
	c.Vehicle.showDetails()
}

func main() {
	suv := Vehicle{
		NumberOfWheels: 4,
		NumberOfPassengers: 6,
	}

	volvoXC90 := Car{
		Make: "Volvo",
		Model: "XC90 T8",
		Year: 2021,
		IsElectric: false,
		IsHybrid: true,
		Vehicle: suv,
	}

	volvoXC90.show()

	fmt.Println()

	teslaModelX := Car {
		Make: "Tesla",
		Model: "Model X",
		Year: 2021,
		IsElectric: true,
		IsHybrid: false,
		Vehicle: suv,
	}

	teslaModelX.Vehicle.NumberOfPassengers = 7

	teslaModelX.show()
}
