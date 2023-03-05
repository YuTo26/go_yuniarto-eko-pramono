package main

import "fmt"

type Car struct {
	Type   string
	FuelIn float64 // bahan bakar dalam liter
}

// https://pastebin.com/VzAK5VLf
func (c Car) EstimateDistance() float64 {
	const fuelPerKm = 1.5 // bahan bakar yang dibutuhkan per kilometer
	return c.FuelIn / fuelPerKm
}

func main() {
	car := Car{"Sedan", 10.5} // membuat objek Car dengan 10.5 liter bahan bakar
	distance := car.EstimateDistance()
	fmt.Printf("Jarak perkiraan yang bisa ditempuh oleh %s dengan %0.2f liter bahan bakar adalah %0.2f km\n", car.Type, car.FuelIn, distance)
}
