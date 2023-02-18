package main

import "fmt"

func main() {
	//deklarasi variable
	var nilai, x int

	//proses
	fmt.Printf("segitiga astrik\n")
	fmt.Printf("input nilai n = ")
	fmt.Scanf("%d\t", &nilai)
	fmt.Println("output : ")
	//statement
	x = nilai + 3
	//nested looping
	for i := 1; i <= nilai; i++ {
		for j := x; j >= 1; j-- {
			fmt.Printf(" ")
		}
		for k := 1; k <= i; k++ {
			fmt.Printf("* ")
		}
		fmt.Println()
		x--
	}
}
