package main

import "fmt"

type Car struct {
	Type       string
	FuelInTank float64
}

func (c Car) EstimateDistance() float64 {
	var distance float64

	switch c.Type {
	case "sedan":
		distance = c.FuelInTank * 10 // 1.5L/mill, 10 mill/L
	case "suv":
		distance = c.FuelInTank * 8 // 1.5L/mill, 8 mill/L
	default:
		distance = 0.0
	}

	return distance
}

func main() {
	car1 := Car{Type: "sedan", FuelInTank: 1.5}

	fmt.Println("Estimated distance for car1:", car1.EstimateDistance())

	car2 := Car{Type: "suv", FuelInTank: 1.5}

	fmt.Println("Estimated distance for car2:", car2.EstimateDistance())
}
